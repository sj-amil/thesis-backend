package models

import (
	"time"
)

// SupplierProfile stores manufacturer/factory data for suppliers.
// Includes verification status and rating for marketplace trust scoring.
type SupplierProfile struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	// UserID         uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_suppliers_user;not null" json:"user_id"`
	CompanyName    string    `gorm:"size:255;not null" json:"company_name"`
	ContactInfo    string    `gorm:"size:255" json:"contact_info"`
	Address        string    `gorm:"type:text" json:"address"`
	ComplianceDocs string    `gorm:"size:255" json:"compliance_docs_url,omitempty"`
	IsVerified     bool      `gorm:"default:false" json:"is_verified"`
	Rating         float64   `gorm:"type:decimal(3,2);default:0" json:"rating"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relations
	// User       *User       `gorm:"foreignKey:UserID" json:"-"`
	// Quotations []Quotation `gorm:"foreignKey:SupplierID" json:"quotations,omitempty"`
}