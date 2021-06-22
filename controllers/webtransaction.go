package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func WebTransaction(webTransaction v1alpha1.WebTransaction) thousandeyes.WebTransaction {
	payload := thousandeyes.WebTransaction{}
	for _, agent := range webTransaction.Agents {
		payload.Agents = append(payload.Agents, thousandeyes.Agent{AgentID: agent.AgentID})
	}
	payload.URL = webTransaction.URL
	payload.Interval = webTransaction.Interval
	payload.TransactionScript = webTransaction.TransactionScript
	return payload
}
