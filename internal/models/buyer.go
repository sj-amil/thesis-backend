package models

import "time"

type BuyerProfile struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint      `gorm:"uniqueIndex:idx_buyers_user;not null" json:"user_id"`
	CompanyName  string    `gorm:"size:255;not null" json:"company_name"`
	ContactInfo  string    `gorm:"size:255" json:"contact_info"`
	TradeLicense string    `gorm:"size:255" json:"trade_license_url,omitempty"`
	IsVerified   bool      `gorm:"default:false" json:"is_verified"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations
	User *User `gorm:"foreignKey:UserID" json:"-"`
	RFQs []RFQ `gorm:"foreignKey:BuyerID" json:"rfqs,omitempty"`
}