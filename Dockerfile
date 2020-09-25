# Build backend
FROM golang:1.15.1-alpine3.12 AS builder
WORKDIR /go/src/github.com/hypertrace/demo-events-app
COPY . .
COPY ./events.json /go/bin/
RUN go build -o /go/bin/events-app ./

# Copy backend to actual image
FROM alpine:3.12.0
WORKDIR /go/bin
COPY --from=builder /go/bin/events-app .
COPY --from=builder /go/bin/events.json .
CMD [ "./events-app", "start-all" ]

