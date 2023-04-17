# Use an official Go runtime as a base image
FROM golang:1.20.2-alpine

RUN go version
ENV GOPATH=/

COPY . .

RUN go mod download

RUN go build -o employeeService ./cmd/main.go

EXPOSE 8000

CMD ["./employeeService"]