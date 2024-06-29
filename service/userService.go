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

func (S *Server) UpdateUser(ctx context.Context, user *pb.User) (*pb.Status, error) {
	status, err := S.U.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return status, err
}

func (S *Server) DeleteUser(ctx context.Context, userId *pb.UserId) (*pb.Status, error) {
	status, err := S.U.DeleteUser(userId)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (S *Server) GetUserProfile(ctx context.Context, userId *pb.UserId) (*pb.UserProfile, error) {
	userProfil, err := S.U.GetUserProfile(userId)
	if err != nil {
		return nil, err
	}
	return userProfil, nil
}

func (S *Server) UpdateUserProfile(ctx context.Context, userProfile *pb.UserProfile) (*pb.Status, error) {
	status, err := S.U.UpdateUserProfile(userProfile)
	if err != nil {
		return nil, err
	}
	return status, nil
}
