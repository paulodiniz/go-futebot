FROM golang:latest

ADD . /go/src/myapp
WORKDIR /go/src/myapp

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .
CMD ["./main"]