package main

import (
	"log"

	"github.com/dawit_hopes/saas/auth/internal/bootstrap"
)

func main() {
	r, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatalf("failed to start: %v", err)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
