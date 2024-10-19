package service

import (
	"encoding/hex"
	"log"
	"os"
)

type BasicAuth struct {
	Login    [32]byte
	Password [32]byte
	IsSet    bool
}

func stringTo32Byte(hashStr string) [32]byte {
	hash2Bytes, err := hex.DecodeString(hashStr)
	if err != nil {
		log.Fatal(err)
	}

	if len(hash2Bytes) != 32 {
		log.Fatal("Hash string is not of the correct length")
	}

	var hash2 [32]byte
	copy(hash2[:], hash2Bytes)
	return hash2
}

func NewBasicAuth() BasicAuth {
	login := os.Getenv("BASIC_AUTH_SHA256_LOGIN")
	passwd := os.Getenv("BASIC_AUTH_SHA256_PASSWORD")
	if login == "" || passwd == "" {
		return BasicAuth{}
	}
	log.Println("Basic auth enabled")
	return BasicAuth{
		IsSet:    true,
		Login:    stringTo32Byte(os.Getenv("BASIC_AUTH_SHA256_LOGIN")),
		Password: stringTo32Byte(os.Getenv("BASIC_AUTH_SHA256_PASSWORD")),
	}
}
