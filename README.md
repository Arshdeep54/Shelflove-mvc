# Shelflove-mvc
 
## Setup  
- open a your development folder
- ```wget https://raw.githubusercontent.com/Arshdeep54/Shelflove-mvc/main/setup.sh``` 
- ```chmod +x host.sh #To change execution permission```
- ```./setup.sh```

### Install air 
- Also asked to install in setup script 
```bash
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```
### Install make 
- If while running you get error make not found 
- ```sudo apt install make```
### For Runnning via  docker

- Rename .env.example to .env 
  - Add any jwt secret and dont change the db configurations , docker compose uses that 

```bash

#install docker-compose if not 
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo docker-compose --version

sudo docker-compose up --build #to see the logs and build the image 
sudo docker-compose up --build -d #to run in detach mode 


```
- Go to http://cosign.org 