package main

import (
	"github.com/matryer/is"
	"os"
	"strconv"
	"testing"
)

const fileBlockchainTest = "blockchain_test.txt"

func TestCreateBlockchainToFile(t *testing.T) {
	i := is.New(t)
	b, err := NewBlockchain(fileBlockchainTest)
	if err != nil {
		t.Errorf("Error creating blockchain: %v", err)
	}

	lastData := generateFakeDataBlockchain(b)

	i.Equal(b.last.data, lastData)
	os.Remove(fileBlockchainTest)
}

func TestRestoreBlochainFromFile(t *testing.T) {
	i := is.New(t)
	b, err := NewBlockchain(fileBlockchainTest)
	if err != nil {
		t.Errorf("Error creating blockchain: %v", err)
	}
	generateFakeDataBlockchain(b)

	bFromFile, err := NewBlockchain(fileBlockchainTest)
	i.Equal(b.last.data, bFromFile.last.data)
	i.Equal(b.last.id, bFromFile.last.id)
	i.Equal(b.last.previous.id, bFromFile.last.previous.id)
	i.Equal(b.last.previous.data, bFromFile.last.previous.data)

	os.Remove(fileBlockchainTest)
}

func generateFakeDataBlockchain(b *Blockchain) string {
	var lastData string
	for i := 0; i < 1000; i++ {
		b.Add(strconv.Itoa(i))
		lastData = strconv.Itoa(i)
	}
	return lastData
}
