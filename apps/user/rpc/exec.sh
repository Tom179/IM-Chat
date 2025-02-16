# go zero命令
goctl rpc protoc ./apps/user/rpc/user.proto --go_out=./apps/user/rpc/ --go-grpc_out=./apps/user/rpc --zrpc_out=./apps/user/rpc/
# --go_out=指定pb.go文件的位置
# --grpc_out=指定grpc.pb.go位置
# --zgrpc= gozero框架生成的相关代码

goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models/" -c
# go-zero生成mysql模型

goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero