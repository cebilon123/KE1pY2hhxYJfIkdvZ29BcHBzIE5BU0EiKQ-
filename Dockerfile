FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY ./ /app/
RUN go mod download

COPY *.go ./

RUN go build -o /api

EXPOSE 8080

CMD [ "/api" ]