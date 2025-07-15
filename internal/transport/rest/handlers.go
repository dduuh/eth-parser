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

	ctx := r.Context()

	ethAddress, privateKey := generateETHAddress()

	addr := domain.Addresses{
		Address:    ethAddress,
		PrivateKey: privateKey,
	}

	createdAddr, err := s.services.AddAddress(ctx, addr)
	if err != nil {
		response(w, http.StatusInternalServerError, err.Error())

		return
	}

	response(w, http.StatusCreated, Dict{
		"address": createdAddr,
	})
}
