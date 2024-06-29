package service

import (
	"context"
	pb "userService/generated"
	"userService/storage/postgres"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	U postgres.NewUser
}

func NewUserServer(u postgres.NewUser) *Server {
	return &Server{U: u}
}

func (S *Server) GetUser(ctx context.Context, userId *pb.UserId) (*pb.User, error) {
	user, err := S.U.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
