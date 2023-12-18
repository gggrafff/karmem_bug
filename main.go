package main

import (
	"example/karmem_bug/serialization_api"
	"fmt"

	karmem "karmem.org/golang"
)

func main() {
	fmt.Println("Test of karmem")
	msg1 := serialization_api.NewBtype()
	msg1.A.IntField = 42
	msg1.A.BoolField = true
	msg1.A.UintField = 33
	writer := karmem.NewWriter(256)
	_, err := msg1.WriteAsRoot(writer)
	if err != nil {
		panic(err)
	}
	encoded := writer.Bytes()
	fmt.Printf("%v was encoded to %d bytes\n", msg1, len(encoded))

	reader := karmem.NewReader(encoded)
	msg2 := serialization_api.NewBtype()
	msg2.ReadAsRoot(reader)
	fmt.Printf("%d bytes was decoded to %v\n", len(encoded), msg2)

	if msg1.A.IntField != msg2.A.IntField {
		panic(fmt.Sprintf("%d == msg1.A.IntField != msg2.A.IntField == %d", msg1.A.IntField, msg2.A.IntField))
	}
}
