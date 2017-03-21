package main

import (
	"fmt"

	data "github.com/tendermint/go-data"
	"github.com/tendermint/tendermint/types"
)

func gen_validator() {
	privValidator := types.GenPrivValidator()
	privValidatorJSONBytes, _ := data.ToJSON(privValidator)
	fmt.Println(string(privValidatorJSONBytes))
}
