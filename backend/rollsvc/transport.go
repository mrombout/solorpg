package rollsvc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/mrombout/solorpg/dice"
)

// MakeRollEndpoint creates a Roll endpoint for gRPC.
func MakeRollEndpoint(service RollService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		rollRequest := request.(rollRequest)
		result, err := service.Roll(rollRequest.DiceNotation, 0) // TODO: Use proper seed
		if err != nil {
			return rollResponse{0, []dice.NumeralDie{}, err.Error()}, nil
		}
		return rollResponse{result.Result, result.Dice, ""}, nil
	}
}

// DecodeRollRequest decodes a roll request to from HTTP.
func DecodeRollRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request rollRequest

	diceNotation, ok := r.URL.Query()["diceNotation"]
	if !ok || len(diceNotation[0]) < 1 {
		return nil, errors.New("diceNotation parameter not given")
	}

	request.DiceNotation = diceNotation[0]

	return request, nil
}

// EncodeResponse encodes a roll response to JSON.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type rollRequest struct {
	DiceNotation string
}

type rollResponse struct {
	Result int
	Dice   []dice.NumeralDie
	Err    string
}
