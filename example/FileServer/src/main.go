package main

import (
	"net"
	"github.com/TriangleGo/gylib/logger"
	"github.com/TriangleGo/gylib/cache"
	"github.com/TriangleGo/gylib/mongo"
	"github.com/TriangleGo/gylib/service/proto"
	"github.com/TriangleGo/gylib/service/etcd"
	"service"
	"google.golang.org/grpc"
	"fmt"
)

func init() {
	logger.InitLogger("./conf/logger.conf")
	cache.InitCache("./conf/cache.conf")
	mongo.InitMongo("./conf/mongo.conf")
	etcd.InitEtcd("./conf/etcd.conf")
}

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", etcd.ServicePort))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	} else {
		logger.Debugf("listening on port %s", etcd.ServicePort)
	}

	etcd.DeleteService()
	etcd.RegisterService()
	// 创建grpc实例
	grpcServer := grpc.NewServer()
	// 注册fileService服务
	proto.RegisterServiceServer(grpcServer, service.InitServer())
	// 启动服务
	grpcServer.Serve(lis)
}