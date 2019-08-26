## 1. Run docker container

    docker-compose up --build --force-recreate

## 2. View Docker status

    docker container ls -a

## 3. SSH to docker

    docker exec -it core_app_1 /bin/bash

### Start the web server:

    cd $GOPATH/src
    git clone <this repo>
    go get ./...
    revel run revel-golang-app

### Go to http://localhost:9000
