package handlers

import (
	"encoding/json"
	"net/http"

	"sss-wallet-core/core"
)

func WalletHandler(w http.ResponseWriter, r *http.Request) {
	pub := core.PubKey
	res := map[string]string{
		"x": pub.X.Text(16),
		"y": pub.Y.Text(16),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
