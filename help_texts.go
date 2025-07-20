package main

import (
	"fmt"
	"runtime"
	"strings"
)

// Prints options/subcommands as well as a description. optionLenght is the total lenght of the option name
func PrintOptions(options [][2]string, optionLenght int) {
	stringedSubcommands := ""

	for _,subcommand := range options{
		subcommandPrefix := "  " + subcommand[0]
		stringedSubcommands += subcommandPrefix
		spaces := max(optionLenght - len(subcommandPrefix), 2)
		stringedSubcommands += strings.Repeat(" ", spaces)
		stringedSubcommands += subcommand[1] + "\n"
	}
	fmt.Println(stringedSubcommands)
}

// Prints the help for the CLI, including subcommands
func HelpCmd() {
	subcommands := [][2]string{
		{"search", "Search project, optionnally with filters"},
		{"list", "Lists possible options for project categories, mod loaders etc"},
		{"install", "Installs a project (mod, ressource pack, shader)"},
		{"link", "Links detected minecraft version with modrinth-cli"},
		{"unlink", "Unlinks Minecraft version"},
		{"set-version", "Manually set minecraft version"},
	}
	flags := [][2]string{
		{"-v, --version", "Prints version info"},
	}
	
	helpText := "modrinth-cli, an unofficial WIP mod manager for minecraft servers and clients\n\n" +
	"USAGE\n" +
	"  modrinth-cli <subcommand> [flags]\n\n" +
	"For help in subcommands, add --help after: \"modrinth-cli list --help\"\n\n" +
	"SUBCOMMANDS"

	
	fmt.Println(helpText)
	PrintOptions(subcommands, 20)
	fmt.Println("FLAGS")
	PrintOptions(flags, 20)
}

func PrintVersion(version string) {
	goVersion := runtime.Version()
	fmt.Println(version + " " + goVersion)
}