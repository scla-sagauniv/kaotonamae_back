FROM golang:1.22.4

RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY . .

RUN go mod tidy
CMD ["air"]