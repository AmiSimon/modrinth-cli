package main

// modrinth-cli search "rocket" -c physics -c aerospace --match-all -v 1.21.8 -l fabric -s (server, != -C client) -t mod (default)
// modrinth-cli get categories mod (default, == get categories)
// prob later: modrinth-cli get categories shader
// may add a .modrinth-cli in the mods folder for hashes and to identify which is manual and wich isn't
// modrinth-cli unlink (stops searching for jar file version)
// modrinth-cli set-version 1.21.8 (after unlinking)
// modrinth-cli relink
// modrinth-cli upgrade (check version, if different upgrade [Y/n])
// modrinth-cli install sodium
// modrinth-cli uninstall sodium (removes sodium-only deps, asks if ok)

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const Version = "0.0.1"

func ParseArguments() {
	if len(os.Args) < 2 {
		HelpCmd()
		return
	}
	switch os.Args[1] {
	case "help":
		HelpCmd()
	case "-v", "--version":
		PrintVersion(Version)
	case "search":
		PrintSearchResults()
	}

}

func GetFlags(arguments []string) ([][2]string, string) {
	var flags [][2]string
	var keyword string
	skipFlagValue := false

	args := arguments[2:]
	for i := range args {
		if skipFlagValue {
			skipFlagValue = false
			continue
		}
		if !strings.HasPrefix(args[i], "-") && keyword == "" {
			keyword = args[i]
			continue
			// flags = append(flags, {args[i], })
		}
		if strings.HasPrefix(args[i], "-") && strings.HasPrefix(args[i+1], "-") {
			flags = append(flags, [2]string{args[i], ""})
			continue
		}
		if !strings.HasPrefix(args[i], "-") {
			log.Fatalf("Unknown flag %q", args[i])
		}
		flags = append(flags, [2]string{args[i], args[i + 1]})
		skipFlagValue = true
	}
	return flags, keyword
}

func PrintSearchResults() {
	if len(os.Args) < 3 {
		SearchHelp()
		return
	}

}

func GetApiData(url string) ([]byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	request.Header.Set("User-Agent", fmt.Sprintf("AmiSimon/modrinth-cli/%v (simon.leneveu@gmail.com)", Version))

	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	if response.StatusCode == 410 {
		log.Fatal("API Returned 410 Gone, it is deprecated, switch to newer")
	}

	body, err := io.ReadAll(response.Body)
	return body, err
}

func main() {
	ParseArguments()
}
