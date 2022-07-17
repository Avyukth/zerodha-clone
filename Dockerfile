FROM golang:1.18-alpine AS builder
RUN mkdir -p /app
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o zerodha


FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/zerodha /app
COPY --from=builder /app/*.env /app
EXPOSE 8080
ENTRYPOINT ["./zerodha"]
