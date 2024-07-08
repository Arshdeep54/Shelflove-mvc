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

```bash
chmod +x docker.sh
./docker.sh
```
- Go to http://cosign.org 