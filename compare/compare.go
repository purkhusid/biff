package compare

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/purkhusid/bazel_vcs_differ/calculate"
)

func readFileToJSON(path string) map[string]calculate.LabelOutput {
	file, err := os.Open(path)

	if err != nil {
		log.Panic(err)
	}

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		log.Panic(err)
	}

	var output map[string]calculate.LabelOutput
	err = json.Unmarshal(bytes, &output)

	if err != nil {
		log.Panic(err)
	}

	return output
}

func compare(left map[string]calculate.LabelOutput, right map[string]calculate.LabelOutput) map[string]calculate.LabelOutput {
	changedLabels := make(map[string]calculate.LabelOutput)
	for k, v := range right {
		leftValue, exists := left[k]

		if exists {
			if leftValue.Hash != v.Hash {
				changedLabels[k] = v
			}
		} else {
			changedLabels[k] = v
		}

	}

	return changedLabels
}

//CompareAndWriteToOutput ...
func CompareAndWriteToOutput(leftPath string, rightPath string, outputPath string) {
	leftOutput := readFileToJSON(leftPath)
	rightOutput := readFileToJSON(rightPath)

	changedLabels := compare(leftOutput, rightOutput)

	f, err := os.Create(outputPath)
	if err != nil {
		f.Close()
		log.Panic(err)
		return
	}

	for k := range changedLabels {
		fmt.Fprintln(f, k)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("file written successfully")
}
