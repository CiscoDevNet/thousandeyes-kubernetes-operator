package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func WebTransaction(webTransaction v1alpha1.WebTransaction) thousandeyes.WebTransaction {
	data := thousandeyes.WebTransaction{}
	data.URL = webTransaction.URL
	data.Interval = webTransaction.Interval
	data.TransactionScript = webTransaction.TransactionScript
	return data
}

func CompareWebTransaction(spec v1alpha1.WebTransaction, te thousandeyes.WebTransaction) bool {
	if spec.URL != te.URL ||
		spec.Interval != te.Interval ||
		spec.TransactionScript != te.TransactionScript {
		return false
	}
	return CompareAgents(spec.Agents, te.Agents) &&
		CompareAlertRules(spec.AlertRules, te.AlertRules)
}
