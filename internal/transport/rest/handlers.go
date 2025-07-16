package rest

import (
	"eth-parser/internal/domain"
	"net/http"
)

const incorrectMethod = "incorrect method"

func (s *Server) addAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response(w, http.StatusMethodNotAllowed, incorrectMethod)

		return
	}

	ethAddress, privateKey := generateAddresses()

	addr := domain.Addresses{
		Address:    ethAddress,
		PrivateKey: privateKey,
	}

	createdAddr, err := s.services.AddAddress(r.Context(), addr)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())

		return
	}

	response(w, http.StatusCreated, Dict{
		"address": createdAddr,
	})
}
