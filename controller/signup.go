package controller

import (
	"context"
	signuppb "craftsman-api-gateway/pkg/grpc"
	"fmt"
)

type NewMemberRegistrationService struct {
	signuppb.UnimplementedNewMemberRegistrationServiceServer
}

func NewNewMemberRegistrationService() *NewMemberRegistrationService {
	return &NewMemberRegistrationService{}
}

func (n *NewMemberRegistrationService) Signup(ctx context.Context, req *signuppb.SignupRequest) (*signuppb.SignupResponse, error) {
	fmt.Printf("req.UserName : %v", req.UserName)
	fmt.Println()
	fmt.Printf("req.GetUserName() : %v", req.GetUserName())

	fmt.Println()
	fmt.Println()

	fmt.Printf("req.MailAddress : %v", req.MailAddress)
	fmt.Println()
	fmt.Printf("req.GetMailAddress() : %v", req.GetMailAddress())

	fmt.Println()
	fmt.Println()

	fmt.Printf("req.Password : %v", req.Password)
	fmt.Println()
	fmt.Printf("req.GetPassword() : %v", req.GetPassword())
	fmt.Println()
	fmt.Println()

	return &signuppb.SignupResponse{
		ID:          "1",
		UserName:    "unkoburi",
		MailAddress: "inogan38@gmail.com",
	}, nil

}
