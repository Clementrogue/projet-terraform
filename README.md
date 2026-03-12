# Phase 5 – Mise en place d'une pipeline CI/CD

## Objectif

L'objectif de cette phase est d'automatiser le déploiement de l'application en utilisant une pipeline CI/CD.

La pipeline permet d'exécuter automatiquement différentes étapes lors d'un push dans le dépôt Git :

- Vérification de la qualité du code
- Build de l'application
- Création et publication d'une image Docker
- Test de charge
- Déploiement automatique sur l'environnement de production

Cette approche permet d'améliorer la fiabilité et la rapidité du déploiement.

---

# Architecture CI/CD

Le workflow de la pipeline est le suivant :

Développeur  
↓  
GitLab Repository  
↓  
GitLab CI/CD Pipeline  
↓  
Build Docker Image  
↓  
Docker Registry  
↓  
Test de charge (k6)  
↓  
Déploiement automatique sur le serveur EC2

---

# Pipeline GitLab

La pipeline est définie dans le fichier :

```
.gitlab-ci.yml
```

Elle est composée des stages suivants :

- **quality**
- **build**
- **package**
- **loadtest**
- **deploy**

---

# Étapes de la pipeline

## 1. Quality

Cette étape vérifie la qualité du projet :

- installation des dépendances Node.js
- audit de sécurité avec npm
- exécution des tests

Commande exécutée :

```
npm ci
npm audit
npm test
```

---

## 2. Build

Compilation de l'application.

Commande :

```
npm run build
```

---

## 3. Package

Création de l'image Docker de l'application et publication dans le registry GitLab.

Commandes utilisées :

```
docker build
docker push
```

---

## 4. Load Test

Un test de charge est exécuté avec **k6** pour vérifier que l'application supporte plusieurs utilisateurs.

Exemple :

```
k6 run k6-loadtest.js
```

Ce test simule plusieurs utilisateurs accédant à l'application.

---

## 5. Deploy

Si toutes les étapes précédentes réussissent, l'application est déployée automatiquement sur le serveur de production.

Le script utilisé est :

```
scripts/deploy.sh
```

Le script :

- se connecte au serveur EC2 via SSH
- télécharge la dernière image Docker
- redémarre le conteneur de l'application

---

# Technologies utilisées

- GitLab CI/CD
- Docker
- Node.js
- Grafana k6
- SSH
- Amazon EC2

---

# Résultat

Grâce à la pipeline CI/CD :

- chaque modification du code déclenche automatiquement un build
- les tests sont exécutés automatiquement
- l'image Docker est générée et stockée
- l'application est déployée automatiquement

Cela permet d'avoir un **déploiement rapide, fiable et automatisé**.

---

# Améliorations possibles

Pour améliorer cette architecture :

- ajout d'un environnement de staging
- ajout de tests automatisés plus avancés
- déploiement sur Kubernetes
- monitoring de la pipeline

---

# Conclusion

La mise en place de la pipeline CI/CD permet d'automatiser entièrement le cycle de déploiement de l'application.

Cela améliore la productivité des développeurs et réduit les erreurs humaines lors des déploiements.
