

FROM golang:1.8

WORKDIR /go/src/app
COPY ./main/hmm.go .

EXPOSE 8080/tcp
EXPOSE 8080/udp



CMD ["go", "run", "hmm.go"]