# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

    cd $GOPATH/src
    git clone <this repo>
    go get ./...
    revel run revel-golang-app

### Go to http://localhost:9000

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites
