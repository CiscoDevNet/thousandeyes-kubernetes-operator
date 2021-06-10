package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"reflect"
	devnetv1alpha1 "wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func CreateWebTransaction(client *thousandeyes.Client, spec devnetv1alpha1.WebTransactionTestSpec) error {
	transaction := thousandeyes.WebTransaction{}
	for _, agent := range spec.WebTransaction.Agents {
		transaction.Agents = append(transaction.Agents, thousandeyes.Agent{AgentID: agent.AgentID})
	}
	transaction.TestName = spec.WebTransaction.TestName
	transaction.URL = spec.WebTransaction.URL
	transaction.Interval = spec.WebTransaction.Interval
	transaction.TransactionScript = spec.WebTransaction.TransactionScript
	_, err := client.CreateWebTransaction(transaction)
	return err
}

func DeleteWebTransaction(client *thousandeyes.Client, spec devnetv1alpha1.WebTransactionTestSpec) error {
	return client.DeleteWebTransaction(spec.WebTransaction.TestID)
}

func UpdateWebTransaction(client *thousandeyes.Client, spec devnetv1alpha1.WebTransactionTestSpec) error {
	transaction := thousandeyes.WebTransaction{}
	for _, agent := range spec.WebTransaction.Agents {
		transaction.Agents = append(transaction.Agents, thousandeyes.Agent{AgentID: agent.AgentID})
	}
	transaction.TestName = spec.WebTransaction.TestName
	transaction.URL = spec.WebTransaction.URL
	transaction.Interval = spec.WebTransaction.Interval
	transaction.TransactionScript = spec.WebTransaction.TransactionScript
	_, err := client.UpdateWebTransaction(spec.WebTransaction.TestID, transaction)
	return err
}

func EqualWebTransaction(metadata devnetv1alpha1.WebTransaction, transaction thousandeyes.WebTransaction) bool {
	if metadata.URL == transaction.URL &&
		metadata.Interval == transaction.Interval &&
		metadata.TransactionScript == transaction.TransactionScript &&
		reflect.DeepEqual(metadata.Agents, transaction.Agents) {
		return true
	}
	return false
}
