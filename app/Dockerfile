FROM node:18

# Installation des outils nécessaires
RUN apt-get update && apt-get install -y wget unzip

# Dossier de travail dans le conteneur
WORKDIR /home/ubuntu

# Téléchargement et extraction du vrai code AWS
RUN wget https://aws-tc-largeobjects.s3.us-west-2.amazonaws.com/CUR-TF-200-ACCAP1-1-91571/1-lab-capstone-project-1/code.zip
RUN unzip code.zip -x "resources/codebase_partner/node_modules/*"

# On se place dans le dossier de l'application
WORKDIR /home/ubuntu/resources/codebase_partner

# Installation des dépendances demandées par ton script
RUN npm install aws aws-sdk
RUN npm install

# On expose le port 80
EXPOSE 80

# On lance l'application
CMD ["npm", "start"]
