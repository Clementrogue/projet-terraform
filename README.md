<img width="1024" height="1536" alt="image" src="https://github.com/user-attachments/assets/58302e7d-fb47-43ad-b0b7-d518bedb2f35" />

# Phase 6 — Déploiement d’une application sur AWS ECS avec Terraform

## Présentation du projet

Cette phase du projet consiste à déployer une application conteneurisée sur **Amazon ECS (Elastic Container Service)** en utilisant **Terraform** pour automatiser la création de l’infrastructure.

L'objectif est de déployer une application Docker sur une infrastructure AWS entièrement configurée par **Infrastructure as Code (IaC)**.

L’infrastructure déployée comprend :

* Un **VPC**
* Deux **subnets publics**
* Une **Internet Gateway**
* Des **tables de routage**
* Un **Security Group**
* Un **cluster ECS**
* Une **instance EC2 avec l’agent ECS**
* Un **Auto Scaling Group**
* Une **Task Definition ECS**
* Un **Service ECS**

L’application est exécutée dans un **conteneur Docker** et accessible via l’adresse IP publique de l’instance EC2.

---

# Architecture de l’infrastructure

L’architecture déployée est la suivante :

```
Internet
   │
Internet Gateway
   │
VPC (10.2.0.0/16)
   │
   ├── Subnet Public A (10.2.1.0/24)
   │       │
   │       └── Instance EC2 (ECS)
   │              │
   │              └── Conteneur Docker (Application)
   │
   └── Subnet Public B (10.2.2.0/24)
```

Les principaux composants utilisés sont :

| Composant          | Rôle                               |
| ------------------ | ---------------------------------- |
| VPC                | Réseau privé pour l’infrastructure |
| Subnets            | Segmentation du réseau             |
| Internet Gateway   | Accès Internet                     |
| Security Group     | Contrôle des accès réseau          |
| ECS Cluster        | Orchestration des conteneurs       |
| Launch Template    | Configuration des instances EC2    |
| Auto Scaling Group | Gestion automatique des instances  |
| Task Definition    | Définition du conteneur            |
| ECS Service        | Maintien des conteneurs actifs     |

---

# Prérequis

Avant de lancer le projet, les outils suivants doivent être installés :

* **Terraform ≥ 1.3**
* **AWS CLI**
* **Docker**
* Un **compte AWS**
* Une **clé SSH**

Configurer les identifiants AWS :

```bash
aws configure
```

---

# Structure du projet

```
phase6-ecs/
│
├── main.tf
├── variables.tf
├── terraform.tfvars
├── outputs.tf
└── README.md
```

---

# Déploiement de l’infrastructure

Initialiser Terraform :

```bash
terraform init
```

Vérifier la configuration :

```bash
terraform validate
```

Visualiser les ressources qui seront créées :

```bash
terraform plan
```

Déployer l’infrastructure :

```bash
terraform apply
```

Terraform va créer automatiquement :

* le VPC
* les subnets
* l’internet gateway
* les règles réseau
* le cluster ECS
* l’instance EC2
* le service ECS
* le conteneur Docker

---

# Accès à l’application

Une fois le déploiement terminé, récupérer l’adresse IP publique de l’instance EC2 :

```bash
aws ec2 describe-instances \
--filters "Name=tag:Name,Values=student-phase6-ecs-instance" \
--query "Reservations[*].Instances[*].PublicIpAddress" \
--output text
```

Accéder ensuite à l’application via le navigateur :

```
http://IP_PUBLIQUE
```

ou

```
http://IP_PUBLIQUE:81
```

selon la configuration du port.

---

# Vérification du déploiement ECS

Lister les clusters ECS :

```bash
aws ecs list-clusters
```

Lister les tâches en cours d’exécution :

```bash
aws ecs list-tasks --cluster student-phase6-cluster
```

Vérifier le service ECS :

```bash
aws ecs describe-services \
--cluster student-phase6-cluster \
--services student-phase6-service
```

---

# Connexion à l’instance ECS

Connexion SSH :

```bash
ssh -i vockey.pem ec2-user@IP_PUBLIQUE
```

Vérifier les conteneurs Docker :

```bash
docker ps
```

Voir les logs du conteneur :

```bash
docker logs ID_CONTENEUR
```

---

# Dépannage

Si l’application n’est pas accessible :

### Vérifier les tâches ECS

```bash
aws ecs list-tasks --cluster student-phase6-cluster
```

### Vérifier les événements du service

```bash
aws ecs describe-services \
--cluster student-phase6-cluster \
--services student-phase6-service \
--query "services[0].events"
```

### Vérifier les conteneurs sur l’instance

```bash
docker ps -a
```

### Vérifier l’écoute du port

```bash
sudo ss -tulpn | grep :80
```

---

# Suppression de l’infrastructure

Pour supprimer toutes les ressources AWS :

```bash
terraform destroy
```

Projet Cloud — Déploiement d’infrastructure sur AWS
Phase 6 — ECS avec Terraform
