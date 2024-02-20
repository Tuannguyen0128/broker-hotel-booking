package server

import (
	"broker-hotel-booking/internal/proto"
	"context"
	"time"
)

func (sv *server) GetAllAccount(ctx context.Context, req *proto.AccountRequest) (*proto.AccountResponse, error) {
	accounts := make([]*proto.Account, 1)
	accounts = append(accounts, &proto.Account{
		ID:           0,
		Fullname:     "Jack",
		Email:        "Jack@gmail.com",
		Password:     "123",
		CreatedAt:    time.Now().String(),
		UpdatedAt:    time.Now().String(),
		MerchantCode: "123",
	})
	res := &proto.AccountResponse{
		Accounts: accounts,
	}
	return res, nil
}
