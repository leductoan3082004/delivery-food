package WebsiteStorage

import (
	"context"
	"food_delivery/modules/websites/WebsiteModel"
)

func (s *sqlStore) GetList(ctx context.Context, filter *WebsiteModel.Filter) ([]WebsiteModel.User, error) {
	db := s.db
	db = db.Where(filter)
	var result []WebsiteModel.User
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
