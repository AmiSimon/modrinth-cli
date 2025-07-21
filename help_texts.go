package main

import (
	"fmt"
	"runtime"
	"strings"
)

// Prints options/subcommands as well as a description. optionLenght is the total lenght of the option name
func AssembleOptions(options [][2]string, optionLenght int) string {
	stringedSubcommands := ""

	for _,subcommand := range options{
		subcommandPrefix := "  " + subcommand[0]
		stringedSubcommands += subcommandPrefix
		spaces := max(optionLenght - len(subcommandPrefix), 2)
		stringedSubcommands += strings.Repeat(" ", spaces)
		stringedSubcommands += subcommand[1] + "\n"
	}
	return stringedSubcommands
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
	fmt.Print(AssembleOptions(subcommands, 20))
	fmt.Println("\nFLAGS")
	fmt.Print(AssembleOptions(flags, 20))
}

func PrintVersion(version string) {
	goVersion := runtime.Version()
	fmt.Println(version + " " + goVersion)
}

func SearchHelp() {
	flags := [][2]string{
		{"-c, --category", "Search for mods with a specific category"},
		{"--match-any", "Searches for mods matching one of the categories, instead of all of them"},
		{"-v, --mod-version", "Searches for a specific version. By default the version set for the instance"},
		{"-l, --loader", "Searches for a specific loader. By default the instance's loader"},
		{"-s, --server-side", "Searches for mods that have to work on servers"},
		{"-C, --client-side", "Searches for mods that have to work on clients"},
		{"-t, --project-type", "Searches for a project type, Default: mod"},
	}
	fmt.Print(
	"Searches for a project on Modrinth\n\n" +
	"USAGE\n"+ 
	"  modrinth-cli search <keywords> [flags]\n\n" +
	"FLAGS\n")
	fmt.Print(AssembleOptions(flags, 30))
}