FROM golang:1.18-alpine as builder

WORKDIR /app
ARG opts
COPY go.mod ./
COPY go.sum ./
COPY ./src ./src

RUN go mod download
RUN env ${opts} go build -o /budg ./src/main.go

FROM alpine:latest
COPY --from=builder /budg /app/bin/budg
COPY docker.env /app.env
EXPOSE 8080
CMD [ "/app/bin/budg" ]
