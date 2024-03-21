FROM golang:1.21.5-bullseye AS build

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o email-service

FROM busybox:latest

WORKDIR /email-service

COPY --from=build /app/email-service .

COPY --from=build /app/.env .

CMD [ "./email-service" ]