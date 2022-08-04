FROM golang:1.18-alpine
RUN apk add --no-cache git
RUN apk add --no-cache bash
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY package.json .
COPY package-lock.json .
RUN npm i

COPY ./ ./

RUN go build -o ./main main.go

CMD ["./main"]