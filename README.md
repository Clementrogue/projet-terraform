# Phase 4 : Packaging de l'application (Docker & ECR)

## Objectif
L'objectif de cette phase est de moderniser le déploiement de l'application en la conteneurisant avec Docker, puis en stockant l'image de manière sécurisée sur un registre privé AWS ECR (Elastic Container Registry). 

## Architecture déployée
* **Dockerfile :** Création d'une image basée sur `node:18`, téléchargeant le code source de l'application depuis S3, installant les dépendances (`npm install`) et exposant le port 80.
* **Amazon ECR :** Registre privé Terraformé pour héberger l'image Docker (`student-app-repo`).
* **Instance EC2 de test :** Déploiement via Terraform d'une machine isolée avec un script `user_data` pour installer Docker, s'authentifier sur ECR, et lancer le conteneur automatiquement.

## Déploiement et Stockage de l'image

1. **Créer le registre ECR avec Terraform :**
   \`\`\`bash
   terraform apply -target=aws_ecr_repository.app_repo -auto-approve
   \`\`\`
2. **Construire l'image Docker localement :**
   \`\`\`bash
   cd app/
   docker build -t student-app-repo:v1 .
   \`\`\`
3. **S'authentifier et pousser sur AWS ECR :**
   \`\`\`bash
   aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com
   docker tag student-app-repo:v1 <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/student-app-repo:v1
   docker push <AWS_ACCOUNT_ID>.dkr.ecr.us-east-1.amazonaws.com/student-app-repo:v1
   \`\`\`

## Validation des conteneurs

**1. Test local (Terminal) :**
Exécution du conteneur en simulant les variables d'environnement de la base de données :
\`\`\`bash
docker run -d -p 8080:80 -e APP_DB_HOST=127.0.0.1 -e APP_DB_USER=nodeapp -e APP_DB_PASSWORD=student12 -e APP_DB_NAME=STUDENTS -e APP_PORT=80 student-app-repo:v1
curl http://localhost:8080
\`\`\`
*(Vérification : Le code HTML de la page d'accueil s'affiche correctement, confirmant que le serveur Node.js tourne dans le conteneur).*

**2. Test Cloud (AWS EC2) :**
Déploiement de l'infrastructure de test Terraform :
\`\`\`bash
terraform apply -auto-approve
\`\`\`
*(Vérification : Copier l'IP publique générée en sortie par Terraform dans un navigateur pour visualiser l'application conteneurisée tournant sur AWS).*
