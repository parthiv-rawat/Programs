package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type VerifiableCredential interface {
	GetID() string
	GetType() []string
	// GetPublicKey() string
	GetIssuer() string
	GetIssuanceDate() string
	GetCredentialSubject() interface{}
}

type UniversityDegreeCredential struct {
	VerifiableCredential struct {
		Context           []string `json:"@context"`
		ID                string   `json:"id"`
		Type              []string `json:"type"`
		Issuer            string   `json:"issuer"`
		IssuanceDate      string   `json:"issuanceDate"`
		CredentialSubject struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			GraduationYear int    `json:"graduationYear"`
			Degree         struct {
				Type           string `json:"type"`
				Specialization string `json:"specialization"`
			} `json:"degree"`
		} `json:"credentialSubject"`
		CredentialStatus struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"credentialStatus"`
	} `json:"vc"`
	Issuer        string `json:"iss"`
	NotBeforeTime int    `json:"nbf"`
	JWTid         string `json:"jti"`
	Subject       string `json:"sub"`
	// Proof         struct {
	// 	Type               string `json:"type"`
	// 	Created            string `json:"created"`
	// 	VerificationMethod string `json:"verificationMethod"`
	// 	ProofPurpose       string `json:"proofPurpose"`
	// 	Challenge          string `json:"challenge"`
	// 	Domain             string `json:"domain"`
	// 	ProofValue         string `json:"proofValue"`
	// } `json:"proof"`
}

type SchoolDegreeCredential struct {
	VerifiableCredential struct {
		Context           []string `json:"@context"`
		ID                string   `json:"id"`
		Type              []string `json:"type"`
		Issuer            string   `json:"issuer"`
		IssuanceDate      string   `json:"issuanceDate"`
		CredentialSubject struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			GraduationYear int    `json:"graduationYear"`
			Marksheet      struct {
				Type  string `json:"type"`
				Class string `json:"class"`
			} `json:"marksheet"`
		} `json:"credentialSubject"`
		CredentialStatus struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"credentialStatus"`
	} `json:"vc"`
	Issuer        string `json:"iss"`
	NotBeforeTime int    `json:"nbf"`
	JWTid         string `json:"jti"`
	Subject       string `json:"sub"`
	// Proof         struct {
	// 	Type               string `json:"type"`
	// 	Created            string `json:"created"`
	// 	VerificationMethod string `json:"verificationMethod"`
	// 	ProofPurpose       string `json:"proofPurpose"`
	// 	Challenge          string `json:"challenge"`
	// 	Domain             string `json:"domain"`
	// 	ProofValue         string `json:"proofValue"`
	// } `json:"proof"`
}

type CredentialGenerationRequest struct {
	Issuer      string
	SubjectID   string
	DegreeName  string
	PublicKey   string
	UniqueKey   string
	DID         string
	GeneratedVC chan VerifiableCredential
}

// University
func (c *UniversityDegreeCredential) GetID() string {
	return c.VerifiableCredential.ID
}

func (c *UniversityDegreeCredential) GetType() []string {
	return c.VerifiableCredential.Type
}

func (c *UniversityDegreeCredential) GetIssuer() string {
	return c.VerifiableCredential.Issuer
}

func (c *UniversityDegreeCredential) GetIssuanceDate() string {
	return c.VerifiableCredential.IssuanceDate
}

func (c *UniversityDegreeCredential) GetCredentialSubject() interface{} {
	return c.VerifiableCredential.CredentialSubject
}

// func (c *UniversityDegreeCredential) GetPublicKey() interface{} {
// 	return c.PublicKey
// }

// Schools
func (s *SchoolDegreeCredential) GetID() string {
	return s.VerifiableCredential.ID
}

func (s *SchoolDegreeCredential) GetType() []string {
	return s.VerifiableCredential.Type
}

func (s *SchoolDegreeCredential) GetIssuer() string {
	return s.VerifiableCredential.Issuer
}

func (s *SchoolDegreeCredential) GetIssuanceDate() string {
	return s.VerifiableCredential.IssuanceDate
}

func (s *SchoolDegreeCredential) GetCredentialSubject() interface{} {
	return s.VerifiableCredential.CredentialSubject
}

