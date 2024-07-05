# Shelflove-mvc
 
## To run locally
- Clone the repo ```git clone git@github.com:Arshdeep54/Shelflove-mvc.git```
- ```cd Shelflove-mvc ```
- rename `.env.example` file to `.env` and fill it with your db info .
- ```make help``` or just ```make``` : It will help you with further installation 

### Install air 
```bash
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

### For Runnning via  docker
```bash
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo docker-compose --version

```