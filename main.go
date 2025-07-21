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
// modrinth-cli uninstall (--only) sodium (removes sodium-only deps, asks if ok)

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const Version = "0.0.1"

const StagingApi = "https://staging-api.modrinth.com/v2/"
const MainApi = "https://api.modrinth.com/v2/"

type QueryParameter struct {
	Key   string
	Value string
}

type Flag struct {
	Flag  string
	Value string
}

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

func GetFlags(arguments []string) ([]Flag, string) {
	var flags []Flag
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
			flags = append(flags, Flag{
				args[i],
				"",
			})
			continue
		}
		if !strings.HasPrefix(args[i], "-") {
			log.Fatalf("Unknown flag %q", args[i])
		}

		flags = append(flags, Flag{
			args[i], 
			args[i+1],
		})

		skipFlagValue = true
	}
	return flags, keyword
}

func FlagsToFacets(flags []Flag) string {
	var categories []string
	var facets, loader, projectType, version string
	var isMatchAny bool //, isServerSide, isClientSide bool
	//TODO figure out wtf is going on with the bool facets shitshow
	for _, flag := range flags {
		switch flag.Flag {
		case "-c", "--category":
			categories = append(categories, flag.Value) // TODO: Check if exists

		case "-v", "--mod-version":
			version = flag.Value

		case "-l", "--loader":
			loader = flag.Value

		// case "-s", "--server-side":
		// 	isServerSide = true

		// case "-C", "--client-side":
		// 	isClientSide = true

		case "-t", "--project-type":
			projectType = flag.Value
		case "--match-any":
			isMatchAny = true
		}
	}

	facets = "["
	if categories != nil {
		categoryFacet := "["

		for i, category := range categories {
			categoryFacet += fmt.Sprintf("\"categories:%v\"", category)
			if i == len(categories) - 1 {
				continue
			}
			if isMatchAny {
				categoryFacet += ","
			} else {
				categoryFacet += "],["
			}
		}
		categoryFacet += "]"
		facets += categoryFacet
	}
	if version != "" {
		facets += fmt.Sprintf(",[\"versions:%v\"]", version)
	}

	if projectType != "" {
		facets += fmt.Sprintf(",[\"project_type:%v\"]", projectType)
	}

	if loader != "" {
		facets += fmt.Sprintf(",[\"categories:%v\"]", loader)
	}
	facets += "]"

	if facets != "[]" {
		return url.PathEscape(facets)
	}
	return ""
}

func PrintSearchResults() {
	if len(os.Args) < 3 || os.Args[2] == "--help" {
		SearchHelp()
		return
	}
	flags, keyword := GetFlags(os.Args)
	facets := FlagsToFacets(flags)

	parameters := []QueryParameter{
		{"query", keyword},
	}

	if facets != "" {
		parameters = append(parameters, QueryParameter{"facets", facets})
	}

	queryUrl := QueryBuilder(MainApi + "search", parameters)

	data, err := GetApiData(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryUrl)
	fmt.Println(string(data))

}

func QueryBuilder(endpoint string, parameters []QueryParameter) string {
	endpoint += "?"
	for i, parameter := range parameters {
		if i != 0 {
			endpoint += "&"
		}
		endpoint += parameter.Key + "=" + parameter.Value
	}
	return endpoint
}

func GetApiData(url string) ([]byte, error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	// request.Header.Set("User-Agent", fmt.Sprintf("AmiSimon/modrinth-cli/%v (simon.leneveu@gmail.com)", Version))

	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	if response.StatusCode == 410 {
		log.Fatal("API Returned 410 Gone, it is deprecated, switch to newer")
	}



	body, err := io.ReadAll(response.Body)
	return body, err
}

func main() {
	ParseArguments()
}
