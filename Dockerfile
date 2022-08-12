# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
ENV WEATHER_CLIENT_TOKEN=f992337d1f4d40218cb153808220808
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /docker-forecast-cities ./src

CMD [ "/docker-forecast-cities" ]
