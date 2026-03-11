package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfrastructure(t *testing.T) {
	t.Parallel()

	tfOptions := &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"vpc_id":                    "vpc-12345",
			"public_subnets":            []string{"subnet-pub1"},
			"private_subnets":           []string{"subnet-priv1", "subnet-priv2"},
			"app_ami_id":                "ami-12345678",
			"iam_instance_profile_name": "test-profile",
		},
	}

	planStruct := terraform.InitAndPlanAndShowWithStruct(t, tfOptions)

	t.Run("ALB_SG_allows_ingress_port_80_from_anywhere", func(t *testing.T) {
		albSG := planStruct.ResourcePlannedValuesMap["aws_security_group.alb_sg"]
		require.NotNil(t, albSG, "aws_security_group.alb_sg not found in plan")

		ingress, ok := albSG.AttributeValues["ingress"].([]interface{})
		require.True(t, ok && len(ingress) > 0, "ALB SG should have ingress rules")

		rule := ingress[0].(map[string]interface{})
		assert.Equal(t, float64(80), rule["from_port"])
		assert.Equal(t, float64(80), rule["to_port"])
		assert.Equal(t, "tcp", rule["protocol"])

		cidrBlocks, ok := rule["cidr_blocks"].([]interface{})
		require.True(t, ok)
		assert.Contains(t, cidrBlocks, "0.0.0.0/0", "ALB SG should allow ingress from anywhere")
	})

	t.Run("ASG_SG_allows_ingress_port_80_only_from_ALB_SG", func(t *testing.T) {
		asgSG := planStruct.ResourcePlannedValuesMap["aws_security_group.asg_sg"]
		require.NotNil(t, asgSG, "aws_security_group.asg_sg not found in plan")

		ingress, ok := asgSG.AttributeValues["ingress"].([]interface{})
		require.True(t, ok && len(ingress) > 0, "ASG SG should have ingress rules")

		rule := ingress[0].(map[string]interface{})
		assert.Equal(t, float64(80), rule["from_port"])
		assert.Equal(t, float64(80), rule["to_port"])
		assert.Equal(t, "tcp", rule["protocol"])

		// Must NOT have CIDR blocks — traffic is restricted to the ALB SG only
		cidrBlocks, _ := rule["cidr_blocks"].([]interface{})
		assert.Empty(t, cidrBlocks, "ASG SG ingress must not allow CIDR-based access")

		// Must reference the ALB security group
		secGroups, _ := rule["security_groups"].([]interface{})
		assert.NotEmpty(t, secGroups, "ASG SG ingress must reference the ALB security group")
	})

	t.Run("ALB_created_with_correct_name_type_and_security_groups", func(t *testing.T) {
		alb := planStruct.ResourcePlannedValuesMap["aws_lb.app_alb"]
		require.NotNil(t, alb, "aws_lb.app_alb not found in plan")

		assert.Equal(t, "app-alb", alb.AttributeValues["name"])
		assert.Equal(t, "application", alb.AttributeValues["load_balancer_type"])
		assert.Equal(t, false, alb.AttributeValues["internal"])

		secGroups, ok := alb.AttributeValues["security_groups"].([]interface{})
		require.True(t, ok)
		assert.Len(t, secGroups, 1, "ALB should have exactly one security group attached")
	})

	t.Run("ALB_target_group_health_check", func(t *testing.T) {
		tg := planStruct.ResourcePlannedValuesMap["aws_lb_target_group.app_tg"]
		require.NotNil(t, tg, "aws_lb_target_group.app_tg not found in plan")

		healthChecks, ok := tg.AttributeValues["health_check"].([]interface{})
		require.True(t, ok && len(healthChecks) > 0, "Target group should have a health check configured")

		hc := healthChecks[0].(map[string]interface{})
		assert.Equal(t, "/", hc["path"])
		assert.Equal(t, "HTTP", hc["protocol"])
	})

	t.Run("ASG_launch_template_vpc_zones_and_scaling_policy", func(t *testing.T) {
		asg := planStruct.ResourcePlannedValuesMap["aws_autoscaling_group.app_asg"]
		require.NotNil(t, asg, "aws_autoscaling_group.app_asg not found in plan")

		// Verify VPC zone identifiers match the private subnets
		vpcZones, ok := asg.AttributeValues["vpc_zone_identifier"].([]interface{})
		require.True(t, ok)
		assert.Contains(t, vpcZones, "subnet-priv1")
		assert.Contains(t, vpcZones, "subnet-priv2")

		// Verify launch template is attached with $Latest version
		launchTemplates, ok := asg.AttributeValues["launch_template"].([]interface{})
		require.True(t, ok && len(launchTemplates) > 0, "ASG should have a launch template")
		lt := launchTemplates[0].(map[string]interface{})
		assert.Equal(t, "$Latest", lt["version"])

		// Verify the CPU-tracking scaling policy
		policy := planStruct.ResourcePlannedValuesMap["aws_autoscaling_policy.cpu_tracking"]
		require.NotNil(t, policy, "aws_autoscaling_policy.cpu_tracking not found in plan")
		assert.Equal(t, "TargetTrackingScaling", policy.AttributeValues["policy_type"])
		assert.Equal(t, "cpu-tracking-policy", policy.AttributeValues["name"])

		ttConfigs, ok := policy.AttributeValues["target_tracking_configuration"].([]interface{})
		require.True(t, ok && len(ttConfigs) > 0)
		ttConfig := ttConfigs[0].(map[string]interface{})
		assert.Equal(t, float64(50), ttConfig["target_value"])
	})
}
