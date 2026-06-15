package models

import "time"

type RFQ struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	RFQNumber string `gorm:"size:50;uniqueIndex;not null" json:"rfq_number"`
	BuyerID   uint   `gorm:"not null;index" json:"buyer_id"`

	BuyerCompanyName string `gorm:"size:255;not null" json:"buyer_company_name"`
	BuyerAddress     string `gorm:"type:text;not null" json:"buyer_address"`
	BuyerContactInfo string `gorm:"size:255" json:"buyer_contact_info"`

	SupplierCompanyName string `gorm:"size:255" json:"supplier_company_name"`
	SupplierAddress     string `gorm:"type:text" json:"supplier_address"`
	SupplierContactInfo string `gorm:"size:255" json:"supplier_contact_info"`

	PONumber        string     `gorm:"size:100" json:"po_number"`
	ShipDate        *time.Time `gorm:"type:date" json:"ship_date"`
	ShipVia         string     `gorm:"size:100" json:"ship_via"`
	FOBPoint        string     `gorm:"size:100" json:"fob_point"`
	PaymentTerms    string     `gorm:"size:100" json:"payment_terms"`
	QuoteValidUntil *time.Time `gorm:"type:date" json:"quote_valid_until"`

	Status              string `gorm:"size:50;default:'draft'" json:"status"`
	Comments            string `gorm:"type:text" json:"comments"`
	SpecialInstructions string `gorm:"type:text" json:"special_instructions"`

	Items      []RFQItem   `gorm:"foreignKey:RFQID" json:"items,omitempty"`
	Quotations []Quotation `gorm:"foreignKey:RFQID" json:"quotations,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy uint      `gorm:"index" json:"created_by"`
}

type RFQItem struct {
	ID          uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	RFQID       uint     `gorm:"not null;index" json:"rfq_id"`
	LineNumber  int      `gorm:"not null" json:"line_number"`
	Quantity    float64  `gorm:"type:decimal(12,2);not null;default:1" json:"quantity"`
	Description string   `gorm:"type:text;not null" json:"description"`
	UnitPrice   *float64 `gorm:"type:decimal(12,2)" json:"unit_price"`
	Taxable     bool     `gorm:"default:false" json:"taxable"`
	Amount      float64  `gorm:"type:decimal(12,2)" json:"amount"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}