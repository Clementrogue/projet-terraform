# Phase 2 – Découplage de l'application

## Objectif

L'objectif de cette phase est de séparer l'application web et la base de données afin qu'elles fonctionnent indépendamment.

Dans la phase 1, l'application Node.js et la base MySQL étaient installées sur la même machine virtuelle.

Dans cette phase, l'architecture est améliorée en utilisant :

- Amazon EC2 pour l'application web
- Amazon RDS MySQL pour la base de données
- AWS Secrets Manager pour stocker les identifiants de la base

Cette approche permet d'améliorer la sécurité, la maintenabilité et la scalabilité de l'application.

---

# Architecture

L'architecture mise en place est la suivante :

Internet  
│  
│  
EC2 (Ubuntu - Application Node.js)  
│  
│  
Amazon RDS (MySQL)

---

# Composants utilisés

### Amazon EC2
- Héberge l'application Node.js
- Instance Ubuntu Server 22.04

### Amazon RDS
- Base de données MySQL
- Service géré par AWS

### AWS Secrets Manager
- Stockage sécurisé des identifiants de la base de données

### VPC
- Réseau isolé pour l'infrastructure

---

# Infrastructure déployée avec Terraform

Les ressources suivantes sont créées :

- VPC
- Subnet public (EC2)
- Subnets privés (RDS)
- Internet Gateway
- Route Table
- Security Groups
- Instance EC2
- Base de données RDS
- Secret dans Secrets Manager

---

# Déploiement

## Initialisation Terraform

```bash
terraform init
```

## Vérification de la configuration

```bash
terraform validate
```

## Plan Terraform

```bash
terraform plan
```

## Déploiement de l'infrastructure

```bash
terraform apply
```

---

# Accès à l'application

Une fois l'infrastructure déployée, Terraform affiche l'adresse IP publique de l'application.

Exemple :

```
http://44.206.0.228
```

L'application permet de :

- consulter la liste des étudiants
- ajouter un étudiant
- modifier un étudiant
- supprimer un étudiant

---

# Technologies utilisées

- Terraform
- AWS EC2
- Amazon RDS (MySQL)
- AWS Secrets Manager
- Node.js
- Ubuntu Server 22.04

---

# Résultat

La base de données est maintenant séparée de l'application, ce qui permet :

- une meilleure sécurité
- une meilleure gestion de la base de données
- une architecture plus scalable

Cette phase prépare l'infrastructure pour la phase suivante où un load balancer et un auto scaling group seront ajoutés.
