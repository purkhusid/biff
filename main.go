package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/golang/protobuf/proto"
	blaze_query "github.com/purkhusid/bazel_vcs_differ/protos"
)

type HashedRuleTarget struct {
	Rule *blaze_query.Rule
	Hash string
}

func getQueryResult(workspacePath, bazel string) *blaze_query.QueryResult {
	ctx, _ := context.WithCancel(context.Background())
	args := append([]string(nil), "query", "--output=proto", "--order_output=full", "//external:all-targets + deps(//...:all-targets)")
	stdoutBuffer := new(bytes.Buffer)
	stdErrBuffer := new(bytes.Buffer)
	command := exec.CommandContext(ctx, bazel, args...)
	command.Dir = workspacePath
	command.Stdout = stdoutBuffer
	command.Stderr = stdErrBuffer
	err := command.Run()

	if err != nil {
		log.Panicf("Failed to run query: %s", string(stdErrBuffer.Bytes()))
	}

	queryResult := &blaze_query.QueryResult{}
	err = proto.Unmarshal(stdoutBuffer.Bytes(), queryResult)

	if err != nil {
		log.Panicf("Failed to unmarshal query output: %s", err)
	}

	return queryResult
}

func writeResultsToFile(hashedRuleTargets []HashedRuleTarget, path string) {
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

func main() {
	workspacePath := os.Args[1]
	bazel := os.Args[2]
	outputPath := os.Args[3]

	queryResult := getQueryResult(workspacePath, bazel)
	calculator := NewCalculator(queryResult)

	hashedTargets := calculator.CalculateHashes()

	writeResultsToFile(hashedTargets, path.Join(workspacePath, outputPath))
}
