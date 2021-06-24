package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func Agents(specAgents []v1alpha1.Agent, teAgents thousandeyes.Agents) thousandeyes.Agents {
	agents := thousandeyes.Agents{}
	for _, specAgent := range specAgents {
		for _, agent := range teAgents {
			if specAgent.AgentName == agent.AgentName {
				agents = append(agents, thousandeyes.Agent{AgentID: agent.AgentID})
				break
			}
		}
	}
	return agents

}

func AlertRules(specRules []v1alpha1.AlertRule, teRules thousandeyes.AlertRules) thousandeyes.AlertRules {
	rules := thousandeyes.AlertRules{}
	for _, specRule := range specRules {
		for _, rule := range teRules {
			if specRule.RuleName == rule.RuleName {
				rules = append(rules, rule)
				break
			}
		}
	}
	return rules
}

func CompareAgents(specAgents []v1alpha1.Agent, teAgents thousandeyes.Agents) bool {
	if len(specAgents) != len(teAgents) {
		return false
	}
	for _, specAgent := range specAgents {
		flg := false
		for _, teAgent := range teAgents {
			if specAgent.AgentName == teAgent.AgentName {
				flg = true
				break
			}
		}
		if !flg {
			return false
		}
	}
	return true
}

func CompareAlertRules(specRules []v1alpha1.AlertRule, teRules thousandeyes.AlertRules) bool {
	if len(specRules) != 0 && (len(specRules) != len(teRules)) {
		return false
	}
	for _, specRule := range specRules {
		flg := false
		for _, teRule := range teRules {
			if specRule.RuleName == teRule.RuleName {
				flg = true
				break
			}
		}
		if !flg {
			return false
		}
	}
	return true
}
