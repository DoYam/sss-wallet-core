package main

import (
	"fmt"
	"log"
	"net/http"

	"sss-wallet-core/core"
	"sss-wallet-core/handlers"
)

func main() {
	if err := core.InitKeyShares(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/sign", handlers.SignHandler)
	http.HandleFunc("/wallet", handlers.WalletHandler)
	http.HandleFunc("/verify", handlers.VerifyHandler)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
