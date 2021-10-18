package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	pb "ordermgt/server/ecommerce"
	"strings"
	"time"
)

const (
	port           = ":50051"
	orderBatchSize = 3
)

var orderMap = make(map[string]pb.Order)

type server struct {
	orderMap map[string]*pb.Order
}

func (s *server) ProcessOrders(stream pb.OrderManagement_ProcessOrdersServer) error {
	batchMarker := 1
	var combinedShipmentMap = make(map[string]pb.CombinedShipment)
	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order ; %s", orderId)
		if err == io.EOF {
			log.Printf("EOF : %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}
		if err != nil {
			log.Println(err)
			return err
		}

		destination := orderMap[orderId.GetValue()].Destination
		shipment, found := combinedShipmentMap[destination]

		if found {
			ord := orderMap[orderId.GetValue()]
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = shipment
		} else {
			comShip := pb.CombinedShipment{Id: "cmb - " + (orderMap[orderId.GetValue()].Description), Status: "Processed!"}
			ord := orderMap[orderId.GetValue()]
			comShip.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[destination] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping : %v -> %v", comb.Id, len(comb.OrdersList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]pb.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}

func (s *server) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	ord := orderMap[orderId.Value]
	return &ord, nil
}

func (s *server) SearchOrders(searchQuery *wrappers.StringValue, strem pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		log.Print(key, order)
		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				err := strem.Send(&order) // 在流中发送匹配的订单
				if err != nil {
					return fmt.Errorf("error sending message to stream: %v", err)
				}
				log.Print("Matching Order Found: ", key)
				break
			}
		}
	}
	return nil
}

func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {
	ordersStr := "Update Order IDs : "
	for {
		order, err := stream.Recv() // 从客户端流中读取消息
		if err == io.EOF {          // 检查流是否已经结束
			return stream.SendAndClose(&wrappers.StringValue{Value: "Orders processed " + ordersStr}) // 服务端发送响应
		}
		orderMap[order.Id] = *order

		log.Println("Order ID ", order.Id, ": Updated")
		ordersStr += order.Id + ","
	}
}

// 一元拦截器 func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler)
// 可以控制传入 gRPC 服务器端的所有一元 RPC
func orderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 前置处理逻辑，可以拦截处理
	// 检查传入的参数，获取当前 RPC 的信息
	log.Println("========= [Server Interceptor] ", info.FullMethod)

	// 调用 UnaryHandler 完成一元 RPC 的正常执行
	m, err := handler(ctx, req)

	// 后置处理逻辑，这里可以处理响应信息
	log.Printf(" Post Proc Message: %s", m)
	return m, err // 发送响应
}

// 服务器流拦截器
// wrappedStream 包装嵌入的 grpc.ServerStream，并拦截对 RecvMsg 函数和 SendMsg 函数的调用
type wrappedStream struct { // grpc.ServerStream 的包装器流
	grpc.ServerStream
}

// RecvMsg 处理流接收到的消息
func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("======= [Server Stream Interceptor Wrapper] "+"Receive a message (Type: %T) at %s", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

// SendMsg 处理流发送的消息
func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("======= [Server Stream Interceptor Wrapper] "+" Send a message (Type: %T) at %v", m,
		time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

// 创建新包装器流的实例
func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

// 流拦截器的实现
func orderServerStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("======= [Server Stream Interceptor] ", info.FullMethod)
	err := handler(srv, newWrappedStream(ss))
	if err != nil {
		log.Printf("RPC failed with error %v", err)
	}
	return err
}

func main() {
	initSampleData()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(orderUnaryServerInterceptor),   // 注册一元拦截器
		grpc.StreamInterceptor(orderServerStreamInterceptor), // 注册流拦截器
	)
	pb.RegisterOrderManagementServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}
