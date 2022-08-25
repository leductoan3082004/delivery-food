package WebsiteStorage

import (
	"context"
	"food_delivery/modules/websites/WebsiteModel"
)

func (s *sqlStore) Create(ctx context.Context, data *WebsiteModel.User) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
