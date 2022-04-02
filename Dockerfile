FROM golang:1.18-buster AS builder

WORKDIR $GOPATH/src/github.com/nazliander/simple-go-actions/
COPY . .

RUN go get .

RUN CGO_ENABLED=0 go build -o /main

FROM scratch

COPY --from=builder /main /main

EXPOSE 8080

CMD ["./main"]
