package models

import "time"

type PurchaseOrder struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PONumber string `gorm:"size:100;uniqueIndex:idx_po_number;not null" json:"po_number"`

	RFQID       uint `gorm:"not null;index" json:"rfq_id"`
	QuotationID uint `gorm:"uniqueIndex:idx_po_quotation;not null" json:"quotation_id"`
	BuyerID     uint `gorm:"not null;index" json:"buyer_id"`
	SupplierID  uint `gorm:"not null;index" json:"supplier_id"`

	PODate   time.Time `gorm:"type:date;not null;default:CURRENT_DATE" json:"po_date"`
	VendorID string    `gorm:"size:100" json:"vendor_id"`

	BuyerCompanyName string `gorm:"size:255;not null" json:"buyer_company_name"`
	BuyerAddress     string `gorm:"type:text;not null" json:"buyer_address"`
	BuyerContactInfo string `gorm:"size:255" json:"buyer_contact_info"`

	SupplierCompanyName   string `gorm:"size:255;not null" json:"supplier_company_name"`
	SupplierAddress       string `gorm:"type:text;not null" json:"supplier_address"`
	SupplierContactInfo   string `gorm:"size:255" json:"supplier_contact_info"`
	SupplierContactPerson string `gorm:"size:255" json:"supplier_contact_person"`

	ShippingCompanyName   string `gorm:"size:255" json:"shipping_company_name"`
	ShippingAddress       string `gorm:"type:text" json:"shipping_address"`
	ShippingContactInfo   string `gorm:"size:255" json:"shipping_contact_info"`
	ShippingContactPerson string `gorm:"size:255" json:"shipping_contact_person"`

	DeliveryDate   *time.Time `gorm:"type:date" json:"delivery_date"`
	ShippingMethod string     `gorm:"size:100" json:"shipping_method"`
	ShippingTerms  string     `gorm:"type:text" json:"shipping_terms"`

	PaymentDueDate *time.Time `gorm:"type:date" json:"payment_due_date"`
	PaymentTerms   string     `gorm:"type:text" json:"payment_terms"`

	Subtotal    float64 `gorm:"type:decimal(12,2);not null" json:"subtotal"`
	TaxRate     float64 `gorm:"type:decimal(5,2);default:0" json:"tax_rate"`
	TaxAmount   float64 `gorm:"type:decimal(12,2);default:0" json:"tax_amount"`
	TotalAmount float64 `gorm:"type:decimal(12,2);not null" json:"total_amount"`

	Status string `gorm:"size:50;default:'draft'" json:"status"`
	Notes  string `gorm:"type:text" json:"notes"`

	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	DeliveredAt *time.Time `json:"delivered_at"`

	Items []PurchaseOrderItem `gorm:"foreignKey:PurchaseOrderID" json:"items,omitempty"`
}

type PurchaseOrderItem struct {
	ID              uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	PurchaseOrderID uint   `gorm:"not null;index" json:"purchase_order_id"`
	RFQItemID       uint   `gorm:"not null;index" json:"rfq_item_id"`
	LineNumber      int    `gorm:"not null" json:"line_number"`

	ItemCode    string  `gorm:"size:100" json:"item_code"`
	Description string  `gorm:"type:text;not null" json:"description"`
	Quantity    int     `gorm:"not null;default:1" json:"quantity"`
	UnitPrice   float64 `gorm:"type:decimal(12,2);not null" json:"unit_price"`
	TotalPrice  float64 `gorm:"type:decimal(12,2)" json:"total_price"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}