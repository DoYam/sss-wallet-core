package handlers

import (
	"encoding/json"
	"math/big"
	"net/http"

	"sss-wallet-core/core"
)

type VerifyRequest struct {
	Message string `json:"message"`
	R       string `json:"r"`
	S       string `json:"s"`
}

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
		return
	}

	rInt := new(big.Int)
	sInt := new(big.Int)
	rInt.SetString(req.R, 16)
	sInt.SetString(req.S, 16)

	ok := core.Verify(req.Message, rInt, sInt, core.PubKey)

	json.NewEncoder(w).Encode(map[string]bool{"valid": ok})
}
