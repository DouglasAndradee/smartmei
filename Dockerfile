FROM golang:1.15

WORKDIR $GOPATH/src/smartmei

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

EXPOSE 5000

CMD ["smartmei"]