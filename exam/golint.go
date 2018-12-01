package exam

import (
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
	"golang.org/x/lint"
)

var (
	// LintMinConfidence sets the min confidence for linting
	LintMinConfidence float64
)

// GoLint runs golint on provided go files and adds suggestions to the response
func GoLint(folder *astrav.Folder, r *extypes.Response, pkgName string) bool {
	files := folder.GetRawFiles()
	resLint := lintCode(files)
	if resLint == "" {
		fmt.Println(aurora.Gray("golint:\t\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("golint:\t\t"), aurora.Red("FAIL"))
	fmt.Println(resLint)
	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovement(gtpl.NotLinted)
	} else {
		r.AppendTodo(gtpl.NotLinted)
	}

	return false
}

func lintCode(files map[string][]byte) string {
	l := lint.Linter{}
	ps, err := l.LintFiles(files)
	if err != nil {
		log.Fatal(err)
	}

	var lintRes string
	for _, p := range ps {
		if p.Confidence < LintMinConfidence {
			continue
		}
		lintRes += fmt.Sprintf("%s: %s\n\t%s\n\tdoc: %s\n", p.Category, p.Text, p.Position.String(), p.Link)
	}
	return lintRes
}
