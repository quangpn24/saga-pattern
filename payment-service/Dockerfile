FROM golang:1.19-alpine as builder

WORKDIR /project

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/server ./cmd/main.go

FROM alpine
COPY --from=builder ./project/bin/server /server
COPY --from=builder ./project/migration /migration

CMD ["/server"]