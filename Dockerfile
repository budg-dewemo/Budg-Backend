FROM golang:1.16-alpine as builder

WORKDIR /app
ARG opts
COPY go.mod ./
COPY go.sum ./
COPY app.env ./
COPY ./src ./src

RUN go mod download
RUN env ${opts} go build -o /budg ./src/main.go

FROM alpine:latest
COPY --from=builder /budg /app/budg
EXPOSE 8080
CMD [ "/app/budg" ]