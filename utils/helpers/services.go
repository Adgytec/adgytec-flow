package helpers

import (
	"log"
	"os"

	"github.com/google/uuid"
)

// this method panics if service namespace is not found
func getServiceNamespace() uuid.UUID {
	namespaceString := os.Getenv("SERVICE_NAMESPACE")
	namespaceVal, namespaceErr := uuid.Parse(namespaceString)
	if namespaceErr != nil {
		log.Fatalf("invalid service namespace value found: %s", namespaceString)
	}

	return namespaceVal
}

func GetServiceIdFromServiceName(name string) uuid.UUID {
	namespace := getServiceNamespace()
	return uuid.NewSHA1(namespace, []byte(name))
}
