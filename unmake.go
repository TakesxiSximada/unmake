package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const version = 2

func DisplayUsage(makefilePath string) {
	fp, err := os.OpenFile(makefilePath, os.O_RDONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	buf, err := ioutil.ReadAll(fp)
	if err != nil {
		panic("")
	}
	regx, err := regexp.Compile(`(?m:^([^._][a-zA-Z0-9_\-]*):(.*)\n(^\t\@\##.*\n)?(^\t\@\#\s*\n)?((^\t\@\#.*\n)*))`)
	if err != nil {
		panic("")
	}
	for _, match := range regx.FindAllSubmatch(buf, -1) {
		cmdName := strings.TrimSpace(string(match[1]))
		paramLine := strings.TrimSpace(strings.Trim(string(match[3]), "\t@#"))
		usage := strings.Replace(string(match[5]), "\t@#", "     ", -1)
		fmt.Printf("* %s %s\n\n%s\n", cmdName, paramLine, usage)
	}
}
