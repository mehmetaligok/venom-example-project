package grpc

import (
	"context"

	"github.com/google/uuid"

	"github.com/mehmetaligok/venom-example-project/src/model"
	pb "github.com/mehmetaligok/venom-example-project/src/proto/user"
)

type UserRepo interface {
	InsertUser(ctx context.Context, user *model.User) error
}

type UserServer struct {
	userRepo UserRepo
}

// NewUserServer returns new user server instance
func NewUserServer(repo UserRepo) *UserServer {
	return &UserServer{userRepo: repo}
}

// AddUser add users to database
func (server *UserServer) AddUser(ctx context.Context, request *pb.NewUserRequest) (*pb.NewUserResponse, error) {
	user := &model.User{
		ID:        uuid.New(),
		FirstName: request.GetFirstName(),
		LastName:  request.GetLastName(),
	}
	err := server.userRepo.InsertUser(ctx, user)

	if err != nil {
		return &pb.NewUserResponse{Status: pb.ResponseStatus_RESPONSE_STATUS_REJECTED}, err
	}

	return &pb.NewUserResponse{
		Status: pb.ResponseStatus_RESPONSE_STATUS_CONFIRMED,
		Id:     user.ID.String(),
	}, nil
}
