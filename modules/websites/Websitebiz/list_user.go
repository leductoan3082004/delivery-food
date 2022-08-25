package Websitebiz

import (
	"context"
	"food_delivery/modules/websites/WebsiteModel"
)

type GetList interface {
	GetList(ctx context.Context, filter *WebsiteModel.Filter) ([]WebsiteModel.User, error)
}
type GetListBiz struct {
	List GetList
}

func NewGetListBiz(List GetList) *GetListBiz {
	return &GetListBiz{List: List}
}
func (biz *GetListBiz) DoGetList(ctx context.Context, filter *WebsiteModel.Filter) ([]WebsiteModel.User, error) {
	result, err := biz.List.GetList(ctx, filter)
	return result, err
}
