package calculate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type LabelOutput struct {
	Hash string `json:hash`
}

func WriteResultsToFile(hashedRuleTargets []HashedRuleTarget, path string) {
	f, err := os.Create(path)
	if err != nil {
		f.Close()
		log.Panic(err)
		return
	}

	output := make(map[string]LabelOutput, 0)
	for _, ruleTarget := range hashedRuleTargets {
		output[ruleTarget.Rule.GetName()] = LabelOutput{
			Hash: ruleTarget.Hash,
		}
	}

	bytes, err := json.MarshalIndent(output, "", "    ")

	if err != nil {
		log.Panic(err)
	}

	_, err = f.Write(bytes)
	if err != nil {
		log.Panic(err)
	}

	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("file written successfully")
}
