package main

import (
	"bufio"
	"crypto/sha512"
	"encoding/hex"
	"github.com/en-vee/alog"
	"log"
	"os"
	"strings"
)

type Blockchain struct {
	filename string
	last     *Chain
}

func NewBlockchain(filename string) (*Blockchain, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(f)
	scanner := bufio.NewScanner(f)

	var previous, current *Chain
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		data, id := line[0], line[1]
		if previous == nil {
			previous = &Chain{id, data, nil}
			continue
		}
		current = &Chain{id, data, previous}
		previous = current
	}

	return &Blockchain{filename: filename, last: current}, nil
}

func (b *Blockchain) Save() error {
	f, err := os.OpenFile(b.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		alog.Error(err.Error())
	}

	defer closeFile(f)

	if _, err := f.WriteString(b.last.data + "|" + b.last.id + "\r\n"); err != nil {
		return err
	}

	return nil
}

func (b *Blockchain) Add(data string) error {
	var previousId string
	if nil == b.last {
		previousId = "0" //generate fixed seed?
	} else {
		previousId = (b.last).ID()
	}

	h := sha512.New()
	h.Write([]byte(previousId))
	b.last = newChain(hex.EncodeToString(h.Sum(nil)), data, b.last)
	return b.Save()
}

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		alog.Error(err.Error())
	}
}
