FROM golang:1.21 as builder

WORKDIR /app

COPY . ./

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -o /datadoll

FROM alpine:latest as build-release

RUN apk add --no-cache \
    ca-certificates \
    # this is needed only if you want to use scp to copy later your pb_data locally
    openssh

WORKDIR /

COPY --from=builder /datadoll /datadoll

EXPOSE 8080

ENTRYPOINT ["/datadoll", "serve", "--http=0.0.0.0:8080"]