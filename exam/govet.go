package exam

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/astrav"
	"github.com/tehsphinx/exalysis/extypes"
	"github.com/tehsphinx/exalysis/gtpl"
)

// GoVet runs govet on provided go files and adds suggestions to the response
func GoVet(_ *astrav.Folder, r *extypes.Response, pkgName string) bool {
	res, state := goVet()

	if state.Success() {
		fmt.Println(aurora.Gray("go vet:\t"), aurora.Green("OK"))
		return true
	}

	fmt.Println(aurora.Gray("go vet:\t"), aurora.Red("FAIL"))
	fmt.Println(res)

	if pkgName == "twofer" || pkgName == "hamming" {
		r.AppendImprovement(gtpl.NotLinted)
	} else {
		r.AppendTodo(gtpl.NotLinted)
	}

	return false
}

func goVet() (string, *os.ProcessState) {
	cmd := exec.Command("go", "vet")

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error running go vet: ", err)
	}

	return string(b), cmd.ProcessState
}
