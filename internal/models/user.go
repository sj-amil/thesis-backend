package models

import (
	"time"

	"github.com/google/uuid"
)

// User is the core authentication entity.
// Role-based access is enforced via the 'role' field.
type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Email        string    `gorm:"size:255;uniqueIndex:idx_users_email;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"` // Never exposed in API
	Role         string    `gorm:"size:50;not null;index;check:role IN ('buyer','supplier','admin')" json:"role"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// 1:1 Profile Relations (loaded only when requested)
	// BuyerProfile    *BuyerProfile    `gorm:"foreignKey:UserID" json:"buyer_profile,omitempty"`
	// SupplierProfile *SupplierProfile `gorm:"foreignKey:UserID" json:"supplier_profile,omitempty"`
}
