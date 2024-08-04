FROM golang:1.22

WORKDIR /app

COPY go.mod go.mod

COPY go.sum go.sum

RUN go mod tidy

COPY . .

RUN go build *.go

ENTRYPOINT [ "./main" ]