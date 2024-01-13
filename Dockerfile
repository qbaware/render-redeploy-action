FROM golang:1.21

WORKDIR /app

COPY ./src/go.mod ./src/go.sum ./

RUN go mod download

COPY ./src/ .

RUN go build -o app ./main.go

ENTRYPOINT [ "./app" ] 
