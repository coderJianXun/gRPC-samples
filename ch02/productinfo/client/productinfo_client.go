package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "productinfo/client/ecommerce"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure()) // 创建到服务端的连接
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn) // 传递连接并创建存根文件，这个实例可以调用服务器的所有远程方法

	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // 创建 Context 以传递给远程调用
	defer cancel()
	r, err := c.AddProduct(ctx,
		&pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s", product.String())
}
