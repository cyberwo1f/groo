package main

import (
	"fmt"
	"github.com/cyberfox1002/groo/commands"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

var (
	target string
	command = &cobra.Command{
		Use: "groo",
		Short: "run to open the web site for git remote. the default target is the origin.",
		Version: Version,
		Run: execute,
	}
)

func init() {
	command.Flags().StringVarP(&target, "remote", "r", "origin", "select of the target name.")
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
	err := commands.OpenRemote(target)
	checkError(err)
}