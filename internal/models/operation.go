package models

import (
	"time"

	"github.com/google/uuid"
)

// ProductionStage represents the predefined stages in manufacturing
type ProductionStage struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"size:100;uniqueIndex;not null" json:"name"` // e.g., "Material Sourcing", "Cutting", "Sewing", "Finishing", "QC Inspection", "Packaging", "Ready for Shipment"
	Order       int       `gorm:"not null" json:"order"` // Sequence order (1, 2, 3...)
	Description string    `gorm:"type:text" json:"description"`
}

// ProductionOrder tracks the manufacturing progress of a purchase order
type ProductionOrder struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	PurchaseOrderID   uuid.UUID `gorm:"type:uuid;uniqueIndex;not null" json:"purchase_order_id"`
	SupplierID        uuid.UUID `gorm:"type:uuid;not null;index" json:"supplier_id"`
	
	// Production Timeline
	StartDate         *time.Time `json:"start_date"`
	EstimatedEndDate  *time.Time `json:"estimated_end_date"`
	ActualEndDate     *time.Time `json:"actual_end_date"`
	
	// Overall Progress (0-100)
	ProgressPercent   int       `gorm:"default:0;check:progress_percent >= 0 AND progress_percent <= 100" json:"progress_percent"`
	
	// Status: pending → in_progress → qc_check → ready_to_ship → shipped → completed
	Status            string    `gorm:"size:50;default:'pending';check:status IN ('pending','in_progress','qc_check','ready_to_ship','shipped','completed','on_hold','cancelled')" json:"status"`
	
	// Notes & Issues
	ProductionNotes   string    `gorm:"type:text" json:"production_notes"`
	Issues            string    `gorm:"type:text" json:"issues"`
	
	// Timestamps
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	
	// Relations
	PurchaseOrder     *PurchaseOrder      `gorm:"foreignKey:PurchaseOrderID" json:"purchase_order,omitempty"`
	Supplier          *SupplierProfile    `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	Milestones        []ProductionMilestone `gorm:"foreignKey:ProductionOrderID" json:"milestones,omitempty"`
	Documents         []ProductionDocument  `gorm:"foreignKey:ProductionOrderID" json:"documents,omitempty"`
}

// ProductionMilestone tracks individual stage completion
type ProductionMilestone struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ProductionOrderID uuid.UUID `gorm:"type:uuid;not null;index" json:"production_order_id"`
	StageID           uuid.UUID `gorm:"type:uuid;not null" json:"stage_id"`
	
	// Status: pending → in_progress → completed → skipped
	Status            string    `gorm:"size:50;default:'pending';check:status IN ('pending','in_progress','completed','skipped')" json:"status"`
	
	// Timeline
	StartedAt         *time.Time `json:"started_at"`
	CompletedAt       *time.Time `json:"completed_at"`
	
	// Progress within this milestone (0-100)
	ProgressPercent   int       `gorm:"default:0;check:progress_percent >= 0 AND progress_percent <= 100" json:"progress_percent"`
	
	// Details
	Notes             string    `gorm:"type:text" json:"notes"`
	WorkerCount       int       `json:"worker_count"`
	MachineCount      int       `json:"machine_count"`
	
	// Quality metrics
	DefectRate        float64   `gorm:"type:decimal(5,2)" json:"defect_rate"` // Percentage
	QualityScore      float64   `gorm:"type:decimal(3,2)" json:"quality_score"` // 0-100
	
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	
	// Relations
	Stage             *ProductionStage `gorm:"foreignKey:StageID" json:"stage,omitempty"`
}

