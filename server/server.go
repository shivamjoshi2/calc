package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	calc "github.com/shivamjoshi2/calc/calc"
	grpc "google.golang.org/grpc"
)

type server struct {
	calc.UnimplementedCalculatorServer
}

func (*server) TwoSum(ctx context.Context, req *calc.TwoNumRequest) (resp *calc.SumResponse, err error) {
	fmt.Println("Sum function is invoked for unary communication...")
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	resp = &calc.SumResponse{
		Sum: num1 + num2,
	}
	return resp, nil
}

func (*server) PrimeNums(req *calc.NumRequest, resp calc.Calculator_PrimeNumbersServer) error {
	fmt.Println("PrimeNumbers function is invoked for server side communication...")
	num := req.GetNum()
	var i int64
	for i = 2; i < num; i++ {
		var val int64 = i
		var check bool = true
		if val <= 1 {
			check = false
		}

		var j int64
		for j = 2; j < val; j++ {
			if val%int64(j) == 0 {
				check = false
			}
		}

		if check {
			res := &calc.AllPrimesResponse{
				Num: i,
			}

			// Delay to notice responses
			time.Sleep(1 * time.Second)
			resp.Send(&res)
		}
	}
	return nil
}

func (*server) Average(stream calc.Calculator_ComputeAverageServer) error {
	fmt.Println("ComputeAverage function is invoked for client side streaming...")
	var sum int64
	var nums int64
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calc.AverageResponse{
				Num: float32(sum) / float32(nums),
			})
		}
		if err != nil {
			log.Fatal(err)
		}
		sum += msg.GetNum()
		nums++
	}
}

func (*server) MaxNumber(stream calc.Calculator_FindMaxNumberServer) error {
	fmt.Println("FindMaxNumber function is invoked for bi direction streaming...")
	var ma int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
			return err
		}
		num := req.GetNum()
		if num > ma {
			ma = num
		}
		stream.Send(&calc.MaxNumResponse{
			Num: ma,
		})
	}
}
func main() {
	fmt.Println("vim-go")

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to Listen: %v", err)
	}

	s := grpc.NewServer()
	calc.RegisterCalculatorServer(s, &server{})

	if err = s.Serve(listen); err != nil {
		log.Fatal("failed to serve : %v", err)
	}
}
