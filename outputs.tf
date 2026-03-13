output "ecs_cluster_name" {
  value = aws_ecs_cluster.main.name
}

output "ecs_service_name" {
  value = aws_ecs_service.app.name
}

output "ecs_task_definition" {
  value = aws_ecs_task_definition.app.family
}

output "instance_sg_id" {
  value = aws_security_group.ecs_instances_sg.id
}

output "public_subnet_a" {
  value = aws_subnet.public_a.id
}
