# tag-service-grpc

基础grpc服务应用

**grpc包获取**

```bash
go get google.golang.org/grpc
```

**grpcurl安装**
```bash
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
.exe文件放入 /go/bin文件夹中(环境变量)


**grpcui安装**

```bash
go get github.com/fullstorydev/grpcui
go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```
.exe文件放入 /go/bin文件夹中

**cmux**

在同端口号同时监听 实现多协议支持
```
go get -u github.com/soheilhy/cmux

```

**grpc-gateway**

将 RESTful 转换为 gRPC 请求
```bash
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

.exe文件放入 /go/bin文件夹中