// ProductionDocument stores all files uploaded during production
type ProductionDocument struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ProductionOrderID uuid.UUID `gorm:"type:uuid;not null;index" json:"production_order_id"`
	MilestoneID       *uuid.UUID `gorm:"type:uuid;index" json:"milestone_id,omitempty"` // Optional: link to specific milestone
	
	// Document Categories
	Category          string    `gorm:"size:100;not null;check:category IN ('raw_material','production','qc','packaging','export','other')" json:"category"`
	DocumentType      string    `gorm:"size:100;not null" json:"document_type"` // e.g., "Material Invoice", "Cutting Report", "QC Photo", "Packing List", "Commercial Invoice", "Bill of Lading"
	
	// File Info
	FileName          string    `gorm:"size:255;not null" json:"file_name"`
	FileURL           string    `gorm:"size:500;not null" json:"file_url"`
	FileSize          int64     `json:"file_size"` // in bytes
	MimeType          string    `gorm:"size:100" json:"mime_type"`
	
	// Metadata
	Description       string    `gorm:"type:text" json:"description"`
	UploadedBy        uuid.UUID `gorm:"type:uuid;not null" json:"uploaded_by"`
	
	// Verification
	IsVerified        bool      `gorm:"default:false" json:"is_verified"`
	VerifiedBy        *uuid.UUID `gorm:"type:uuid" json:"verified_by,omitempty"`
	VerifiedAt        *time.Time `json:"verified_at"`
	
	// Versioning (for documents that get updated)
	Version           int       `gorm:"default:1" json:"version"`
	ReplacesDocID     *uuid.UUID `gorm:"type:uuid" json:"replaces_doc_id,omitempty"`
	
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	
	// Relations
	ProductionOrder   *ProductionOrder `gorm:"foreignKey:ProductionOrderID" json:"-"`
	Milestone         *ProductionMilestone `gorm:"foreignKey:MilestoneID" json:"milestone,omitempty"`
	UploadedByUser    *User            `gorm:"foreignKey:UploadedBy" json:"uploaded_by_user,omitempty"`
}

// ExportDocument specifically for shipping/export paperwork
type ExportDocument struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ProductionOrderID uuid.UUID `gorm:"type:uuid;uniqueIndex:idx_export_doc_po;not null" json:"production_order_id"`
	DocumentType      string    `gorm:"size:100;not null;check:document_type IN ('commercial_invoice','packing_list','bill_of_lading','certificate_of_origin','inspection_certificate','insurance_certificate','export_license','customs_declaration')" json:"document_type"`
	
	// Document Details
	DocumentNumber    string    `gorm:"size:100;uniqueIndex" json:"document_number"`
	IssueDate         time.Time `json:"issue_date"`
	ExpiryDate        *time.Time `json:"expiry_date"`
	
	// File
	FileURL           string    `gorm:"size:500;not null" json:"file_url"`
	FileName          string    `gorm:"size:255;not null" json:"file_name"`
	
	// Status: draft → submitted → approved → rejected
	Status            string    `gorm:"size:50;default:'draft';check:status IN ('draft','submitted','approved','rejected')" json:"status"`
	
	// Approval workflow
	SubmittedAt       *time.Time `json:"submitted_at"`
	ApprovedAt        *time.Time `json:"approved_at"`
	ApprovedBy        *uuid.UUID `gorm:"type:uuid" json:"approved_by,omitempty"`
	RejectionReason   string    `gorm:"type:text" json:"rejection_reason"`
	
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	
	// Relations
	ProductionOrder   *ProductionOrder `gorm:"foreignKey:ProductionOrderID" json:"-"`
}

// ProductionUpdate represents a status update/log from supplier
type ProductionUpdate struct {
	ID                uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	ProductionOrderID uuid.UUID `gorm:"type:uuid;not null;index" json:"production_order_id"`
	UpdatedBy         uuid.UUID `gorm:"type:uuid;not null" json:"updated_by"`
	
	UpdateType        string    `gorm:"size:50;check:update_type IN ('milestone','document','issue','general')" json:"update_type"`
	Message           string    `gorm:"type:text;not null" json:"message"`
	
	// Attachments (photos, files)
	AttachmentURLs    []string  `gorm:"type:text[]" json:"attachment_urls"`
	
	// Visibility
	IsVisibleToBuyer  bool      `gorm:"default:true" json:"is_visible_to_buyer"`
	
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at"`
	
	// Relations
	ProductionOrder   *ProductionOrder `gorm:"foreignKey:ProductionOrderID" json:"-"`
	UpdatedByUser     *User            `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}