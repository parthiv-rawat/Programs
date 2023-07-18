package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

type UserDocument struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
}

func main() {
	// Simulated user document data
	userDocData := `{
		"id": "12345",
		"name": "Parthiv Rawat",
		"publicKey": "4338b8fb2e11a6d153ee709ddc948c87a6c73951e22b1ef9c3af68c876c41a66097ef9392b7e741da30e1eddd8a0b0dceefeeb97355cc18f42991a9e969c7a47",
		"signature": "0e3a01f83521d87ec6b7f0d6b73648a18976e07f45695c1259e84b2a3b7b7d58e8cc607f07525f2b7d7d3a03e5a1d01b9012356b605b2d19a9e41d731ee0bce"
	}`

	// Parse the user document JSON
	userDoc := UserDocument{}
	if err := json.Unmarshal([]byte(userDocData), &userDoc); err != nil {
		log.Fatal(err)
	}

	// Verify the authenticity of the user document
	isAuthentic, err := VerifyUserDocument(&userDoc)
	if err != nil {
		log.Fatal(err)
	}

	if isAuthentic {
		fmt.Println("User document is authentic")
	} else {
		fmt.Println("User document is not authentic")
	}
}

func VerifyUserDocument(userDoc *UserDocument) (bool, error) {

	publicKeyBytes, err := hex.DecodeString(padHex(userDoc.PublicKey))
	if err != nil {
		return false, err
	}

	signatureBytes, err := hex.DecodeString(padHex(userDoc.Signature))
	if err != nil {
		return false, err
	}

	message := []byte(fmt.Sprintf("%s:%s", userDoc.ID, userDoc.Name))

	isValid := ecdsa.Verify(decodePublicKey(publicKeyBytes), message, decodeSignature(signatureBytes[0:16]), decodeSignature(signatureBytes[16:]))
	return isValid, nil
}

func padHex(hexString string) string {
	if len(hexString)%2 != 0 {
		return "0" + hexString
	}
	return hexString
}

func decodePublicKey(publicKeyBytes []byte) *ecdsa.PublicKey {
	curve := elliptic.P256()
	x, y := elliptic.Unmarshal(curve, publicKeyBytes)
	return &ecdsa.PublicKey{
		Curve: curve,
		X:     x,
		Y:     y,
	}
}

func decodeSignature(signatureBytes []byte) *big.Int {
	return new(big.Int).SetBytes(signatureBytes)
}
