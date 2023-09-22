package main

import (
	cmdcfg "github.com/Jchicode/stcevm/cmd/stcevmd/cmd"
	"log"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/Jchicode/stcevm/app"
	"github.com/Jchicode/stcevm/cmd/stcevmd/cmd"
)

func main() {
	cmdcfg.RegisterDenoms()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			log.Println(err)
			os.Exit(1)
		}
	}
}
