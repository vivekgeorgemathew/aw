## 1. Build the Binary
FROM golang:1.22-alpine AS builder

WORKDIR app

COPY ./ ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o aw
EXPOSE 8080

## 2. Build the Image
FROM scratch
WORKDIR /

COPY --from=builder /go/app/app.env /app.env
COPY --from=builder /go/app/aw /aw

ENTRYPOINT ["/aw"]




