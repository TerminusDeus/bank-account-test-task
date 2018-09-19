FROM golang:1.8.5-jessie
RUN go get -u github.com/gin-gonic/gin
WORKDIR /go/src/app
# add source code
ADD src src
# run main.go
CMD ["go", "run", "src/main.go"]