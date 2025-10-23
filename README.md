# Run locally

```
go mod init web-app
go mod tidy
go build -o web-app
go run main.go
```

# Using Docker

## Create Docker Image

docker build -t Z3ter-birthday-app:v1.0 .

docker images | grep Z3ter

## Run Container

docker run -d -p 8000:8080 --name Z3ter-app Z3ter-birthday-app:v1.0

docker ps

## Test application

localhost:8000

## Inside the container

docker exec -it Z3ter-app sh

# Push to DockerHub

docker login

docker tag Z3ter-birthday-app:v1.0 ashrakat112/Z3ter-birthday-app:v1.0

docker tag Z3ter-birthday-app:v1.0 ashrakat112/Z3ter-birthday-app:latest

docker push ashrakat112/Z3ter-birthday-app:v1.0
docker push ashrakat112/Z3ter-birthday-app:latest

# Run Container from image on Dockerhub

docker pull ashrakat112/Z3ter-birthday-app:latest

docker run -d -p 8080:8080 --name Z3ter-app ashrakat112/Z3ter-birthday-app:latest

# Remove everything

## Stop and remove container

docker stop Z3ter-app
docker rm Z3ter-app

## Remove all local images

docker rmi Z3ter-birthday-app:v1.0
docker rmi ashrakat112/Z3ter-birthday-app:v1.0
docker rmi ashrakat112/Z3ter-birthday-app:latest

docker images
