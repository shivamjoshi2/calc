package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	calc "github.com/shivamjoshi2/calc/calc"
	grpc "google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calc.NewCalculatorClient(cc)

	MaxNumber(c)
}

func TwoSum(c calc.CalcClient) {
	fmt.Println("Starting Sum service...")
	req := calc.TwoNumRequest{
		Num1: 10,
		Num2: 20,
	}
	resp, err := c.TwoSum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Sum grpc unary call: %v", err)
	}
	log.Printf("Response from Sum grpc unary call: %v", resp.Sum)
}

func PrimeNums(c calc.CalcClient) {
	fmt.Println("Starting Prime number service...")
	req := calc.NumRequest{
		Num: 10,
	}
	resStream, err := c.PrimeNums(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumbers server-side streaming grpc: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while receving server stream: %v", err)
		}
		fmt.Println("Response From PrimeNumbers Server: ", msg.Num)
	}

}

func Average(c calc.CalcClient) {
	fmt.Println("Starting ComputeAverage service...")
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client-side streaming : %v", err)
	}
	req := []*calc.NumRequest{
		{Num: 10},
		{Num: 20},
		{Num: 30},
		{Num: 40},
		{Num: 50},
		{Num: 60},
	}
	for _, val := range req {
		fmt.Println("\nSending numbers...: ", val)
		stream.Send(val)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server : %v", err)
	}
	fmt.Println("\n****Response From Server : ", resp.Num)

}

func MaxNumber(c calc.CalcClient) {
	fmt.Println("Starting Maxnumber service")
	req := []*calc.NumRequest{
		{Num: 1},
		{Num: 3},
		{Num: 7},
		{Num: 9},
		{Num: 2},
		{Num: 5},
		{Num: 22},
		{Num: 15},
		{Num: 21},
		{Num: 19},
	}
	stream, err := c.MaxNumber(context.Background())
	if err != nil {
		log.Fatalf("error occured while performing client side streaming : %v", err)
	}

	waitchan := make(chan int32)

	go func(req []*calc.NumRequest) {
		for _, val := range req {
			fmt.Println("\nSending numbers...: ", val.Num)
			err := stream.Send(val)
			if err != nil {
				log.Fatalf("error while sending request to FindMaxNumber service : %v", err)
			}
			time.Sleep(1000 * time.Millisecond)

		}
		stream.CloseSend()
	}(req)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(waitchan)
				return
			}
			if err != nil {
				log.Fatalf("Error receiving response from server : %v", err)
			}
			fmt.Printf("\nResponse From Server : %v", resp.Num)
		}
	}()
	<-waitchan
}
