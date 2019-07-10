### server

``` sh
protoc -Iserver/todo/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis server/todo/todo.proto --go_out=plugins=grpc:server/todo
protoc -Iserver/todo/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis server/todo/todo.proto --js_out=import_style=commonjs:client/src/lib
protoc -Iserver/todo/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis server/todo/todo.proto --grpc-web_out=import_style=commonjs,mode=grpcwebtext:client/src/lib
protoc -Iserver/todo/ -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis server/todo/todo.proto --grpc-gateway_out=logtostderr=true:server/todo

cd server
# grpc server
go run main.go
# http server
go run gw.go
```

### envoy

``` sh
cd envoy
docker build -t envoy:v1 .
docker run  -p 8080:8080 -p 9090:9090 envoy:v1
```

### client

``` sh
cd client
npm run serve
```
