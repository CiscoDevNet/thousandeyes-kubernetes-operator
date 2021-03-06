package controllers

import (
	"github.com/CiscoDevNet/thousandeyes-kubernetes-operator/api/v1alpha1"
	"github.com/william20111/go-thousandeyes"
)

func PageLoad(pageLoad v1alpha1.PageLoad) thousandeyes.PageLoad {
	data := thousandeyes.PageLoad{}
	data.URL = pageLoad.URL
	data.Interval = pageLoad.Interval
	data.HTTPInterval = pageLoad.HTTPInterval
	return data
}

func ComparePageLoad(spec v1alpha1.PageLoad, te thousandeyes.PageLoad) bool {
	if spec.URL != te.URL ||
		spec.Interval != te.Interval ||
		spec.HTTPInterval != te.HTTPInterval {
		return false
	}
	return CompareAgents(spec.Agents, te.Agents) &&
		CompareAlertRules(spec.AlertRules, te.AlertRules)
}
