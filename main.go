package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var language string

const ReqUrl = "https://raw.githubusercontent.com/github/gitignore/refs/heads/main/"

func init() {

	flag.StringVar(&language, "l", "", "Language to get gitignore for")
	flag.Parse()

	if language == "" {
		fmt.Println("Specify a language :3")
		os.Exit(1)
	}
	language = strings.ToLower(language)
	language = strings.ToUpper(language[:1]) + language[1:]

}

func main() {
	url := fmt.Sprintf("%s%s.gitignore", ReqUrl, language)

	r, err := http.Get(url)
	if err != nil {
		fmt.Println("error while sending request - invalid language: ", err)
		os.Exit(2)
	}

	if r.StatusCode != 200 {
		fmt.Println("invalid language")
		os.Exit(3)
	}
	defer r.Body.Close()

	content, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("couldn't read request body: ", err)
		os.Exit(4)
	}

	file, err := os.Create(".gitignore")
	if err != nil {
		fmt.Println("couldn't create gitignore: ", err)
		os.Exit(5)
	}
	defer file.Close()
	file.Write(content)

	fmt.Printf("Gitignore for %s created!\n", language)

}
