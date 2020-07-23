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
	command = &cobra.Command{
		Use: "groo",
		Short: "run in client mode",
		Long: "",
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
	fmt.Printf(aurora.Red("\n■■▶︎ Aborting without taking further action\n").String())
}

func execute(cmd *cobra.Command, args []string) {
	err := commands.OpenRemote(args[0])
	checkError(err)
}