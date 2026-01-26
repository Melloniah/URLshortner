package shortener

import (
	"math/rand"
	"time"
)

const (
	// Characters used in short codes
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Default length of short codes
	codeLength = 6
)

// Generator generates short codes
type Generator struct {
	rnd *rand.Rand
}

// NewGenerator creates a new code generator
func NewGenerator() *Generator {
	// Create a new random source with current time as seed
	source := rand.NewSource(time.Now().UnixNano())
	return &Generator{
		rnd: rand.New(source),
	}
}

// Generate creates a random short code
func (g *Generator) Generate() string {
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[g.rnd.Intn(len(charset))]
	}
	return string(code)
}