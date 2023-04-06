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
	"strconv"
	"strings"

	receipt "capn-proto-practice/receipt"

	capnp "zombiezen.com/go/capnproto2"
)

const stateFile = "chain_state.txt"

// Save state to file
func saveState(id int) {
	err := os.WriteFile(stateFile, []byte(strconv.Itoa(id)), 0644)
	if err != nil {
		panic(err)
	}
}

// Read state from file
func readState() int {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		// If file does not exist, set initial state to 0
		return 0
	}
	id, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		panic(err)
	}
	return id
}

func RunAndVerify(data string, id int) {
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

	rcpt.SetId(uint64(id))
	rcpt.SetData([]byte(sampleData))
	rcpt.SetSignature(signature)

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

	deserializedID := deserializedReceipt.Id()

	dataBytes, err := deserializedReceipt.Data()
	if err != nil {
		panic(err)
	}
	signatureBytes, err := deserializedReceipt.Signature()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nID: %d\n", deserializedID)
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
	// Get message from user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter message: ")
	message, _ := reader.ReadString('\n')

	// Get the current ID from the chain state
	currentID := readState()

	// Increment the ID
	newID := currentID + 1

	// Save the new ID to the chain state
	saveState(newID)

	// Run verification with the updated ID
	RunAndVerify(message, newID)
}
