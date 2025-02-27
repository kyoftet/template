####################### Build stage #######################
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main /app/cmd/api/main.go


####################### Run stage #######################
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
CMD [ "./main" ]
