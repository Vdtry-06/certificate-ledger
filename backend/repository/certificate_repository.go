package repository

import (
	"fmt"
	"sync"

	"certificate-ledger/domain"
)

type CertificateRepository struct {
	certificates map[string]*domain.Certificate
	mu           sync.RWMutex
}

func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		certificates: make(map[string]*domain.Certificate),
	}
}

func (r *CertificateRepository) Save(cert *domain.Certificate) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	r.certificates[cert.ID] = cert
	return nil
}

func (r *CertificateRepository) FindByID(id string) (*domain.Certificate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	cert, ok := r.certificates[id]
	if !ok {
		return nil, fmt.Errorf("certificate with ID %s not found", id)
	}
	
	return cert, nil
}

func (r *CertificateRepository) FindByHash(hash string) (*domain.Certificate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	for _, cert := range r.certificates {
		if cert.Hash == hash {
			return cert, nil
		}
	}
	
	return nil, fmt.Errorf("certificate with hash %s not found", hash)
}

func (r *CertificateRepository) FindAll() ([]*domain.Certificate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	certs := make([]*domain.Certificate, 0, len(r.certificates))
	for _, cert := range r.certificates {
		certs = append(certs, cert)
	}
	
	return certs, nil
}

func (r *CertificateRepository) FindByIssuerID(userID string) ([]*domain.Certificate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	var certs []*domain.Certificate
	for _, cert := range r.certificates {
		if cert.IssuerID == userID {
			certs = append(certs, cert)
		}
	}
	
	return certs, nil
}

func (r *CertificateRepository) FindByRecipientEmail(email string) ([]*domain.Certificate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	var certs []*domain.Certificate
	for _, cert := range r.certificates {
		if cert.RecipientEmail == email {
			certs = append(certs, cert)
		}
	}
	
	return certs, nil
}
