package repository

import (
	"database/sql"
	"fmt"

	"certificate-ledger/domain"
)

type CertificateRepository struct {
	db *sql.DB
}

func NewCertificateRepository(db *sql.DB) *CertificateRepository {
	return &CertificateRepository{db: db}
}

func (r *CertificateRepository) Save(cert *domain.Certificate) error {
	query := `
		INSERT INTO certificates (id, hash, recipient_name, recipient_email, certificate_title, issue_date, issuer_id, issuer_name, description, block_number, timestamp)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query,
		cert.ID,
		cert.Hash,
		cert.RecipientName,
		cert.RecipientEmail,
		cert.CertificateTitle,
		cert.IssueDate,
		cert.IssuerID,
		cert.IssuerName,
		cert.Description,
		cert.BlockNumber,
		cert.Timestamp,
	)
	if err != nil {
		return fmt.Errorf("failed to save certificate: %v", err)
	}
	return nil
}

func (r *CertificateRepository) FindByID(id string) (*domain.Certificate, error) {
	query := `SELECT id, hash, recipient_name, recipient_email, certificate_title, issue_date, issuer_id, issuer_name, description, block_number, timestamp
	          FROM certificates WHERE id = ?`
	row := r.db.QueryRow(query, id)

	var cert domain.Certificate
	err := row.Scan(
		&cert.ID,
		&cert.Hash,
		&cert.RecipientName,
		&cert.RecipientEmail,
		&cert.CertificateTitle,
		&cert.IssueDate,
		&cert.IssuerID,
		&cert.IssuerName,
		&cert.Description,
		&cert.BlockNumber,
		&cert.Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("certificate with ID %s not found", id)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find certificate: %v", err)
	}
	return &cert, nil
}

func (r *CertificateRepository) FindByHash(hash string) (*domain.Certificate, error) {
	query := `SELECT id, hash, recipient_name, recipient_email, certificate_title, issue_date, issuer_id, issuer_name, description, block_number, timestamp
	          FROM certificates WHERE hash = ?`
	row := r.db.QueryRow(query, hash)

	var cert domain.Certificate
	err := row.Scan(
		&cert.ID,
		&cert.Hash,
		&cert.RecipientName,
		&cert.RecipientEmail,
		&cert.CertificateTitle,
		&cert.IssueDate,
		&cert.IssuerID,
		&cert.IssuerName,
		&cert.Description,
		&cert.BlockNumber,
		&cert.Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("certificate with hash %s not found", hash)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find certificate: %v", err)
	}
	return &cert, nil
}

func (r *CertificateRepository) FindAll() ([]*domain.Certificate, error) {
	query := `SELECT id, hash, recipient_name, recipient_email, certificate_title, issue_date, issuer_id, issuer_name, description, block_number, timestamp
	          FROM certificates`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query certificates: %v", err)
	}
	defer rows.Close()

	var certs []*domain.Certificate
	for rows.Next() {
		var cert domain.Certificate
		if err := rows.Scan(
			&cert.ID,
			&cert.Hash,
			&cert.RecipientName,
			&cert.RecipientEmail,
			&cert.CertificateTitle,
			&cert.IssueDate,
			&cert.IssuerID,
			&cert.IssuerName,
			&cert.Description,
			&cert.BlockNumber,
			&cert.Timestamp,
		); err != nil {
			return nil, fmt.Errorf("failed to scan certificate: %v", err)
		}
		certs = append(certs, &cert)
	}
	return certs, nil
}

func (r *CertificateRepository) FindByIssuerID(userID string) ([]*domain.Certificate, error) {
	query := `SELECT id, hash, recipient_name, recipient_email, certificate_title, issue_date, issuer_id, issuer_name, description, block_number, timestamp
	          FROM certificates WHERE issuer_id = ?`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query certificates: %v", err)
	}
	defer rows.Close()

	var certs []*domain.Certificate
	for rows.Next() {
		var cert domain.Certificate
		if err := rows.Scan(
			&cert.ID,
			&cert.Hash,
			&cert.RecipientName,
			&cert.RecipientEmail,
			&cert.CertificateTitle,
			&cert.IssueDate,
			&cert.IssuerID,
			&cert.IssuerName,
			&cert.Description,
			&cert.BlockNumber,
			&cert.Timestamp,
		); err != nil {
			return nil, fmt.Errorf("failed to scan certificate: %v", err)
		}
		certs = append(certs, &cert)
	}
	return certs, nil
}

func (r *CertificateRepository) FindByRecipientEmail(email string) ([]*domain.Certificate, error) {
	query := `SELECT id, hash, recipient_name, recipient_email, certificate_title, issue_date, issuer_id, issuer_name, description, block_number, timestamp
	          FROM certificates WHERE recipient_email = ?`
	rows, err := r.db.Query(query, email)
	if err != nil {
		return nil, fmt.Errorf("failed to query certificates: %v", err)
	}
	defer rows.Close()

	var certs []*domain.Certificate
	for rows.Next() {
		var cert domain.Certificate
		if err := rows.Scan(
			&cert.ID,
			&cert.Hash,
			&cert.RecipientName,
			&cert.RecipientEmail,
			&cert.CertificateTitle,
			&cert.IssueDate,
			&cert.IssuerID,
			&cert.IssuerName,
			&cert.Description,
			&cert.BlockNumber,
			&cert.Timestamp,
		); err != nil {
			return nil, fmt.Errorf("failed to scan certificate: %v", err)
		}
		certs = append(certs, &cert)
	}
	return certs, nil
}