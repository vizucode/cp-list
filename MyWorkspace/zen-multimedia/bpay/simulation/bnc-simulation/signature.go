package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func String2Sign(method, endpointUrl, timestamp, payload string) (resp string, err error) {
	payloadHashed, err := GenerateSHA256Encoded(os.Getenv("BNC_PRIVATE_KEY"), payload)
	if err != nil {
		log.Println(err)
		return "", err
	}
	resp = fmt.Sprintf("%s:%s:%s:%s", method, endpointUrl, strings.ToLower(payloadHashed), timestamp)

	return resp, nil
}

func LoadPubKey() (pubKey *rsa.PublicKey, err error) {
	pubKeyPem := fmt.Sprintf("-----BEGIN RSA PUBLIC KEY-----\n%s\n-----END RSA PUBLIC KEY-----", os.Getenv("BNC_PUBLIC_KEY"))

	r := strings.NewReader(pubKeyPem)
	pbKeyBytes, err := io.ReadAll(r)
	if err != nil {
		log.Println(err)
		return pubKey, err
	}

	// decode the pem encoding
	block, _ := pem.Decode(pbKeyBytes)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		log.Println(err)
		return pubKey, err
	}

	pubKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Println(err)
		return pubKey, err
	}

	return pubKey, nil
}

func VerifySignature(signature string, message string) (err error) {

	pubKey, err := LoadPubKey()
	if err != nil {
		log.Println(err)
		return err
	}

	messageEncrypted, err := GenerateSHA256(os.Getenv("BNC_PRIVATE_KEY"), message)
	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, []byte(messageEncrypted), []byte(signature))
	if err != nil {
		return err
	}

	return nil
}

func GenerateSHA256(privateKey string, payload string) (resp string, err error) {
	hash := sha256.New()
	hash.Write([]byte(payload))
	resp = string(hash.Sum(nil))
	return resp, nil
}

func GenerateSHA256Encoded(privateKey string, payload string) (resp string, err error) {
	hash := sha256.New()
	hash.Write([]byte(payload))
	resp = hex.EncodeToString(hash.Sum(nil))
	return resp, nil
}
