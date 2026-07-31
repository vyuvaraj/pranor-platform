package main

import (
	"fmt"
	"os"

	"github.com/vyuvaraj/PranorCore"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run upload_openapi.go <service-name> <spec-file-path>")
		os.Exit(1)
	}

	serviceName := os.Args[1]
	filePath := os.Args[2]

	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading spec file: %v\n", err)
		os.Exit(1)
	}

	sc := PranorCore.NewStoreClient()
	bucket := "openapi-registry"
	key := fmt.Sprintf("%s-openapi.json", serviceName)

	fmt.Printf("Uploading %s spec to PranorVault bucket %s/%s...\n", serviceName, bucket, key)
	if err := sc.Put(bucket, key, data); err != nil {
		fmt.Printf("Error uploading spec: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Spec uploaded successfully!")
}
