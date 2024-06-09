FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "main.go"]