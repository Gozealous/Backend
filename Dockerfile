FROM golang:1.17-alpine
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN pwd

RUN ls

COPY **/* ./

RUN go mod download

RUN ls -R

RUN go build -o /gozealous

CMD [ "/gozealous" ]