// func (s *SchoolDegreeCredential) GetPublicKey() interface{} {
// 	return s.PublicKey
// }

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/generate", generateHandler)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("form.html"))
		tmpl.Execute(w, nil)
	}
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		issuer := r.Form.Get("issuer")
		subjectID := r.Form.Get("subjectID")
		degreeName := r.Form.Get("degreeName")
		DID := r.Form.Get("DID")

		if DID == "" {
			http.Error(w, "Missing DID parameter", http.StatusBadRequest)
			return
		}

		uuid := generateUUIDFromDID(DID)

		credentialChan := make(chan VerifiableCredential)

		credentialRequest := CredentialGenerationRequest{
			Issuer:      issuer,
			SubjectID:   subjectID,
			DegreeName:  degreeName,
			DID:         DID,
			UniqueKey:   uuid,
			GeneratedVC: credentialChan,
		}

		if subjectID == "university" {
			go func() {
				generateUniversityCredential(&credentialRequest)
			}()
		} else {
			go func() {
				generateSchoolCredential(&credentialRequest)
			}()
		}

		vc := <-credentialChan

		vcBytes, err := json.MarshalIndent(vc, "", "  ")
		if err != nil {
			http.Error(w, "Error marshaling VerifiableCredential", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(vcBytes)
	}
}

