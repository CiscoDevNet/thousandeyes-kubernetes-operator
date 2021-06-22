package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func PageLoad(pageLoad v1alpha1.PageLoad) thousandeyes.PageLoad {
	payload := thousandeyes.PageLoad{}
	for _, agent := range pageLoad.Agents {
		payload.AddAgent(agent.AgentID)
	}
	payload.URL = pageLoad.URL
	payload.Interval = pageLoad.Interval
	payload.HTTPInterval = pageLoad.HTTPInterval
	return payload
}
