# Developer Workflow
Notes, config, and general workflow for navpi-go development

## clone [navpi-go](https://github.com/NAVCoin/navpi-go.git)
This is the main project repo for NavPi 2.0 "Kowhai"

    git clone https://github.com/NAVCoin/navpi-go.git

## set project WORKDIR
    /go/src/github.com/NAVCoin/navpi-go/app

## install main dependencies
    go get ./

## install test dependencies
    go get -t -v ./...

## run app
    go run main.go

## run tests
    go test ./...

## clone [nav-docker](https://github.com/NAVCoin/nav-docker)
This repo contains Docker files used to build and run containerized instances of the different Nav projects.

    git clone https://github.com/NAVCoin/nav-docker.git

## start service
To run the service that spawns a Docker container running the navcoind daemon, simply run the command below in the directory containing the Docker files:

    cd nav-docker/docker-navcoind
    docker-compose up

## post-install
the install will take ages, but if you see the following you're good to go:

    Starting dockernavcoind_testnet_1 ... done
    Attaching to dockernavcoind_testnet_1

the navcoind daemon is now running on the testnet

## navcoin-cli access
Open up a new terminal tab/window

    docker exec -it dockernavcoind_testnet_1 /bin/bash

you should now be in cli mode with something like this:

    root@795b5c0525c0:/#

you will now be able to execute rpc commands accordingly

## testing endpoints
- Once the app is running it will be accessible at `127.0.0.1:9002`
- The default `managerApiPort` is `9002` and set in `server-config.json`

## wizard API
These endpoints are used for the NavPi setup wizard.

### setrange
takes the users ip address and saves it to the config as a range

    /api/setup/v1/setrange

### protectui
takes the api response and checks username and password

    /api/setup/v1/protectui

## daemon API
These endpoints are used in the NavPi UI.

### getstakereport
lists last single 30 day stake subtotal and last 24h, 7, 30, 365 day subtotal.

    /api/wallet/v1/getstakereport

### encryptwallet
Encrypts the wallet with _passphrase_. Once encrypted, three new commands are available:
_walletlock, walletpassphrase, walletpassphrasechange_

    /api/wallet/v1/encryptwallet





