package main

// Block Defines a single block
type Block struct {

	// Blockchain header parts
	Hash          []byte
	PrevBlockHash []byte
	Timestamp     int64

	// Data part
	Data []byte
}

func main() {

}
