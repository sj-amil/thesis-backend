package models

import "time"

type ProductionStage struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Order       int    `gorm:"not null" json:"order"`
	Description string `gorm:"type:text" json:"description"`
}

type ProductionOrder struct {
	ID              uint `gorm:"primaryKey;autoIncrement" json:"id"`
	PurchaseOrderID uint `gorm:"uniqueIndex;not null" json:"purchase_order_id"`
	SupplierID      uint `gorm:"not null;index" json:"supplier_id"`

	StartDate        *time.Time `json:"start_date"`
	EstimatedEndDate *time.Time `json:"estimated_end_date"`
	ActualEndDate    *time.Time `json:"actual_end_date"`

	ProgressPercent int `gorm:"default:0" json:"progress_percent"`
	Status          string `gorm:"size:50;default:'pending'" json:"status"`

	ProductionNotes string `gorm:"type:text" json:"production_notes"`
	Issues          string `gorm:"type:text" json:"issues"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	PurchaseOrder *PurchaseOrder        `gorm:"foreignKey:PurchaseOrderID" json:"purchase_order,omitempty"`
	Supplier      *SupplierProfile      `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	Milestones    []ProductionMilestone `gorm:"foreignKey:ProductionOrderID" json:"milestones,omitempty"`
	Documents     []ProductionDocument  `gorm:"foreignKey:ProductionOrderID" json:"documents,omitempty"`
}

type ProductionMilestone struct {
	ID                uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductionOrderID uint `gorm:"not null;index" json:"production_order_id"`
	StageID           uint `gorm:"not null;index" json:"stage_id"`

	Status string `gorm:"size:50;default:'pending'" json:"status"`

	StartedAt   *time.Time `json:"started_at"`
	CompletedAt *time.Time `json:"completed_at"`

	ProgressPercent int `gorm:"default:0" json:"progress_percent"`

	Notes        string `gorm:"type:text" json:"notes"`
	WorkerCount  int    `json:"worker_count"`
	MachineCount int    `json:"machine_count"`

	DefectRate   float64 `gorm:"type:decimal(5,2)" json:"defect_rate"`
	QualityScore float64 `gorm:"type:decimal(3,2)" json:"quality_score"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Stage *ProductionStage `gorm:"foreignKey:StageID" json:"stage,omitempty"`
}

type ProductionDocument struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductionOrderID uint   `gorm:"not null;index" json:"production_order_id"`
	MilestoneID       *uint  `gorm:"index" json:"milestone_id,omitempty"`

	Category     string `gorm:"size:100;not null" json:"category"`
	DocumentType string `gorm:"size:100;not null" json:"document_type"`

	FileName string `gorm:"size:255;not null" json:"file_name"`
	FileURL  string `gorm:"size:500;not null" json:"file_url"`
	FileSize int64  `json:"file_size"`
	MimeType string `gorm:"size:100" json:"mime_type"`

	Description string `gorm:"type:text" json:"description"`
	UploadedBy  uint   `gorm:"not null;index" json:"uploaded_by"`

	IsVerified bool       `gorm:"default:false" json:"is_verified"`
	VerifiedBy *uint      `json:"verified_by,omitempty"`
	VerifiedAt *time.Time `json:"verified_at"`

	Version       int   `gorm:"default:1" json:"version"`
	ReplacesDocID *uint `json:"replaces_doc_id,omitempty"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	ProductionOrder *ProductionOrder     `gorm:"foreignKey:ProductionOrderID" json:"-"`
	Milestone       *ProductionMilestone `gorm:"foreignKey:MilestoneID" json:"milestone,omitempty"`
	UploadedByUser  *User                `gorm:"foreignKey:UploadedBy" json:"uploaded_by_user,omitempty"`
}

type ExportDocument struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductionOrderID uint   `gorm:"uniqueIndex:idx_export_doc_po;not null" json:"production_order_id"`
	DocumentType      string `gorm:"size:100;not null" json:"document_type"`

	DocumentNumber string     `gorm:"size:100;uniqueIndex" json:"document_number"`
	IssueDate      time.Time  `json:"issue_date"`
	ExpiryDate     *time.Time `json:"expiry_date"`

	FileURL  string `gorm:"size:500;not null" json:"file_url"`
	FileName string `gorm:"size:255;not null" json:"file_name"`

	Status string `gorm:"size:50;default:'draft'" json:"status"`

	SubmittedAt     *time.Time `json:"submitted_at"`
	ApprovedAt      *time.Time `json:"approved_at"`
	ApprovedBy      *uint      `json:"approved_by,omitempty"`
	RejectionReason string     `gorm:"type:text" json:"rejection_reason"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	ProductionOrder *ProductionOrder `gorm:"foreignKey:ProductionOrderID" json:"-"`
}

type ProductionUpdate struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductionOrderID uint   `gorm:"not null;index" json:"production_order_id"`
	UpdatedBy         uint   `gorm:"not null;index" json:"updated_by"`

	UpdateType string `gorm:"size:50" json:"update_type"`
	Message    string `gorm:"type:text;not null" json:"message"`

	AttachmentURLs []string `gorm:"serializer:json" json:"attachment_urls"`

	IsVisibleToBuyer bool `gorm:"default:true" json:"is_visible_to_buyer"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	ProductionOrder *ProductionOrder `gorm:"foreignKey:ProductionOrderID" json:"-"`
	UpdatedByUser   *User            `gorm:"foreignKey:UpdatedBy" json:"updated_by_user,omitempty"`
}