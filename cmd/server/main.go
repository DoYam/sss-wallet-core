package main

import (
	"fmt"
	"log"
	"net/http"

	"sss-wallet-core/core"
	"sss-wallet-core/handlers"
	"sss-wallet-core/internal"
)

func main() {
	if err := core.InitKeyShares(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/sign", internal.WithCORS(handlers.SignHandler))
	http.HandleFunc("/wallet", internal.WithCORS(handlers.WalletHandler))
	http.HandleFunc("/verify", internal.WithCORS(handlers.VerifyHandler))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
