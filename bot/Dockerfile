FROM golang:1.18

WORKDIR /api

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /api/bin/main .

ENTRYPOINT [ "/api/bin/main" ]