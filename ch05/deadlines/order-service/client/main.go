package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "ordermgt/client/ecommerce"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	orderMgtClient := pb.NewOrderManagementClient(conn)

	// add deadline
	clientDeadline := time.Now().Add(time.Duration(2 * time.Second)) // 2秒截止时间
	ctx, cancel := context.WithDeadline(context.Background(),clientDeadline)
	defer cancel()

	// 添加订单
	order1 := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"},Destination: "San Jose, CA", Price: 2300.00}
	res, addErr := orderMgtClient.AddOrder(ctx, &order1)
	if addErr != nil {
		 got := status.Code(addErr)
		 log.Printf("Error Occured -> addOrder : , %v",got)
	} else {
		 log.Print("AddOrder Response -> ", res.Value)
	}

	// 获取订单
	retrievedOrder, err := orderMgtClient.GetOrder(ctx, &wrappers.StringValue{Value: "106"})
	log.Print("GetOrder Response -> : ", retrievedOrder)

	searchStream, _ := orderMgtClient.SearchOrders(ctx, &wrappers.StringValue{Value: "Google"})
	for {
		searchOrder, err := searchStream.Recv()
		if err == io.EOF {
			break
		}
		log.Print("Search Result: ", searchOrder)
	}

	// updateOrders
	updOrder1 := pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "Mountain View, CA", Price: 1100.00}
	updOrder2 := pb.Order{Id: "103", Items: []string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination: "San Jose, CA", Price: 2800.00}
	updOrder3 := pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Mountain View, CA", Price: 2200.00}

	updateStream, err := orderMgtClient.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("%v.UpdateOrders(_) = _, %v", orderMgtClient, err)
	}

	// updateOrders 1
	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	}

	// updateOrders 2
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	}

	// updateOrders 3
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder3, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}
	log.Printf("Update Orders Res : %s", updateRes)

	// 处理订单
	streamProcOrder, err := orderMgtClient.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("%v.ProcessOrders(_) = _, %v", orderMgtClient, err)
	}
	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "102"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", orderMgtClient, "102", err)
	}

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "103"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", orderMgtClient, "103", err)
	}

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "104"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", orderMgtClient, "104", err)
	}

	channel := make(chan struct{})
	go asyncClientBidirectionalRPC(streamProcOrder, channel)
	time.Sleep(time.Millisecond * 1000)

	if err := streamProcOrder.Send(&wrappers.StringValue{Value: "101"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", orderMgtClient, "101", err)
	}

	if err := streamProcOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}
	channel <- struct{}{}
}

func asyncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan struct{}) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder == io.EOF {
			break
		}
		log.Print("Combined shipment : ", combinedShipment.OrdersList)
	}
	<-c
}
