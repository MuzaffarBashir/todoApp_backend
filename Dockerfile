FROM golang:1.17-alpine

ENV WORKDIR=/app
WORKDIR $WORKDIR

COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download
COPY src $WORKDIR

RUN go build -o todoapp-go .
EXPOSE 8090
CMD ["./todoapp-go"]