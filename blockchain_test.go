package main

import (
	"bufio"
	"github.com/matryer/is"
	"log"
	"os"
	"strconv"
	"testing"
)

const fileBlockchainTest = "blockchain_test.txt"
const fileBlockchainOtherTest = "blockchain_test_1.txt"

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

func TestRestoreBlockchainFromFile(t *testing.T) {
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
func TestSameHashSameSeedDifferentInstance(t *testing.T) {
	i := is.New(t)
	b, _ := NewBlockchain(fileBlockchainTest)
	bo, _ := NewBlockchain(fileBlockchainOtherTest)

	generateFakeDataBlockchain(b)
	generateFakeDataBlockchain(bo)

	f, err := os.OpenFile(fileBlockchainTest, os.O_RDONLY|os.O_CREATE, 0600)
	fo, err := os.OpenFile(fileBlockchainOtherTest, os.O_RDONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(f)
	defer closeFile(fo)

	scanner := bufio.NewScanner(f)

	var l, lo []string
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	scanner = bufio.NewScanner(fo)
	for scanner.Scan() {
		lo = append(lo, scanner.Text())
	}

	for it, _ := range l {
		i.Equal(l[it], lo[it])
	}

	os.Remove(fileBlockchainTest)
	os.Remove(fileBlockchainOtherTest)
}

func generateFakeDataBlockchain(b *Blockchain) string {
	var lastData string
	for i := 0; i < 1000; i++ {
		b.Add(strconv.Itoa(i))
		lastData = strconv.Itoa(i)
	}
	return lastData
}
