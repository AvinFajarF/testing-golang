package service

import (
	"context"
	"errors"

	"github.com/AvinFajarF/config"
	"github.com/AvinFajarF/model"
	pb "github.com/AvinFajarF/proto"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func (*UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := req.GetUser()

	data := model.UserModel{
		ID: user.GetId(),
		Name: user.GetName(),
		Email: user.GetEmail(),
	}

	config.DB.Create(&data)

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id: user.GetId(),
            Name: user.GetName(),
            Email: user.GetEmail(),
		},
	}, nil

}

func (*UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var users = []*pb.User{}

	result := config.DB.Find(&users)


	if result.Error != nil {
		return nil, errors.New("failed get all users")
	}

	return &pb.GetUserResponse{
		User: users,
	},nil

}

func (*UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var user model.UserModel
	args := req.GetUser()

	result := config.DB.Model(&user).Where("id = ?", args.Id).Updates(model.UserModel{
		Name: args.Name,
        Email: args.Email,
	})

	if result.Error != nil {
		return nil, errors.New("failed get all users")
	}

	return &pb.UpdateUserResponse{
		User: &pb.User{
            Id: args.Id,
            Name: args.Name,
            Email: args.Email,
        },
	}, nil


}


func (*UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	var user model.UserModel
	args := req.Id

	result := config.DB.Where("id = ?", args).Delete(&user)

	if result.Error != nil {
		return nil, errors.New("failed delete user")
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil


}