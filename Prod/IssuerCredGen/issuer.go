package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type VerifiableCredential interface {
	GetContext() []string
	GetID() string
	GetType() []string
	GetIssuer() string
	GetIssuanceDate() string
	GetCredentialSubject() interface{}
}

type Config struct {
	PrivateKey string `json:"PrivateKey"`
	Port       string `json:"Port"`
}

type BaseCredential struct {
	VerifiableCredential struct {
		Context           []string    `json:"@context"`
		ID                string      `json:"id"`
		Type              []string    `json:"type"`
		Issuer            string      `json:"issuer"`
		IssuanceDate      string      `json:"issuanceDate"`
		CredentialSubject interface{} `json:"credentialSubject"`
		CredentialStatus  struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"credentialStatus"`
	} `json:"vc"`
	Issuer        string `json:"iss"`
	NotBeforeTime int    `json:"nbf"`
	JWTid         string `json:"jti"`
	Subject       string `json:"sub"`
}

type UniversityDegreeCredential struct {
	BaseCredential
}

type SchoolDiplomaCredential struct {
	BaseCredential
}

type UniversityDegreeTemplate struct {
	DegreeType string `json:"degreeType"`
	Program    string `json:"program"`
	Department string `json:"department"`
	Graduation int    `json:"graduation"`
}

type SchoolDiplomaTemplate struct {
	Stream     string `json:"stream"`
	Graduation int    `json:"graduation"`
}

type CredentialGenerationRequest struct {
	Issuer      string
	Template    interface{}
	DID         string
	GeneratedVC chan VerifiableCredential
}

func (c *UniversityDegreeCredential) GetContext() []string {
	return c.BaseCredential.VerifiableCredential.Context
}

func (c *UniversityDegreeCredential) GetID() string {
	return c.BaseCredential.VerifiableCredential.ID
}

func (c *UniversityDegreeCredential) GetType() []string {
	return c.BaseCredential.VerifiableCredential.Type
}

func (c *UniversityDegreeCredential) GetIssuer() string {
	return c.BaseCredential.VerifiableCredential.Issuer
}

func (c *UniversityDegreeCredential) GetIssuanceDate() string {
	return c.BaseCredential.VerifiableCredential.IssuanceDate
}

func (c *UniversityDegreeCredential) GetCredentialSubject() interface{} {
	return c.BaseCredential.VerifiableCredential.CredentialSubject
}

func (s *SchoolDiplomaCredential) GetContext() []string {
	return s.BaseCredential.VerifiableCredential.Context
}

func (s *SchoolDiplomaCredential) GetID() string {
	return s.BaseCredential.VerifiableCredential.ID
}

func (s *SchoolDiplomaCredential) GetType() []string {
	return s.BaseCredential.VerifiableCredential.Type
}

func (s *SchoolDiplomaCredential) GetIssuer() string {
	return s.BaseCredential.VerifiableCredential.Issuer
}

func (s *SchoolDiplomaCredential) GetIssuanceDate() string {
	return s.BaseCredential.VerifiableCredential.IssuanceDate
}

func (s *SchoolDiplomaCredential) GetCredentialSubject() interface{} {
	return s.BaseCredential.VerifiableCredential.CredentialSubject
}

func loadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	err = json.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	// Create a CredentialGenerationRequest for UniversityDegreeCredential
	universityRequest := &CredentialGenerationRequest{
		Issuer: "Example University",
		Template: UniversityDegreeTemplate{
			DegreeType: "Bachelor's Degree",
			Program:    "Computer Science",
			Department: "Engineering",
			Graduation: 2023,
		},
		DID: "did:example:12345",
	}

	// Generate UniversityDegreeCredential
	universityCredential := generateUniversityCredential(universityRequest)
	fmt.Println("University Degree Credential:")
	fmt.Printf("%+v\n", universityCredential)

	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	privateKey := []byte(config.PrivateKey)

	jwtTokenUniversity, err := generateJWT(universityCredential, privateKey)
	if err != nil {
		log.Fatal("Failed to generate JWT token:", err)
	}

	fmt.Println("JWT Token:")
	fmt.Println(jwtTokenUniversity)
	fmt.Println()

	// Create a CredentialGenerationRequest for SchoolDiplomaCredential
	schoolRequest := &CredentialGenerationRequest{
		Issuer: "Example School",
		Template: SchoolDiplomaTemplate{
			Stream:     "Science",
			Graduation: 2023,
		},
		DID: "did:example:54321",
	}

	// Generate SchoolDiplomaCredential
	schoolCredential := generateSchoolCredential(schoolRequest)
	fmt.Println("School Diploma Credential:")
	fmt.Printf("%+v\n", schoolCredential)

	jwtTokenSchool, err := generateJWT(schoolCredential, privateKey)
	if err != nil {
		log.Fatal("Failed to generate JWT token:", err)
	}

	fmt.Println("JWT Token:")
	fmt.Println(jwtTokenSchool)
}

func generateCredential(request *CredentialGenerationRequest, credentialType string) VerifiableCredential {
	baseCredential := BaseCredential{
		VerifiableCredential: struct {
			Context           []string    `json:"@context"`
			ID                string      `json:"id"`
			Type              []string    `json:"type"`
			Issuer            string      `json:"issuer"`
			IssuanceDate      string      `json:"issuanceDate"`
			CredentialSubject interface{} `json:"credentialSubject"`
			CredentialStatus  struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			} `json:"credentialStatus"`
		}{
			Context:      []string{"https://www.w3.org/2018/credentials/v1"},
			ID:           "http://" + request.Issuer + ".edu/credentials/" + string(rand.Intn(1000)),
			Type:         []string{"VerifiableCredential", credentialType},
			Issuer:       request.Issuer,
			IssuanceDate: time.Now().UTC().Format(time.RFC3339),
			CredentialStatus: struct {
				ID   string `json:"id"`
				Type string `json:"type"`
			}{
				ID:   "https://" + request.Issuer + ".edu/status/" + string(rand.Intn(1000)),
				Type: "CredentialStatusList2023",
			},
		},
		Issuer:        request.Issuer,
		NotBeforeTime: rand.Intn(10000),
		JWTid:         "http://" + request.Issuer + ".edu/credentials/" + string(rand.Intn(1000)),
		Subject:       request.DID,
	}

	switch credentialType {
	case "UniversityDegreeCredential":
		template := request.Template.(UniversityDegreeTemplate)
		baseCredential.VerifiableCredential.CredentialSubject = struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			GraduationYear int    `json:"graduationYear"`
			Degree         struct {
				Type       string `json:"type"`
				Program    string `json:"program"`
				Department string `json:"department"`
			} `json:"degree"`
		}{
			ID:             request.DID,
			Name:           request.Issuer,
			GraduationYear: template.Graduation,
			Degree: struct {
				Type       string `json:"type"`
				Program    string `json:"program"`
				Department string `json:"department"`
			}{
				Type:       template.DegreeType,
				Program:    template.Program,
				Department: template.Department,
			},
		}
	case "SchoolDiplomaCredential":
		template := request.Template.(SchoolDiplomaTemplate)
		baseCredential.VerifiableCredential.CredentialSubject = struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			GraduationYear int    `json:"graduationYear"`
			Stream         string `json:"stream"`
		}{
			ID:             request.DID,
			Name:           request.Issuer,
			GraduationYear: template.Graduation,
			Stream:         template.Stream,
		}
	default:
		log.Fatalf("Unsupported credential type: %s", credentialType)
		return nil
	}

	switch credentialType {
	case "UniversityDegreeCredential":
		return &UniversityDegreeCredential{BaseCredential: baseCredential}
	case "SchoolDiplomaCredential":
		return &SchoolDiplomaCredential{BaseCredential: baseCredential}
	default:
		log.Fatalf("Unsupported credential type: %s", credentialType)
		return nil
	}
}

func generateUniversityCredential(request *CredentialGenerationRequest) VerifiableCredential {
	return generateCredential(request, "UniversityDegreeCredential")
}

func generateSchoolCredential(request *CredentialGenerationRequest) VerifiableCredential {
	return generateCredential(request, "SchoolDiplomaCredential")
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
