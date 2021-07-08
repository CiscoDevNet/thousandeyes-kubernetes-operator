package controllers

import (
	"github.com/william20111/go-thousandeyes"
	"wwwin-github.cisco.com/DevNet/thousandeyes-operator/api/v1alpha1"
)

func HTTPServer(httpServer v1alpha1.HTTPServer) thousandeyes.HTTPServer {
	data := thousandeyes.HTTPServer{}
	data.URL = httpServer.URL
	data.Interval = httpServer.Interval
	return data
}

func CompareHTTPServer(spec v1alpha1.HTTPServer, te thousandeyes.HTTPServer) bool {
	if spec.URL != te.URL ||
		spec.Interval != te.Interval {
		return false
	}
	return CompareAgents(spec.Agents, te.Agents) &&
		CompareAlertRules(spec.AlertRules, te.AlertRules)
}
