FROM golang:1.20-alpine
WORKDIR /src
COPY . .

# env vars for configuring go build
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -v -o ./simple-budget main.go
ENTRYPOINT ["./simple-budget"]
EXPOSE 8080