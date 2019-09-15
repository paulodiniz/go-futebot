FROM golang:latest

ADD . /go/src/myapp
WORKDIR /go/src/myapp

COPY . .
RUN go get
RUN go install
RUN go build -o main .
CMD ["./main"]