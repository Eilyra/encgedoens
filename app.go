package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	key := make([]byte, 16)
	iv := make([]byte, 16)
	rand.Read(iv)
	kf, err := os.Open("C:/Users/Ellen/film.key")
	defer kf.Close()
	if err != nil {
		panic(err)
	}
	_, err = io.ReadFull(kf, key)
	if err != nil {
		panic(err)
	}
	ivf, err := os.Create("C:/Users/Ellen/lipsum.iv")
	if err != nil {
		panic(err)
	}
	n, err := ivf.Write(iv)
	if err != nil {
		panic(err)
	}
	if n != len(iv) {
		panic("Didn't write the whole IV!")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	c := newCTR(block, iv)

	inf, err := os.Open("C:/Users/Ellen/lipsum.txt")
	defer inf.Close()
	if err != nil {
		panic(err)
	}
	ouf, err := os.Create("C:/Users/Ellen/lipsum.enc")
	if err != nil {
		panic(err)
	}
	var counter uint64 = 0
	for {
		buf := make([]byte, 16)
		n, err := inf.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error: %s", err)
			}
			break
		}
		encrypted := c.XOR(buf, counter)[:n]
		on, err := ouf.Write(encrypted)
		if err != nil {
			panic(err)
		}
		if n != on {
			fmt.Printf("In bytes not equal to out, %d, %d\n", n, on)
		}
		counter++
	}
	ouf.Close()
	counter = 0
	ouf, err = os.Open("C:/Users/Ellen/lipsum.enc")
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, 16)
		n, err := ouf.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		fmt.Printf("%s", c.XOR(buf, counter)[:n])
		counter++
	}
}
