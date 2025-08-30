package helpers

import (
	"log"
	"os"
	"sync"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var (
	idNamespace uuid.UUID
	idOnce      sync.Once
)

// this method panics if id namespace is not found
func getIDNamespace() uuid.UUID {
	idOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("error loading .env file")
		}

		namespaceString := os.Getenv("ID_NAMESPACE")
		namespaceVal, namespaceErr := uuid.Parse(namespaceString)
		if namespaceErr != nil {
			log.Fatalf("invalid id namespace value found: %s\n%v", namespaceString, namespaceErr)
		}
		idNamespace = namespaceVal
	})

	return idNamespace
}

func GetIDFromPayload(payload []byte) uuid.UUID {
	namespace := getIDNamespace()
	return uuid.NewSHA1(namespace, payload)
}
