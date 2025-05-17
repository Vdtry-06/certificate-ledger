package handler

import (
	"encoding/json"
	"net/http"

	"certificate-ledger/domain"
	"certificate-ledger/service"

	"github.com/gorilla/mux"
)

type CertificateHandler struct {
	service *service.CertificateService
}

func NewCertificateHandler(service *service.CertificateService) *CertificateHandler {
	return &CertificateHandler{
		service: service,
	}
}

func (h *CertificateHandler) CreateCertificate(w http.ResponseWriter, r *http.Request) {
    var req domain.CertificateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    user, ok := r.Context().Value("user").(*domain.User)
    if !ok || user == nil {
        http.Error(w, "Unauthorized: User not found in context", http.StatusUnauthorized)
        return
    }

    cert, err := h.service.CreateCertificate(req, user.ID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(cert)
}

func (h *CertificateHandler) GetCertificate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	cert, err := h.service.GetCertificate(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cert)
}

func (h *CertificateHandler) VerifyCertificate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	isValid, err := h.service.VerifyCertificate(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{"valid": isValid})
}

func (h *CertificateHandler) GetAllCertificates(w http.ResponseWriter, r *http.Request) {
	certs, err := h.service.GetAllCertificates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certs)
}
