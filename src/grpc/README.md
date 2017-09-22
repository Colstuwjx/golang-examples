# grpc

### build

```
cd customer
protoc -I customer/ customer/customer.proto --go_out=plugins=grpc:customer
```

### run

```
# for server side
go run server.go
```

```
# for client side
go run client.go
```
