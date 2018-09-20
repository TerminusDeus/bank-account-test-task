FROM golang:latest

RUN go get -u github.com/gin-gonic/gin

WORKDIR /go/src/bank-account-test-task
ADD src /go/src/bank-account-test-task/src
CMD ["go", "run", "src/main.go"]

EXPOSE 8080