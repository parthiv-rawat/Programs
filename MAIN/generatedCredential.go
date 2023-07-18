/*
1. Define the requirements in which
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", generateUUIDHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateUUIDHandler(w http.ResponseWriter, r *http.Request) {
	did := r.URL.Query().Get("did")

	if did == "" {
		http.Error(w, "Missing DID parameter", http.StatusBadRequest)
		return
	}

	uuid := generateUUIDFromDID(did)

	fmt.Fprintf(w, "Generated UUID: %s", uuid)
}

func generateUUIDFromDID(did string) string {
	hash := sha256.Sum256([]byte(did))
	uuid := hex.EncodeToString(hash[:])
	return uuid
}
