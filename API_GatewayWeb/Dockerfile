FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /build

FROM alpine:latest

COPY --from=builder /build /app/build

EXPOSE 8080

CMD [ "/app/build" ]