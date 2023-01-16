package main

import (
	"api-gateway/discovery"
	"api-gateway/internal/rpc"
	"api-gateway/internal/service"
	"api-gateway/routes"
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	InitConfig()
	go startListen() //转载路由
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		fmt.Println("exit! ", s)
	}
	fmt.Println("gateway listen on :3000")
}

func startListen() {
	InitRpc()

	ginRouter := routes.NewRouter()
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("server start err: %s", err.Error())
	}
}

func InitRpc() {
	// etcd注册
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// 各个RPC服务连接
	userServiceName := viper.GetString("domain.user")
	connUser, err := RPCConnect(ctx, userServiceName, etcdRegister)
	if err != nil {
		return
	}
	rpc.InitUserRpcRepo(service.NewUserServiceClient(connUser))

	dormitoryServiceName := viper.GetString("domain.dormitory")
	connDormitory, err := RPCConnect(ctx, dormitoryServiceName, etcdRegister)
	if err != nil {
		return
	}
	rpc.InitDormitoryRpcRepo(service.NewDormitoryServiceClient(connDormitory))
}

func RPCConnect(ctx context.Context, serviceName string, etcdRegister *discovery.Resolver) (conn *grpc.ClientConn, err error) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	addr := fmt.Sprintf("%s:///%s", etcdRegister.Scheme(), serviceName)
	conn, err = grpc.DialContext(ctx, addr, opts...)
	return
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
