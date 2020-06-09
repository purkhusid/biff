package calculate

import (
	"bytes"
	"context"
	"log"
	"os"
	"os/exec"

	"github.com/golang/protobuf/proto"
	blaze_query "github.com/purkhusid/bazel_vcs_differ/protos"
)

func GetQueryResult(workspacePath, bazel, query string) *blaze_query.QueryResult {
	ctx, _ := context.WithCancel(context.Background())
	args := append([]string(nil), "query", "--output=proto", "--order_output=full", query)
	stdoutBuffer := new(bytes.Buffer)
	stdErrBuffer := new(bytes.Buffer)
	command := exec.CommandContext(ctx, bazel, args...)
	command.Dir = workspacePath
	command.Stdout = stdoutBuffer
	command.Stderr = stdErrBuffer
	err := command.Run()

	if err != nil {
		log.Println("Failed to run query")
		log.Println("Stdout:")
		log.Println(string(stdoutBuffer.Bytes()))
		log.Println("Stderr:")
		log.Println(string(stdErrBuffer.Bytes()))
		os.Exit(1)
	}

	queryResult := &blaze_query.QueryResult{}
	err = proto.Unmarshal(stdoutBuffer.Bytes(), queryResult)

	if err != nil {
		log.Panicf("Failed to unmarshal query output: %s", err)
	}

	return queryResult
}
