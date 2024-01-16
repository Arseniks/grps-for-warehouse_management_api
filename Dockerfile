FROM golang:latest

RUN go version
ENV GOPATH=/

COPY . .

RUN go mod download
RUN go build -o jsonrpc_warehouse_management_api ./cmd/root/main.go

CMD ["./jsonrpc_warehouse_management_api"]
