package calculate

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	blaze_query "github.com/purkhusid/bazel_vcs_differ/protos"
	"github.com/purkhusid/bazel_vcs_differ/query"
)

type HashedRuleTarget struct {
	Rule *blaze_query.Rule
	Hash string
}

// Calculator ...
type Calculator struct {
	QueryResult  *blaze_query.QueryResult
	LabelTargets map[string]*blaze_query.Target
	LabelCache   map[string]string
}

// NewCalculator ...
func NewCalculator(queryResult *blaze_query.QueryResult) Calculator {
	labelTargets := make(map[string]*blaze_query.Target)

	for _, target := range queryResult.GetTarget() {
		targetLabel := query.GetTargetLabel(target)
		labelTargets[targetLabel] = target
	}

	return Calculator{
		QueryResult:  queryResult,
		LabelTargets: labelTargets,
		LabelCache:   make(map[string]string),
	}
}

func (c *Calculator) calculateSourceFileHash(sourceFile *blaze_query.SourceFile) string {
	pathToTarget := filepath.Dir(strings.Split(sourceFile.GetLocation(), ":")[0])
	pathFromTargetSplit := strings.Split(strings.Replace(sourceFile.GetName(), "//", "", -1), ":")
	pathFromTarget := pathFromTargetSplit[len(pathFromTargetSplit)-1]
	combinedPath := path.Join(pathToTarget, pathFromTarget)

	// Get absolute path of file
	fileAbsPath, err := filepath.Abs(combinedPath)
	if err != nil {
		log.Panicf("Failed to get absolute path of %s", combinedPath)
	}

	// Open the file for use
	file, err := os.Open(fileAbsPath)
	if err != nil {
		log.Printf("Failed to open %s", fileAbsPath)
		c.LabelCache[sourceFile.GetName()] = ""
		return ""
	}
	defer file.Close()

	checksum := sha256.New()
	if _, err := io.Copy(checksum, file); err != nil {
		log.Panicf("Failed to get checksum of %s", fileAbsPath)
	}

	hashString := fmt.Sprintf("%x", checksum.Sum(nil))
	c.LabelCache[sourceFile.GetName()] = hashString
	return hashString
}

func (c *Calculator) calculateInputHash(label string) string {
	cachedHash, exists := c.LabelCache[label]

	if exists {
		return cachedHash
	}

	target, exists := c.LabelTargets[label]

	if !exists {
		log.Printf("Did not find target with label %s", label)
		return ""
	}

	if target.GetRule() != nil {
		return c.calculateRuleHash(target.GetRule())
	}

	if target.GetSourceFile() != nil {
		return c.calculateSourceFileHash(target.GetSourceFile())
	}

	return ""
}

func (c *Calculator) calculateRuleHash(rule *blaze_query.Rule) string {
	hash := sha256.New()

	// Add the attributes to the hash
	attributes := query.GetRuleAttributes(rule)
	for _, attribute := range attributes {
		attributeName := attribute.GetName()

		// Exclude attributes that include absolute paths
		if attributeName != "generator_location" && attributeName != "path" && attributeName != "build_file" {
			attributeBytes, err := proto.Marshal(attribute)

			if err != nil {
				log.Panicf("Failed to marshal attribute %s on target %s", attribute.GetName(), rule.GetName())
			}

			hash.Write(attributeBytes)
		}
	}

	for _, input := range rule.GetRuleInput() {
		inputHash := c.calculateInputHash(input)
		hash.Write([]byte(inputHash))
	}

	hashString := fmt.Sprintf("%x", hash.Sum(nil))
	c.LabelCache[rule.GetName()] = hashString
	return hashString
}

// CalculateHashes ...
func (c *Calculator) CalculateHashes() []HashedRuleTarget {
	hashedRuleTargets := make([]HashedRuleTarget, 0)
	ruleTargets := query.GetRuleTargets(c.QueryResult)
	for _, ruleTarget := range ruleTargets {
		ruleHash := c.calculateRuleHash(ruleTarget)
		hashedRuleTarget := HashedRuleTarget{
			Rule: ruleTarget,
			Hash: ruleHash,
		}
		hashedRuleTargets = append(hashedRuleTargets, hashedRuleTarget)
	}

	return hashedRuleTargets
}