func generateUniversityCredential(request *CredentialGenerationRequest) {
	vc := &UniversityDegreeCredential{
		VerifiableCredential: struct {
			Context           []string `json:"@context"`
			ID                string   `json:"id"`
			Type              []string `json:"type"`
			Issuer            string   `json:"issuer"`
			IssuanceDate      string   `json:"issuanceDate"`
			CredentialSubject struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				GraduationYear int    `json:"graduationYear"`
				Degree         struct {
					Type           string `json:"type"`
					Specialization string `json:"specialization"`
				} `json:"degree"`
			} `json:"credentialSubject"`
			CredentialStatus struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"credentialStatus"`
		}{
			Context: []string{
				"https://www.w3.org/2018/credentials/v1",
				"https://www.w3.org/2018/credentials/examples/v1"},
			ID:           "http://" + request.Issuer + ".edu/credentials/3732",
			Type:         []string{"VerifiableCredential", "UniversityDegreeCredential"},
			Issuer:       request.Issuer,
			IssuanceDate: time.Now().UTC().Format(time.RFC3339),
			CredentialSubject: struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				GraduationYear int    `json:"graduationYear"`
				Degree         struct {
					Type           string `json:"type"`
					Specialization string `json:"specialization"`
				} `json:"degree"`
			}{
				ID:             request.DID,
				Name:           request.Issuer,
				GraduationYear: 2018,
				Degree: struct {
					Type           string `json:"type"`
					Specialization string `json:"specialization"`
				}{
					Type:           "Bachelor's Degree",
					Specialization: request.DegreeName,
				},
			},
			CredentialStatus: struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			}{
				ID:   "https://" + request.Issuer + ".edu/status/24",
				Type: "CredentialStatusList2023",
			},
		},
		Issuer:        request.Issuer,
		NotBeforeTime: 12345,
		JWTid:         "http://" + request.Issuer + ".edu/credentials/3732",
		Subject:       request.DID,
		// Proof: struct {
		// 	Type               string `json:"type"`
		// 	Created            string `json:"created"`
		// 	VerificationMethod string `json:"verificationMethod"`
		// 	ProofPurpose       string `json:"proofPurpose"`
		// 	Challenge          string `json:"challenge"`
		// 	Domain             string `json:"domain"`
		// 	ProofValue         string `json:"proofValue"`
		// }{
		// 	Type:               "Ed25519Signature2020",
		// 	Created:            time.Now().UTC().Format(time.RFC3339),
		// 	VerificationMethod: request.Issuer + "#" + "4338b8fb2e11a6d153ee709ddc948c87a6c73951e22b1ef9c3af68c876c41a66097ef9392b7e741da30e1eddd8a0b0dceefeeb97355cc18f42991a9e969c7a474338b8fb2e11a6d153ee709ddc948c87a6c73951e22b1ef9c3af68c876c41a66097ef9392b7e741da30e1eddd8a0b0dceefeeb97355cc18f42991a9e969c7a47",
		// 	ProofPurpose:       "assertionMethod",
		// 	Challenge:          "1f44d55f-f161-4938-a659-f8026467f126",
		// 	Domain:             "4jt78h47fh47",
		// 	ProofValue:         "z4kWncP1KLByEaSU3oaijUNk8GPGCCgntz8q4Gk55QuMwQe1dsWgSmf7RsRNYgfJUChdSV22khsfpBsX7ub14aYbe",
		// },
	}

	request.GeneratedVC <- vc
}

func generateSchoolCredential(request *CredentialGenerationRequest) {
	vc := &SchoolDegreeCredential{
		VerifiableCredential: struct {
			Context           []string `json:"@context"`
			ID                string   `json:"id"`
			Type              []string `json:"type"`
			Issuer            string   `json:"issuer"`
			IssuanceDate      string   `json:"issuanceDate"`
			CredentialSubject struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				GraduationYear int    `json:"graduationYear"`
				Marksheet      struct {
					Type  string `json:"type"`
					Class string `json:"class"`
				} `json:"marksheet"`
			} `json:"credentialSubject"`
			CredentialStatus struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"credentialStatus"`
		}{
			Context: []string{
				"https://www.w3.org/2018/credentials/v1",
				"https://www.w3.org/2018/credentials/examples/v1"},
			ID:           "http://" + request.Issuer + ".edu/credentials/3732",
			Type:         []string{"VerifiableCredential", "SchoolDegreeCredential"},
			Issuer:       request.Issuer,
			IssuanceDate: time.Now().UTC().Format(time.RFC3339),
			CredentialSubject: struct {
				ID             string `json:"id"`
				Name           string `json:"name"`
				GraduationYear int    `json:"graduationYear"`
				Marksheet      struct {
					Type  string `json:"type"`
					Class string `json:"class"`
				} `json:"marksheet"`
			}{
				ID:             request.DID,
				Name:           request.Issuer,
				GraduationYear: 2018,
				Marksheet: struct {
					Type  string `json:"type"`
					Class string `json:"class"`
				}{
					Type:  "Bachelor's Degree",
					Class: request.DegreeName,
				},
			},
			CredentialStatus: struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			}{
				ID:   "https://" + request.Issuer + ".edu/status/24",
				Type: "CredentialStatusList2023",
			},
		},
		Issuer:        request.Issuer,
		NotBeforeTime: 12345,
		JWTid:         "http://" + request.Issuer + ".edu/credentials/3732",
		Subject:       request.DID,
		// Proof: struct {
		// 	Type               string `json:"type"`
		// 	Created            string `json:"created"`
		// 	VerificationMethod string `json:"verificationMethod"`
		// 	ProofPurpose       string `json:"proofPurpose"`
		// 	Challenge          string `json:"challenge"`
		// 	Domain             string `json:"domain"`
		// 	ProofValue         string `json:"proofValue"`
		// }{
		// 	Type:               "Ed25519Signature2020",
		// 	Created:            time.Now().UTC().Format(time.RFC3339),
		// 	VerificationMethod: request.Issuer + "#" + "4338b8fb2e11a6d153ee709ddc948c87a6c73951e22b1ef9c3af68c876c41a66097ef9392b7e741da30e1eddd8a0b0dceefeeb97355cc18f42991a9e969c7a474338b8fb2e11a6d153ee709ddc948c87a6c73951e22b1ef9c3af68c876c41a66097ef9392b7e741da30e1eddd8a0b0dceefeeb97355cc18f42991a9e969c7a47",
		// 	ProofPurpose:       "assertionMethod",
		// 	Challenge:          "1f44d55f-f161-4938-a659-f8026467f126",
		// 	Domain:             "4jt78h47fh47",
		// 	ProofValue:         request.UniqueKey,
		// },
	}

	request.GeneratedVC <- vc
}

func generateUUIDFromDID(did string) string {
	hash := sha256.Sum256([]byte(did))
	uuid := hex.EncodeToString(hash[:])
	return uuid
}

func generateJWT(vc VerifiableCredential, privateKey []byte) (string, error) {

	vcBytes, err := json.Marshal(vc)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"vc":  string(vcBytes),
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Expiration time: 1 hour
	})

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil

}
