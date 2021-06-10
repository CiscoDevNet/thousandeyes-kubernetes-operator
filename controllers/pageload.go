package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"reflect"
	devnetv1alpha1 "wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func CreatePageLoad(client *thousandeyes.Client, spec devnetv1alpha1.PageLoadTestSpec) error {
	payload := &thousandeyes.PageLoad{}
	for _, agent := range spec.PageLoad.Agents {
		payload.AddAgent(agent.AgentID)
	}
	payload.TestName = spec.PageLoad.TestName
	payload.URL = spec.PageLoad.URL
	payload.Interval = spec.PageLoad.Interval
	payload.HTTPInterval = spec.PageLoad.HttpInterval
	_, err := client.CreatePageLoad(*payload)
	return err
}

func DeletePageLoad(client *thousandeyes.Client, spec devnetv1alpha1.PageLoadTestSpec) error {
	return client.DeletePageLoad(spec.PageLoad.TestID)
}

func UpdatePageLoad(client *thousandeyes.Client, spec devnetv1alpha1.PageLoadTestSpec) error {
	payload := &thousandeyes.PageLoad{}
	for _, agent := range spec.PageLoad.Agents {
		payload.AddAgent(agent.AgentID)
	}
	payload.TestName = spec.PageLoad.TestName
	payload.URL = spec.PageLoad.URL
	payload.Interval = spec.PageLoad.Interval
	payload.HTTPInterval = spec.PageLoad.HttpInterval
	_, err := client.UpdatePageLoad(spec.PageLoad.TestID, *payload)
	return err
}

func EqualPageLoad(metadata devnetv1alpha1.PageLoad, pageload thousandeyes.PageLoad) bool {
	if metadata.URL == pageload.URL &&
		metadata.Interval == pageload.Interval &&
		metadata.HttpInterval == pageload.HTTPInterval &&
		reflect.DeepEqual(metadata.Agents, pageload.Agents) {
		return true
	}
	return false
}
