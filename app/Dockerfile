FROM golang:alpine as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

RUN go install main

EXPOSE 8000


FROM scratch

WORKDIR /app

COPY --from=builder /app/main /app/

ENTRYPOINT [ "/app/main" ]


