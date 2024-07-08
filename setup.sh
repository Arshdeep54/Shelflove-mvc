#!/bin/bash
read -p "Want to clone ? y/N " clone

if [[ "$clone" =~ ^[Yy]$ ]]; then
  git clone git@github.com:Arshdeep54/Shelflove-mvc.git
fi

BLUE='\033[0;34m'
DEFAULT='\033[0m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'

cd Shelflove-mvc

read -p "Enter Your mysql username: " MYSQL_USERNAME
read -s -p "Enter Your mysql password: " MYSQL_PASSWORD
echo
read -s -p "Enter a secure jwt secret: " JWT_SECRET
echo
echo "You can run in three ways"
echo -e "1. at localhost:8000 with ${YELLOW}make dev${DEFAULT}"
echo -e "2. at http://cosign.org with apache hosting ${YELLOW}make host${DEFAULT}"
echo -e "3. with docker using docker-compose.yml (run ./docker.sh) "
echo 
MYSQL_HOST=localhost

read -p "You want to run via docker :y/N " docker

if [[ "$docker" =~ ^[Yy]$ ]]; then
  echo "Setting Docker .env variables ..."
  sleep 2 
  MYSQL_USERNAME=user
  MYSQL_PASSWORD=password
  MYSQL_HOST=db   
fi

ENV_CONFIG=$(cat <<EOF
MYSQL_HOST=${MYSQL_HOST}  #for docker change it to db
MYSQL_USERNAME=${MYSQL_USERNAME}    #for docker change it to user
MYSQL_PASSWORD='${MYSQL_PASSWORD}'    #for docker change it to password
MYSQL_DATABASE=shelflove  
JWT_SECRET='${JWT_SECRET}'
JWT_EXPIRATION=5184000 #1 day
MYSQL_PORT=3306
NODE_ENV="production"
EOF
)

# rm .env.example
touch .env
echo "${ENV_CONFIG}" >> .env
if [[ "$docker" =~ ^[Yy]$ ]]; then
  echo " "
else
  echo "Migrating ... Please enter your mysql password"
  mysql -u ${MYSQL_USERNAME} -p < ./pkg/config/migrations/initial.up.sql
  make install
  read -p "$(echo -e "Want to install air(hot reload) for ${YELLOW}make dev${DEFAULT}: y/N ")" air
  echo " "
  if [[ "$air" =~ ^[Yy]$ ]]; then
    echo "Installing air "
    curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
  fi
  echo " "
  echo -e "Running ${YELLOW}make help${DEFAULT} for further testing"
  make help
fi
echo " "
echo -e "    ${BLUE}cd Shelflove-mvc${DEFAULT}"





