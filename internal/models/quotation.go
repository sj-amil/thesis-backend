package models

import (
	"time"

	"github.com/google/uuid"
)

// Quotation represents supplier responses to RFQ
type Quotation struct {
    ID            uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
    RFQID         uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_rfq_supplier" json:"rfq_id"`
    SupplierID    uuid.UUID  `gorm:"type:uuid;not null;uniqueIndex:idx_rfq_supplier" json:"supplier_id"`
    
    // Pricing
    BasePrice     float64    `gorm:"type:decimal(12,2);not null" json:"base_price"`
    MarginPct     float64    `gorm:"type:decimal(5,2);default:0" json:"margin_pct"`
    FinalPrice    float64    `gorm:"type:decimal(12,2);not null" json:"final_price"`
    
    // Terms
    LeadTimeDays  *int       `json:"lead_time_days"`
    MOQ           *int       `json:"moq"`
    ValidityDays  *int       `json:"validity_days"`
    
    // Status
    Status        string     `gorm:"size:50;default:'pending'" json:"status"`
    
    // Metadata
    SubmittedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"submitted_at"`
    LockedAt      *time.Time `json:"locked_at"`
    LockedBy      *uuid.UUID `gorm:"type:uuid" json:"locked_by"`
    
    // Relationships
    // RFQ           *RFQ       `gorm:"foreignKey:RFQID" json:"rfq,omitempty"`
    // Supplier      *Supplier  `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	// PurchaseOrder *PurchaseOrder `gorm:"foreignKey:QuotationID" json:"purchase_order,omitempty"`
}

