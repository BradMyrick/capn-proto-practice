package main

import (
	"fmt"
	"io/ioutil"
	"os"

	receipt "capn-proto-practice/receipt/receipt"

	capnp "zombiezen.com/go/capnproto2"
)

func main() {
	// Create a new Receipt message
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	r, err := receipt.NewRootReceipt(seg)
	if err != nil {
		panic(err)
	}

	r.SetId(1)
	r.SetData([]byte("example data"))
	r.SetSignature([]byte("example signature"))

	// Serialize the Receipt message
	err = capnp.NewEncoder(os.Stdout).Encode(msg)
	if err != nil {
		panic(err)
	}

	// Save the serialized message to a file
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("serialized_receipt.bin", data, 0644)
	if err != nil {
		panic(err)
	}

	// Deserialize the Receipt message
	fileData, err := ioutil.ReadFile("serialized_receipt.bin")
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

	fmt.Printf("ID: %d\n", deserializedReceipt.Id())
	fmt.Printf("Data: %s\n", string(deserializedReceipt.Data()))
	fmt.Printf("Signature: %s\n", string(deserializedReceipt.Signature()))
}
