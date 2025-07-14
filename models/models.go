package models

import "time"

type User struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"` // admin, researcher, guest, lab_technician
	CreatedAt    time.Time `json:"created_at"`
}

type Genome struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Species          string    `json:"species"`
	ReferenceVersion string    `json:"reference_version"`
	CreatedBy        int       `json:"created_by"`
	CreatedAt        time.Time `json:"created_at"`
}

type Sample struct {
	ID             int       `json:"id"`
	GenomeID       int       `json:"genome_id"`
	DonorID        string    `json:"donor_id"`
	CollectionDate string    `json:"collection_date"`
	SampleType     string    `json:"sample_type"`
	Metadata       string    `json:"metadata"` // JSON as raw string or use map[string]interface{}
	CollectedBy    int       `json:"collected_by"`
	CreatedAt      time.Time `json:"created_at"`
}

type SequenceFile struct {
	ID         int       `json:"id"`
	SampleID   int       `json:"sample_id"`
	FilePath   string    `json:"file_path"`
	FileType   string    `json:"file_type"`
	Checksum   string    `json:"checksum"`
	UploadedBy int       `json:"uploaded_by"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type VariantFile struct {
	ID         int       `json:"id"`
	SampleID   int       `json:"sample_id"`
	GenomeID   int       `json:"genome_id"`
	FilePath   string    `json:"file_path"`
	FileType   string    `json:"file_type"`
	Checksum   string    `json:"checksum"`
	UploadedBy int       `json:"uploaded_by"`
	UploadedAt time.Time `json:"uploaded_at"`
}

type AuditLog struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	Action       string    `json:"action"`
	ResourceType string    `json:"resource_type"`
	ResourceID   int       `json:"resource_id"`
	Timestamp    time.Time `json:"timestamp"`
	Details      string    `json:"details"`
}
