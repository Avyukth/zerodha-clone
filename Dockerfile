

# FROM golang:1.18-alpine
# WORKDIR /app
# ADD go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o bin/zerodha -v .
# EXPOSE 8080

# CMD [ "bin/zerodha" ]



#first stage - builder
FROM golang:1.18-alpine AS build

WORKDIR /

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . . ./

RUN go build -o /zerodha

##
## Deploy
##
# FROM gcr.io/distroless/base-debian10

# WORKDIR /
# # ARG BUILD_CONFIG=ls
# COPY --from=build /zerodha /
# RUN echo ${BUILD_CONFIG}
# EXPOSE 8080

# USER nonroot:nonroot
ENTRYPOINT ["/zerodha"]
