FROM golang:1.17-alpine AS builder

WORKDIR /build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o sendsms .

FROM scratch

COPY --from=builder ["/build/sendsms", "/"]

EXPOSE 3005

ENTRYPOINT ["/sendsms"]