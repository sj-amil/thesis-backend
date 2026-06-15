package models

import "time"

type Quotation struct {
	ID         uint `gorm:"primaryKey;autoIncrement" json:"id"`
	RFQID      uint `gorm:"not null;uniqueIndex:idx_rfq_supplier" json:"rfq_id"`
	SupplierID uint `gorm:"not null;uniqueIndex:idx_rfq_supplier" json:"supplier_id"`

	BasePrice  float64 `gorm:"type:decimal(12,2);not null" json:"base_price"`
	MarginPct  float64 `gorm:"type:decimal(5,2);default:0" json:"margin_pct"`
	FinalPrice float64 `gorm:"type:decimal(12,2);not null" json:"final_price"`

	LeadTimeDays *int `json:"lead_time_days"`
	MOQ          *int `json:"moq"`
	ValidityDays *int `json:"validity_days"`

	Status string `gorm:"size:50;default:'pending'" json:"status"`

	SubmittedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"submitted_at"`
	LockedAt    *time.Time `json:"locked_at"`
	LockedBy    *uint      `json:"locked_by"`

	RFQ           *RFQ             `gorm:"foreignKey:RFQID" json:"rfq,omitempty"`
	Supplier      *SupplierProfile `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	PurchaseOrder *PurchaseOrder   `gorm:"foreignKey:QuotationID" json:"purchase_order,omitempty"`
}