#!/bin/bash

make build

read -p "Want to install dcoker-conpose ? y/N " install 
if [[ "$install" =~ ^[Yy]$ ]]; then
 sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
 sudo chmod +x /usr/local/bin/docker-compose
 sudo docker-compose --version
fi 


ENV_CONFIG=$(cat <<EOF
MYSQL_HOST=db  #for docker change it to db
MYSQL_USERNAME=user   #for docker change it to user
MYSQL_PASSWORD=password   #for docker change it to password
MYSQL_DATABASE=shelflove  
JWT_SECRET=dockerSecretly
JWT_EXPIRATION=5184000 #1 day
MYSQL_PORT=3306
NODE_ENV="production"
EOF
)
rm .env
touch .env
echo "${ENV_CONFIG}" >> .env


echo " Go to http://cosign.org once done "

sudo docker-compose up --build #to see the logs and build the image 
# sudo docker-compose up --build -d #to run in detach mode 
