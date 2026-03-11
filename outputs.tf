output "web_public_ip" {
  value = aws_instance.web.public_ip
}

output "web_url" {
  value = "http://${aws_instance.web.public_ip}"
}

output "rds_endpoint" {
  value = aws_db_instance.mysql.address
}

output "secret_name" {
  value = aws_secretsmanager_secret.db_secret.name
}
