FROM golang:1.17-alpine
WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

COPY * ./

RUN go mod download

RUN ls

RUN go build -o /gozealous

CMD [ "/gozealous" ]