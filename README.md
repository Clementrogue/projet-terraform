# AWS Terraform - Student Records (Phase 1)

## Description

Ce projet déploie une application web de gestion des dossiers étudiants sur AWS en utilisant **Terraform**.

L'objectif de la phase 1 est de créer une **infrastructure simple** permettant d'héberger l'application sur une **instance EC2** accessible depuis Internet.

---

## Architecture
Internet
│
Internet Gateway
│
VPC
│
Public Subnet
│
EC2 (Ubuntu)
│
Node.js Application + MySQL


---

## Infrastructure AWS

Les ressources suivantes sont créées avec Terraform :

- VPC
- Subnet public
- Internet Gateway
- Route Table
- Security Group
- Instance EC2

Ports ouverts :

- **22** : SSH
- **80** : HTTP

---
<img width="501" height="681" alt="image" src="https://github.com/user-attachments/assets/0f3b0c67-7a19-49b5-8a33-6784aad7882e" />

## Déploiement

### Initialiser Terraform

```bash
terraform init
```

### Vérifier la configuration

```bash
terraform validate
```

### Voir le plan

```bash
terraform plan
```

### Déployer l'infrastructure

```bash
terraform apply
```

## Accès à l'application
Une fois l'infrastructure déployée, l'application est accessible via l'adresse publique de l'instance EC2 :
```bash
http://PUBLIC_IP
```
Terraform affiche l'adresse IP dans les outputs après le déploiement.

