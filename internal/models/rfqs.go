package models

import (
	"time"

	"github.com/google/uuid"
)

// RFQ represents the main Request for Quote document
type RFQ struct {
    ID                  uuid.UUID      `gorm:"type:uuid;primary_key" json:"id"`
    RFQNumber           string         `gorm:"size:50;uniqueIndex;not null" json:"rfq_number"`
    BuyerID             uuid.UUID      `gorm:"type:uuid;not null" json:"buyer_id"`
    
    // Buyer Company Info
    BuyerCompanyName    string         `gorm:"size:255;not null" json:"buyer_company_name"`
    BuyerAddress        string         `gorm:"type:text;not null" json:"buyer_address"`
    BuyerContactInfo    string         `gorm:"size:255" json:"buyer_contact_info"`
    
    // Quote Recipient (Supplier)
    SupplierCompanyName string         `gorm:"size:255" json:"supplier_company_name"`
    SupplierAddress     string         `gorm:"type:text" json:"supplier_address"`
    SupplierContactInfo string         `gorm:"size:255" json:"supplier_contact_info"`
    
    // RFQ Details
    PONumber            string         `gorm:"size:100" json:"po_number"`
    ShipDate            *time.Time     `gorm:"type:date" json:"ship_date"`
    ShipVia             string         `gorm:"size:100" json:"ship_via"`
    FOBPoint            string         `gorm:"size:100" json:"fob_point"`
    PaymentTerms        string         `gorm:"size:100" json:"payment_terms"`
    QuoteValidUntil     *time.Time     `gorm:"type:date" json:"quote_valid_until"`
    
    // Status & Metadata
    Status              string         `gorm:"size:50;default:'draft'" json:"status"`
    Comments            string         `gorm:"type:text" json:"comments"`
    SpecialInstructions string         `gorm:"type:text" json:"special_instructions"`
    
    // Relationships
    Items               []RFQItem      `gorm:"foreignKey:RFQID" json:"items,omitempty"`
    // Quotations          []Quotation    `gorm:"foreignKey:RFQID" json:"quotations,omitempty"`
    
    // Timestamps
    CreatedAt           time.Time      `json:"created_at"`
    UpdatedAt           time.Time      `json:"updated_at"`
    CreatedBy           uuid.UUID      `gorm:"type:uuid" json:"created_by"`
}

// RFQItem represents line items in an RFQ
type RFQItem struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
    RFQID       uuid.UUID `gorm:"type:uuid;not null;index" json:"rfq_id"`
    LineNumber  int       `gorm:"not null" json:"line_number"`
    Quantity    float64   `gorm:"type:decimal(12,2);not null;default:1" json:"quantity"`
    Description string    `gorm:"type:text;not null" json:"description"`
    UnitPrice   *float64  `gorm:"type:decimal(12,2)" json:"unit_price"`
    Taxable     bool      `gorm:"default:false" json:"taxable"`
    Amount      float64   `gorm:"type:decimal(12,2)" json:"amount"` // Auto-calculated
    CreatedAt   time.Time `json:"created_at"`
}



