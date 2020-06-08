package query

import (
	blaze_query "github.com/purkhusid/bazel_vcs_differ/protos"
)

func GetRuleTargets(queryResult *blaze_query.QueryResult) []*blaze_query.Rule {
	rules := make([]*blaze_query.Rule, 0)

	for _, target := range queryResult.GetTarget() {
		if target.GetRule() != nil {
			rules = append(rules, target.GetRule())
		}
	}

	return rules
}

func GetRuleAttributes(rule *blaze_query.Rule) []*blaze_query.Attribute {
	attributes := make([]*blaze_query.Attribute, 0)
	for _, attribute := range rule.GetAttribute() {
		name := attribute.GetName()
		if name != "generator_location" && name != "path" && name != "build_file" {
			attributes = append(attributes, attribute)
		}
	}

	return attributes
}

func GetTargetLabel(target *blaze_query.Target) string {
	if target.GetRule() != nil {
		return target.GetRule().GetName()
	}

	if target.GetSourceFile() != nil {
		return target.GetSourceFile().GetName()
	}

	if target.GetGeneratedFile() != nil {
		return target.GetGeneratedFile().GetName()
	}

	// TODO: Should I include PackageGroup ?
	// if target.GetPackageGroup() != nil {
	// 	return target.GetPackageGroup().GetName()
	// }

	// TODO: Should I include EnvironmentGroup ?
	// if target.GetEnvironmentGroup() != nil {
	// 	return target.GetEnvironmentGroup().GetName()
	// }

	return ""
}
