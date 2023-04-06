package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	receipt "capn-proto-practice/receipt"

	capnp "zombiezen.com/go/capnproto2"
)

func RunAndVerify(data string) {
	// Receive sample data from args
	sampleData := data

	// Generate a key pair
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	publicKey := &privateKey.PublicKey

	// Hash the data
	hash := sha256.Sum256([]byte(sampleData))

	// Sign the hashed data
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)

	// Create a new Receipt message
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	rcpt, err := receipt.NewRootReceipt(seg)
	if err != nil {
		panic(err)
	}

	rcpt.SetId(1)
	rcpt.SetData([]byte(sampleData))
	rcpt.SetSignature(signature)

	// Serialize the Receipt message
	err = capnp.NewEncoder(os.Stdout).Encode(msg)
	if err != nil {
		panic(err)
	}

	// Save the serialized message to a file
	msgBytes, err := msg.Marshal()
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("serialized_receipt.bin", msgBytes, 0644)
	if err != nil {
		panic(err)
	}

	// Deserialize the Receipt message
	fileData, err := os.ReadFile("serialized_receipt.bin")
	if err != nil {
		panic(err)
	}

	deserializedMsg, err := capnp.Unmarshal(fileData)
	if err != nil {
		panic(err)
	}

	deserializedReceipt, err := receipt.ReadRootReceipt(deserializedMsg)
	if err != nil {
		panic(err)
	}

	id := deserializedReceipt.Id()

	dataBytes, err := deserializedReceipt.Data()
	if err != nil {
		panic(err)
	}
	signatureBytes, err := deserializedReceipt.Signature()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nID: %d\n", id)
	fmt.Printf("Data: %s\n", string(dataBytes))
	fmt.Printf("Signature: %s\n", hex.EncodeToString(signatureBytes))

	// Verify the signature
	rBytes, sBytes := signatureBytes[:len(signatureBytes)/2], signatureBytes[len(signatureBytes)/2:]
	r = new(big.Int).SetBytes(rBytes)
	s = new(big.Int).SetBytes(sBytes)
	isValid := ecdsa.Verify(publicKey, hash[:], r, s)
	fmt.Printf("Signature valid: %v\n", isValid)
}

func main() {
	// get message from user input
	var message string
	fmt.Print("Enter message: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}
	RunAndVerify(message)
}
