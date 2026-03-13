# Phase 3 : Haute Disponibilité et Mise à l'Échelle Automatique

## Objectif
L'objectif de cette phase est de rendre l'application web hautement disponible et résiliente aux pics de trafic en utilisant les services AWS d'équilibrage de charge et d'Auto Scaling via Terraform.

## Architecture déployée
* **Application Load Balancer (ALB) :** Répartit le trafic entrant (HTTP port 80) sur plusieurs instances EC2 situées dans des sous-réseaux publics.
* **Target Group & Health Checks :** Vérifie en continu la santé des instances (HTTP 200).
* **Launch Template :** Modèle de configuration EC2 intégrant la bonne AMI (Phase 2) et le profil IAM `LabInstanceProfile`.
* **Auto Scaling Group (ASG) :** Déploie les instances dans des sous-réseaux privés (min: 2, max: 4).
* **Target Tracking Scaling Policy :** Alarme CloudWatch configurée (pour les besoins du lab) pour déclencher une mise à l'échelle dès que le CPU moyen dépasse **20%**.


<img width="1330" height="1810" alt="image" src="https://github.com/user-attachments/assets/d6535dc8-4ad6-421a-9fdf-895820f6e78c" />


## Déploiement

1. **Initialiser et appliquer l'infrastructure :**
   \`\`\`bash
   terraform init
   terraform apply -auto-approve
   \`\`\`
2. **Récupérer l'URL de l'application :**
   Terraform affichera en sortie la variable `alb_dns_name`. Copiez cette URL dans un navigateur pour accéder à l'application.

## Validation & Test de charge
Pour prouver le bon fonctionnement de la mise à l'échelle (Scale-out) :
1. Installer l'outil de test de charge :
   \`\`\`bash
   npm install -g loadtest
   \`\`\`
2. Lancer un bombardement de requêtes pendant 3 minutes (100 utilisateurs simultanés) :
   \`\`\`bash
   loadtest -t 180 -c 100 http://<VOTRE_ALB_DNS_NAME>/
   \`\`\`
3. Observer la console AWS (EC2 > Auto Scaling Groups > Activity) : de nouvelles instances sont automatiquement créées pour absorber la charge.

<img width="3054" height="1156" alt="image" src="https://github.com/user-attachments/assets/1a6171ac-c0a3-4c74-828d-50ee5b3576ac" />

<img width="3032" height="1612" alt="image" src="https://github.com/user-attachments/assets/a2ccff37-0adb-44f3-abaa-1aca278530d8" />
