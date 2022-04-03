package models

import "time"

type BasicEntity struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	TenantID  uint       `gorm:"tenant_id" json:"tenantId"`
	CreatedAt time.Time  `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"deleted_at" json:"deletedAt"`
}
