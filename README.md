
## Clone this repo

    $ cd $GOPATH/src
    $ git clone <this repo>

## 1. Work with non docker

### Get some go module

    $ go get github.com/revel/revel
    $ go get github.com/revel/cmd/revel
    $ go get github.com/go-gorp/gorp

    // Config file
    $ cp conf/app.conf.dev conf/app.conf
    // Then open app.conf and edit sql config

### Install mysql

    // Do by yourself

### Run this app

    $ revel run revel-golang-app dev

## 2. Play with docker and mysql

### Run app

    $ docker-compose up --build --force-recreate

### View Docker status

    $ docker container ls -a

### SSH to docker

    $ docker exec -it revel-golang-app /bin/bash
    // Or ssh to mysql service inside docker
    $ docker exec -it mysql-db /bin/bash

##### GET: http://localhost:9000/api/book
##### GET: http://localhost:9000/api/book/1
##### POST: http://localhost:9000/api/book -> Use PostMan to insert recore into DB