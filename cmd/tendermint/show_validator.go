package main

import (
	"fmt"

	data "github.com/tendermint/go-data"
	"github.com/tendermint/tendermint/types"
)

func show_validator() {
	privValidatorFile := config.GetString("priv_validator_file")
	privValidator := types.LoadOrGenPrivValidator(privValidatorFile)
	pubKeyJSONBytes, _ := data.ToJSON(privValidator.PubKey)
	fmt.Println(string(pubKeyJSONBytes))
}
