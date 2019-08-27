# Telling to use Docker's golang ready image
FROM golang

MAINTAINER abinhho <abinh.ho@gmail.com>

EXPOSE 9000

# Set pwd to the go folder
RUN export GOPATH=/go
WORKDIR ${GOPATH}

RUN apt-get update -y && \
   apt-get install -y \
   vim

# Install and build binaries revel for this server
RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN go get github.com/go-gorp/gorp

# Copy the entrypoint file
# ADD /docker/app/infinite_loop.sh /infinite_loop.sh
# RUN chmod +x /infinite_loop.sh

# Start process
# ENTRYPOINT ["/infinite_loop.sh"]

RUN ls /go/src

# RUN cp /go/conf/app.conf.dev-docker /go/conf/app.conf
# RUN revel run revel-golang-app
