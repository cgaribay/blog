package main

import (
	"os"

	"github.com/cgaribay/blog/app"
	"github.com/cgaribay/blog/cmd/blogd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
