package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
	"abcdefghijklmnopqrstuvwxyzåäö" +
	"0123456789" +
	"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

// GeneratePasswords generates a password with the given length
func GeneratePassword(ln int) string {
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < ln; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// GeneratePasswords is a cycle for generating a batch of passwords
func GeneratePasswords(qnt, ln int) []string {
	resp := make([]string, qnt, qnt)
	for i := 0; i < qnt; i++ {
		resp[i] = GeneratePassword(ln)
	}
	// return "q65hvP4&WnQZ#s"
	return resp
}

// main takes params and generates a password with the given length and quantity of examples.
// TODO: take args.
// If no input is given, all the default values are used.
func main() {
	passLen := 14
	qntOfPasses := 10

	fmt.Println("Generating ...")
	for _, v := range GeneratePasswords(qntOfPasses, passLen) {
		fmt.Println(v)
	}
}
