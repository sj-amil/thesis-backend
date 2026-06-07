package models

import (
	"time"

	"github.com/google/uuid"
)

// PurchaseOrder represents the main order document after quotation acceptance
type PurchaseOrder struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"` // Generated via Go hook below
	PONumber        string    `gorm:"size:100;uniqueIndex:idx_po_number;not null" json:"po_number"`
	
	// Crucial Workflow Links
	RFQID           uuid.UUID `gorm:"type:uuid;not null;index" json:"rfq_id"` // 👈 Added: Traces directly back to the original RFQ
	QuotationID     uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_po_quotation;not null" json:"quotation_id"`
	BuyerID         uuid.UUID `gorm:"type:uuid;not null;index" json:"buyer_id"`
	SupplierID      uuid.UUID `gorm:"type:uuid;not null;index" json:"supplier_id"`
	
	// PO Metadata
	PODate          time.Time `gorm:"type:date;not null;default:CURRENT_DATE" json:"po_date"`
	VendorID        string    `gorm:"size:100" json:"vendor_id"`
	
	// Buyer Company Info (Snapshot at time of order)
	BuyerCompanyName string     `gorm:"size:255;not null" json:"buyer_company_name"`
	BuyerAddress     string     `gorm:"type:text;not null" json:"buyer_address"`
	BuyerContactInfo string     `gorm:"size:255" json:"buyer_contact_info"`
	
	// Supplier Info (Purchased From)
	SupplierCompanyName string  `gorm:"size:255;not null" json:"supplier_company_name"`
	SupplierAddress     string  `gorm:"type:text;not null" json:"supplier_address"`
	SupplierContactInfo string  `gorm:"size:255" json:"supplier_contact_info"`
	SupplierContactPerson string `gorm:"size:255" json:"supplier_contact_person"`
	
	// Shipping Address (Shipped To)
	ShippingCompanyName string `gorm:"size:255" json:"shipping_company_name"`
	ShippingAddress     string `gorm:"type:text" json:"shipping_address"`
	ShippingContactInfo string `gorm:"size:255" json:"shipping_contact_info"`
	ShippingContactPerson string `gorm:"size:255" json:"shipping_contact_person"`
	
	// Delivery & Shipping
	DeliveryDate    *time.Time `gorm:"type:date" json:"delivery_date"`
	ShippingMethod  string     `gorm:"size:100" json:"shipping_method"`
	ShippingTerms   string     `gorm:"type:text" json:"shipping_terms"`
	
	// Payment Terms
	PaymentDueDate *time.Time `gorm:"type:date" json:"payment_due_date"`
	PaymentTerms   string     `gorm:"type:text" json:"payment_terms"`
	
	// Pricing
	Subtotal    float64 `gorm:"type:decimal(12,2);not null" json:"subtotal"`
	TaxRate     float64 `gorm:"type:decimal(5,2);default:0" json:"tax_rate"`
	TaxAmount   float64 `gorm:"type:decimal(12,2);default:0" json:"tax_amount"`
	TotalAmount float64 `gorm:"type:decimal(12,2);not null" json:"total_amount"`
	
	// Status Workflow
	Status        string     `gorm:"size:50;default:'draft';check:status IN ('draft','sent','confirmed','in_production','shipped','delivered','completed','cancelled')" json:"status"`
	Notes         string     `gorm:"type:text" json:"notes"`
	
	// Timestamps
	CreatedAt     time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	ConfirmedAt   *time.Time `json:"confirmed_at"`
	ShippedAt     *time.Time `json:"shipped_at"`
	DeliveredAt   *time.Time `json:"delivered_at"`
	
	// Relations
	Items         []PurchaseOrderItem `gorm:"foreignKey:PurchaseOrderID" json:"items,omitempty"`
}

// PurchaseOrderItem represents line items in a purchase order originating from an RFQ item
type PurchaseOrderItem struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"` // Generated via Go hook below
	PurchaseOrderID uuid.UUID `gorm:"type:uuid;not null;index" json:"purchase_order_id"`
	RFQItemID       uuid.UUID `gorm:"type:uuid;not null;index" json:"rfq_item_id"` // 👈 Added: Links directly to the original RFQ input field/item row
	LineNumber      int       `gorm:"not null" json:"line_number"`
	
	ItemCode        string    `gorm:"size:100" json:"item_code"`
	Description     string    `gorm:"type:text;not null" json:"description"`
	Quantity        int       `gorm:"not null;default:1" json:"quantity"`
	UnitPrice       float64   `gorm:"type:decimal(12,2);not null" json:"unit_price"`
	TotalPrice      float64   `gorm:"type:decimal(12,2)" json:"total_price"` // Managed via Go hooks
	
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
