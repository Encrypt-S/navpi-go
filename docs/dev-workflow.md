### Developer Workflow
Notes, config, and general workflow for navpi-go development

#### clone [navpi-go](https://github.com/NAVCoin/navpi-go.git)
This is the main project repo for NavPi 2.0 "Kowhai"

    git clone https://github.com/NAVCoin/navpi-go.git

#### set project WORKDIR
    /go/src/github.com/NAVCoin/navpi-go/app

#### install main dependencies
    go get ./

#### install test dependencies
    go get -t -v ./...

#### run app
    go run main.go

#### run tests
    go test ./...

#### clone [nav-docker](https://github.com/NAVCoin/nav-docker)
This repo contains Docker files used to build and run containerized instances of the different Nav projects.

    git clone https://github.com/NAVCoin/nav-docker.git

#### start service (navcoind daemon)
To run the service that spawns a Docker container running the navcoind daemon, simply run the command below in the directory containing the Docker files:

    cd nav-docker/docker-navcoind
    docker-compose up

#### local testing
- Once the app is running it will be accessible at `127.0.0.1:9002`
- The default `managerApiPort` is `9002` and set in `server-config.json`

#### wizard api
> http://127.0.0.1:9002/api/setup/v1/setrange
  - takes the users ip address and saves it to the config as a range

> http://127.0.0.1:9002/api/setup/v1/protectui
  - takes the api response and checks username and password

#### daemon api
> http://127.0.0.1:9002/api/wallet/v1/getstakereport
    - writes out stake report





