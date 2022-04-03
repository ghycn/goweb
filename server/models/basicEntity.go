package models

import "time"

type BasicEntity struct {
	ID        int        `gorm:"primary_key" json:"id"`
	TenantID  int        `gorm:"tenant_id" json:"tenantId"`
	CreatedAt time.Time  `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"deleted_at" json:"deletedAt"`
}
