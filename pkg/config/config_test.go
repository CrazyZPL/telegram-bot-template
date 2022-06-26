package config

import (
	"fmt"
	"log"
	"testing"
)

func TestInitConfigFile(t *testing.T) {
	config, err := InitConfigFile("../../config.yaml")
	if err != nil {
		log.Fatalf("Test InitConfigFile error: %v", err)
	}

	fmt.Println(config)
}
