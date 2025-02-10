// // package main

// // import (
// // 	"encoding/base64"
// // 	"fmt"
// // 	"io/ioutil"

// // 	"github.com/docusign/docusign-esign-go/model"
// // )

// // // const codeVerifier string = "R8zFoqs0yey29G71QITZs3dK1YsdIvFNBfO4D1bukBw"

// // // func GenerateCodeChallenge() (string, error) {
// // // 	hash := sha256.Sum256([]byte(codeVerifier))
// // // 	codeChallenge := base64.RawURLEncoding.EncodeToString(hash[:])

// // // 	return codeChallenge, nil
// // // }

// // func makeEnvelope(args map[string]string, docPath string) (*model.EnvelopeDefinition, error) {
// // 	// Ler o arquivo do documento
// // 	docBytes, err := ioutil.ReadFile(docPath)
// // 	if err != nil {
// // 		return nil, fmt.Errorf("erro ao ler %s: %w", docPath, err)
// // 	}
// // 	docB64 := base64.StdEncoding.EncodeToString(docBytes)

// // 	// Criar o documento
// // 	document := model.Document{
// // 		DocumentBase64: &docB64,
// // 		Name:           "Contrato de Assinatura",
// // 		FileExtension:  "pdf", // Altere conforme o tipo do arquivo
// // 		DocumentId:     "1",
// // 	}

// // 	// Criar o destinatário (assinante)
// // 	signer := model.Signer{
// // 		Email:        args["signer_email"],
// // 		Name:         args["signer_name"],
// // 		RecipientId:  "1",
// // 		RoutingOrder: "1",
// // 		Tabs: &model.Tabs{
// // 			SignHereTabs: &[]model.SignHere{
// // 				{
// // 					AnchorString:  "/assinatura/",
// // 					AnchorUnits:   "pixels",
// // 					AnchorXOffset: "20",
// // 					AnchorYOffset: "10",
// // 				},
// // 			},
// // 		},
// // 	}

// // 	// Criar o envelope
// // 	env := &model.EnvelopeDefinition{
// // 		EmailSubject: "Por favor, assine este documento",
// // 		Status:       args["status"], // "sent" para envio, "created" para rascunho
// // 		Documents:    &[]model.Document{document},
// // 		Recipients: &model.Recipients{
// // 			Signers: &[]model.Signer{signer},
// // 		},
// // 	}

// // 	return env, nil
// // }

// // func main() {
// // 	// codeChallenge, err := GenerateCodeChallenge()
// // 	// if err != nil {
// // 	// 	panic(err)
// // 	// }

// // 	// fmt.Println(codeChallenge)
// // }

// // /*
// // https://account-d.docusign.com/oauth/auth?response_type=code&scope=signature
// // &client_id=6e1a57cc-ffa4-4f61-a633-c47ee6550678
// // &redirect_uri=http://localhost/
// // &code_challenge_method=S256
// // &code_challenge=RqN6kvc2fJwD-BQG3SzsDfQcX54BxuyuM40alAt8b5M
// // */

// package main

// import (
// 	"bytes"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.POST("/upload", func(c *gin.Context) {
// 		body, err := io.ReadAll(c.Request.Body)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Falha ao ler o corpo da requisição",
// 			})
// 			return
// 		}

// 		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

// 		log.Println("Corpo da requisição recebido:", string(body))

// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Corpo da requisição recebido e logado com sucesso",
// 		})
// 	})

//		if err := r.Run(":8888"); err != nil {
//			log.Fatal("Erro ao iniciar o servidor:", err)
//		}
//	}
package main
