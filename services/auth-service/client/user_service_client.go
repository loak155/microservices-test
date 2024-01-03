package client

import "bufio"

type IUserServiceClient interface {
	GetUser(id int) (domain.User, error)
}

type userServiceClient struct {
	scanner *bufio.Scanner
	client  pb.UserServiceClient
}

func NewUserServiceClient() IUserServiceClient {
	return &userServiceClient{}
}

func (c *IUserServiceClient) GetUser() (User, error) {

	return User{}, nil
}
