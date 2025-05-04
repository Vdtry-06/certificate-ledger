package domain

import (
	"time"
)

type Certificate struct {
	ID              string    `json:"id"`
	Hash            string    `json:"hash"`
	RecipientName   string    `json:"recipientName"`
	RecipientEmail  string    `json:"recipientEmail"`
	CertificateTitle string   `json:"certificateTitle"`
	IssueDate       time.Time `json:"issueDate"`
	IssuerID        string    `json:"issuerId"`
	IssuerName      string    `json:"issuerName"`
	Description     string    `json:"description"`
	BlockNumber     int       `json:"blockNumber"`
	Timestamp       time.Time `json:"timestamp"`
}

type CertificateRequest struct {
	RecipientName   string `json:"recipientName"`
	RecipientEmail  string `json:"recipientEmail"`
	CertificateTitle string `json:"certificateTitle"`
	IssueDate       string `json:"issueDate"`
	IssuerName      string `json:"issuerName"`
	Description     string `json:"description"`
}
