FROM golang:1.21 as base

FROM base as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /app

EXPOSE 8090

CMD ["air"]