FROM golang:1.8

COPY . /go/src/github.com/Simulalex/fortune-cookie-api
COPY fortunes.json .

RUN go install github.com/Simulalex/fortune-cookie-api

ENTRYPOINT /go/bin/fortune-cookie-api

EXPOSE 8080
