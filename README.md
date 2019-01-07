# GRPC Template

1. Generate grpc code

```
$cd grpc-template
$protoc structure/person.proto --go_out=plugins=grpc:.
```
will generate `person.pb.go` under directory `structure`.

2. Run grpc server

```
$go run server/main.go
```

3. Run grpc client

```
$go run client/main.go
```
