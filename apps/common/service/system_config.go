package commonservice

import (
	"lazyfury.github.com/yoomall-server/apps/common/model"
	"lazyfury.github.com/yoomall-server/core/driver"
)

type SystemConfigService struct {
	DB *driver.DB
}

func NewSystemConfigService(db *driver.DB) *SystemConfigService {
	return &SystemConfigService{
		DB: db,
	}
}

// get config by group as map
func (s *SystemConfigService) GetConfigByGroup(groupId uint64) map[string]string {
	var configs []model.SystemConfig
	s.DB.Model(&model.SystemConfig{}).Where("groupId = ?", groupId).Find(&configs)
	res := make(map[string]string)
	for _, config := range configs {
		res[config.Key] = config.Value
	}
	return res
}

// update config by group
func (s *SystemConfigService) UpdateConfigByGroup(groupId uint, configs []map[string]string) error {
	tx := s.DB.Begin()
	for _, config := range configs {
		tx.Model(&model.SystemConfig{}).Where("groupId = ? AND key = ?", groupId, config["key"]).Update("value", config["value"])
	}
	return tx.Commit().Error
}
