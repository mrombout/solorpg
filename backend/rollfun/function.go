package rollfun

import (
	"encoding/json"
	"fmt"
	"github.com/mrombout/solorpg/backend/rollsvc"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Roll rolls one or more dice based on the given dice notation and returns the result.
func Roll(w http.ResponseWriter, r *http.Request) {
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

	rollService := rollsvc.RollServiceImpl{}
	result, err := rollService.Roll(diceNotation, seed)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json.NewEncoder(w).Encode(result)
}
