package rollfun

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mrombout/solorpg/dice"
)

// Roll rolls one or more dice based on the given dice notation and returns the result.
func Roll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	queryParams := r.URL.Query()

	diceNotation := queryParams.Get("dice")
	if diceNotation == "" {
		diceNotation = "1d6"
	}

	seedParam := queryParams.Get("seed")
	var seed int64
	if seedParam == "" {
		seed = time.Now().UTC().UnixNano()
	} else {
		var err error
		seed, err = strconv.ParseInt(seedParam, 10, 64)
		if err != nil {
			fmt.Println("seed not valid")
			os.Exit(1)
		}
	}

	result, err := roll(diceNotation, seed)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}

func roll(diceNotation string, seed int64) (RollResult, error) {
	dice, err := dice.Parse(diceNotation)
	if err != nil {
		return emptyResult, err
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for key := range dice {
		die := &dice[key]
		die.Roll()
	}

	totalRoll := 0
	for _, dice := range dice {
		totalRoll += dice.Result
	}

	return RollResult{
		Result: totalRoll,
		Dice:   dice,
	}, nil
}

// RollResult contains the result of a single dice roll.
type RollResult struct {
	Result int
	Dice   []dice.NumeralDie
}

var emptyResult = RollResult{}

type rollRequest struct {
	DiceNotation string
}

type rollResponse struct {
	Result int
	Dice   []dice.NumeralDie
	Err    string
}
