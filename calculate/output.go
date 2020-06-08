package calculate

import (
	"fmt"
	"log"
	"os"
)

func WriteResultsToFile(hashedRuleTargets []HashedRuleTarget, path string) {
	f, err := os.Create(path)
	if err != nil {
		f.Close()
		log.Panic(err)
		return
	}

	for _, ruleTarget := range hashedRuleTargets {
		line := fmt.Sprintf("%s %s", ruleTarget.Rule.GetName(), ruleTarget.Hash)
		fmt.Fprintln(f, line)
		if err != nil {
			log.Panic(err)
		}
	}
	err = f.Close()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("file written successfully")
}
