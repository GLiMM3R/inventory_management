FROM golang:1.22.5 AS builder
WORKDIR /app
COPY . .
COPY go.mod go.sum .
RUN go mod download
# RUN go build -o ./inventory_management ./main.go
RUN COPY *.go ./ 
RUN CGO_ENABLED=0 GOOS=linux go build -o /inventory_management /cmd/api/main.go
 
 
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/inventory_management .
EXPOSE 8080
ENTRYPOINT ["./inventory_management"]