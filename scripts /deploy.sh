#!/bin/bash
set -e

mkdir -p ~/.ssh
echo "$SSH_PRIVATE_KEY" > ~/.ssh/id_rsa
chmod 600 ~/.ssh/id_rsa

ssh-keyscan -H "$DEPLOY_HOST" >> ~/.ssh/known_hosts

ssh ubuntu@$DEPLOY_HOST << EOF
  docker login -u "$CI_REGISTRY_USER" -p "$CI_REGISTRY_PASSWORD" "$CI_REGISTRY"
  docker pull $CI_REGISTRY_IMAGE/student-app:latest
  docker stop student-app || true
  docker rm student-app || true
  docker run -d \
    --name student-app \
    --restart always \
    -p 80:3000 \
    -e APP_DB_HOST="$APP_DB_HOST" \
    -e APP_DB_USER="$APP_DB_USER" \
    -e APP_DB_PASSWORD="$APP_DB_PASSWORD" \
    -e APP_DB_NAME="$APP_DB_NAME" \
    -e APP_PORT="3000" \
    $CI_REGISTRY_IMAGE/student-app:latest
EOF
