package main

import (
	"fmt"
	"github.com/cyberfox1002/groo/commands"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	version = "v0.5.0"

	command = &cobra.Command{
		Use: "groo [remote]",
		Short: "run to open the web site for git remote.",
		Version: version,
		Run: execute,
		Args: cobra.MinimumNArgs(1),
	}
)

func init() {
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
}

func checkError(err error) {
	if err != nil {
		displayError(err)
		os.Exit(1)
	}
}

func displayError(err error) {
	fmt.Printf(aurora.Red("\n■■▶︎ ERROR: %s").String(), err.Error())
}

func execute(cmd *cobra.Command, args []string) {
	err := commands.OpenRemote(args[0])
	checkError(err)
}