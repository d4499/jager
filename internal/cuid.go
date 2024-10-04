package internal

import (
	"log"
	"math/rand"

	"github.com/nrednav/cuid2"
)

func NewCUID() string {
	generate, err := cuid2.Init(
		cuid2.WithRandomFunc(rand.Float64),
		cuid2.WithLength(30),
	)
	if err != nil {
		log.Fatalf("Unable to generate cuid")
	}

	id := generate()

	return id
}
