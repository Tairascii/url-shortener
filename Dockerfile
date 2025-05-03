FROM golang:1.23.4-alpine AS builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/server/main.go

FROM scratch

COPY --from=builder server /bin/server
COPY config/values_docker.yaml /bin/config/values_docker.yaml

ENV CONFIG_FILE=/bin/config/values_docker.yaml

ENTRYPOINT ["/bin/server"]
