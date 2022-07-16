

FROM golang:1.18-alpine
WORKDIR /usr/src/app
ADD go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /zerodha
EXPOSE 8080

CMD [ "/zerodha" ]
