FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./
COPY config.deploy.yml ./config.yml

RUN go build -o ./url-counter

CMD ["./url-counter"]
