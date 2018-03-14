FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/NAVCoin/navpi-go/app
WORKDIR /go/src/github.com/NAVCoin/navpi-go/app

# Install dependencies
RUN go get ./

# Install test dependencies
RUN go get -t -v ./...

# Build
RUN go build

# Test
RUN go test

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8080