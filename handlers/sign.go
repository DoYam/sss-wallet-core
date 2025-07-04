package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"sss-wallet-core/core"
)

type SignRequest struct {
	Message string `json:"message"`
}

func SignHandler(w http.ResponseWriter, r *http.Request) {
	var req SignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid", http.StatusBadRequest)
		return
	}

	// 랜덤하게 3개 share 선택
	var selected [][]byte
	perm := rand.Perm(len(core.Shares))
	for i := 0; i < 3; i++ {
		selected = append(selected, core.Shares[perm[i]])
	}

	priv, err := core.RebuildPrivateKey(selected)
	if err != nil {
		http.Error(w, "combine fail", http.StatusInternalServerError)
		return
	}

	sig, err := core.SignWithKey(priv, req.Message)
	if err != nil {
		http.Error(w, "sign fail", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sig)
}
