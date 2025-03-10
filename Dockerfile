FROM golang:1.23.2
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd
EXPOSE 5005
CMD ["./api"]
