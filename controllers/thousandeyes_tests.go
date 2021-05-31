package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"reflect"
	devnetv1alpha1 "wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func CreateTest(client *thousandeyes.Client, spec devnetv1alpha1.ThousandEyesTestSpec) (metadata devnetv1alpha1.Metadata, err error) {
	switch spec.TestType {
	case "page-load":
		payload := &thousandeyes.PageLoad{}
		for _, agent := range spec.Metadata.Agents {
			payload.AddAgent(agent.AgentID)
		}
		payload.TestName = spec.Metadata.TestName
		payload.URL = spec.Metadata.URL
		payload.Interval = spec.Metadata.Interval
		payload.HTTPInterval = spec.Metadata.HttpInterval
		result, err := client.CreatePageLoad(*payload)
		if err != nil {
			return devnetv1alpha1.Metadata{}, err
		}
		//save the observed info
		metadata.TestName = result.TestName
		metadata.TestID = result.TestID
		metadata.HttpInterval = result.HTTPInterval
		metadata.Interval = result.Interval
		metadata.URL = result.URL
		return metadata, nil
	}
	return
}

func DeleteTest(client *thousandeyes.Client, spec devnetv1alpha1.ThousandEyesTestSpec) (err error) {
	switch spec.TestType {
	case "page-load":
		id := spec.Metadata.TestID
		return client.DeletePageLoad(id)
	}
	return
}

func UpdateTest(client *thousandeyes.Client, spec devnetv1alpha1.ThousandEyesTestSpec) (testID *int, err error) {
	switch spec.TestType {
	case "page-load":
		payload := &thousandeyes.PageLoad{}
		for _, agent := range spec.Metadata.Agents {
			payload.AddAgent(agent.AgentID)
		}
		payload.TestName = spec.Metadata.TestName
		payload.URL = spec.Metadata.URL
		payload.Interval = spec.Metadata.Interval
		payload.HTTPInterval = spec.Metadata.HttpInterval
		result, err := client.UpdatePageLoad(spec.Metadata.TestID, *payload)
		if err != nil {
			return nil, err
		}
		return &result.TestID, nil
	}
	return
}

func EqualMetadata(metadata devnetv1alpha1.Metadata, pageload thousandeyes.PageLoad) bool {
	if metadata.URL == pageload.URL &&
		metadata.Interval == pageload.Interval &&
		metadata.HttpInterval == pageload.HTTPInterval &&
		reflect.DeepEqual(metadata.Agents, pageload.Agents) {
		return true
	}
	return false
}
