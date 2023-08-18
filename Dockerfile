FROM golang:1.19
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o build/Plugin-Go ./main.go
CMD ["go", "run", "./main.go"]