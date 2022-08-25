package Websitebiz

import (
	"context"
	"food_delivery/modules/websites/WebsiteModel"
)

type CreateUser interface {
	Create(ctx context.Context, data *WebsiteModel.User) error
}
type CreateUserBiz struct {
	store CreateUser
}

func NewCreateUserBiz(store CreateUser) *CreateUserBiz {
	return &CreateUserBiz{store: store}
}
func (biz *CreateUserBiz) Register(ctx context.Context, data *WebsiteModel.User) error {
	err := biz.store.Create(ctx, data)
	return err
}
