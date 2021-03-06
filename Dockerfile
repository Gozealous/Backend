FROM golang:1.17-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./


RUN go mod download
COPY . ./

RUN ls -R

RUN go build -o /gozealous

CMD [ "/gozealous" ]