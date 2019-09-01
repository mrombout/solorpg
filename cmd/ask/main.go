package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mrombout/solorpg/dice"
)

var a1d6 = dice.NumeralDie{
	Faces: 6,
}

var results = map[int]string{
	6: "Yes, and...",
	4: "Yes...",
	2: "Yes, but...",
	5: "No, but...",
	3: "No...",
	1: "No, and...",
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	roll := dice.RollOne(a1d6)
	result := results[roll]

	fmt.Println(result)
}
