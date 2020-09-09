# Build backend
FROM golang:1.15.1-alpine3.12 AS builder
WORKDIR /go/src/github.com/hypertrace/demo-events-app
COPY . .
RUN go build -o /go/bin/events-app ./cmd/main.go

# Copy backend to actual image
FROM alpine:3.12.0
WORKDIR /go/bin
COPY --from=builder /go/bin/events-app .
CMD [ "./events-app" ]
