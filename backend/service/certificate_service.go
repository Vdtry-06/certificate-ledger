package service

import (
	"encoding/json"
	"fmt"
	"time"

	"certificate-ledger/blockchain"
	"certificate-ledger/domain"
	"certificate-ledger/repository"
	"github.com/google/uuid"
)

type CertificateService struct {
	repo       *repository.CertificateRepository
	blockchain *blockchain.Blockchain
}

func NewCertificateService(repo *repository.CertificateRepository, bc *blockchain.Blockchain) *CertificateService {
	return &CertificateService{
		repo:       repo,
		blockchain: bc,
	}
}

func (s *CertificateService) CreateCertificate(req domain.CertificateRequest, userID string) (*domain.Certificate, error) {
	certID := fmt.Sprintf("CERT-%s", uuid.New().String()[:6])

	issueDate, err := time.Parse("2006-01-02", req.IssueDate)
	if err != nil {
		return nil, fmt.Errorf("invalid issue date format: %v", err)
	}

	cert := &domain.Certificate{
		ID:               certID,
		RecipientName:    req.RecipientName,
		RecipientEmail:   req.RecipientEmail,
		CertificateTitle: req.CertificateTitle,
		IssueDate:        issueDate,
		IssuerID:         userID,
		IssuerName:       req.IssuerName,
		Description:      req.Description,
		Timestamp:        time.Now(),
	}

	certData, err := json.Marshal(cert)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal certificate: %v", err)
	}

	block := s.blockchain.AddBlock(certData)

	cert.Hash = block.Hash
	cert.BlockNumber = block.Index

	if err := s.repo.Save(cert); err != nil {
		return nil, fmt.Errorf("failed to save certificate: %v", err)
	}

	return cert, nil
}

func (s *CertificateService) GetCertificate(id string) (*domain.Certificate, error) {
	return s.repo.FindByID(id)
}

func (s *CertificateService) GetCertificateByHash(hash string) (*domain.Certificate, error) {
	return s.repo.FindByHash(hash)
}

func (s *CertificateService) VerifyCertificate(hash string) (bool, error) {
	cert, err := s.repo.FindByHash(hash)
	if err != nil {
		return false, err
	}

	block, err := s.blockchain.GetBlock(hash)
	if err != nil {
		return false, err
	}

	var blockCert domain.Certificate
	if err := json.Unmarshal(block.Data, &blockCert); err != nil {
		return false, err
	}

	if cert.ID != blockCert.ID || 
	   cert.RecipientName != blockCert.RecipientName || 
	   cert.CertificateTitle != blockCert.CertificateTitle {
		return false, nil
	}

	return true, nil
}

func (s *CertificateService) GetAllCertificates() ([]*domain.Certificate, error) {
	return s.repo.FindAll()
}

func (s *CertificateService) GetCertificatesByIssuer(userID string) ([]*domain.Certificate, error) {
	return s.repo.FindByIssuerID(userID)
}

func (s *CertificateService) GetCertificatesByRecipient(email string) ([]*domain.Certificate, error) {
	return s.repo.FindByRecipientEmail(email)
}
