package main

import (
	"log"
	"net"
	pb "userService/generated"
	"userService/service"
	"userService/storage"
	"userService/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.Connect()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()

	s := service.NewUserServer(*postgres.NewUserRepo(db))
	grpc := grpc.NewServer()
	pb.RegisterUserServiceServer(grpc, s)
	if err = grpc.Serve(listener); err != nil{
		log.Fatal(err)
	}
}
