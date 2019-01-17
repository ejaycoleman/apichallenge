FROM golang:1.8

WORKDIR /go/src/app
COPY ./main/facade.go .

CMD ["go", "run", "hmm.go"]