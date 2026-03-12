resource "aws_ecr_repository" "app_repo" {
  name                 = "student-app-repo"
  image_tag_mutability = "MUTABLE"
  force_delete         = true # Très utile en environnement de lab pour pouvoir tout détruire facilement à la fin

  image_scanning_configuration {
    scan_on_push = true
  }
}

output "ecr_repository_url" {
  value       = aws_ecr_repository.app_repo.repository_url
  description = "L'URL du registre ECR pour pousser l'image Docker"
}
