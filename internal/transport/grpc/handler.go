package grpc

import (
	"context"

	userpb "github.com/JustSelection/project-protos/proto/users"
	"github.com/JustSelection/users-service/internal/user"
)

type Handler struct {
	svc user.Service
	userpb.UnimplementedUserServiceServer
}

func NewHandler(svc user.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := h.svc.CreateUser(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.CreateUserResponse{
		User: &userpb.User{
			Id:    u.ID,
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	u, err := h.svc.GetUserByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserResponse{
		User: &userpb.User{
			Id:    u.ID,
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	u, err := h.svc.UpdateUser(req.Id, &req.Email, &req.Password)
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: &userpb.User{
			Id:    u.ID,
			Email: u.Email,
		},
	}, nil
}

func (h *Handler) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	err := h.svc.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &userpb.DeleteUserResponse{}, nil
}

func (h *Handler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	users, err := h.svc.GetAllUsers()
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*userpb.User, len(users))
	for i, u := range users {
		pbUsers[i] = &userpb.User{
			Id:    u.ID,
			Email: u.Email,
		}
	}

	return &userpb.ListUsersResponse{Users: pbUsers}, nil
}
