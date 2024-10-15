FROM golang:1.23

WORKDIR /app

COPY ./api/go/src/go.sum ./api/go/src/go.mod ./
RUN go mod download && go mod verify

CMD ["go", "run", "."]