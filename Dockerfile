FROM golang:1.20.5-alpine3.17

WORKDIR /app
RUN echo "" > /app/.env

COPY . /app

EXPOSE 5000

CMD ["go","run","main.go"]
