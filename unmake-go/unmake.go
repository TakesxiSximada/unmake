package unmake

import (
	"strings"
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"path/filepath"
	"io/ioutil"
)


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


func ParseInclude(makefilePath string) {
	absMakefilePath, err := filepath.Abs(makefilePath)
	if err != nil {
		return
	}
	makefileDirPath := filepath.Dir(absMakefilePath)
	DisplayUsage(makefilePath)
	fp, err := os.OpenFile(absMakefilePath, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer fp.Close()

	regx, err := regexp.Compile(`^include ([^#]+)`)
	if err != nil {
		log.Println(err)
		return
    }
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := scanner.Text()

		for _, matching := range regx.FindAllStringSubmatch(line, -1) {
			if len(matching) > 0 {
				for _, includeFilePath := range matching[1:] {
					if !path.IsAbs(includeFilePath) {
						includeFilePath = path.Join(makefileDirPath, includeFilePath)
						includeFilePath, err = filepath.Abs(includeFilePath)
					}
					ParseInclude(includeFilePath)
				}
			}
		}
	}
}
