package global

import (
	"gin-web/config"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Config
	GVA_LOG    *zap.Logger
)
