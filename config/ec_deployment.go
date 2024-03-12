package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func ConfigureECDeployment(p *config.Provider) {
	p.AddResourceConfigurator("ec_deployment", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "elastic"
		r.UseAsync = true
		r.Kind = "ElasticCloudDeployment"
		//r.ExternalName.OmittedFields = []string{
		//	"apm",
		//	"elasticsearch",
		//	"enterprise_search",
		//	"integrations_server",
		//	"kibana",
		//	"observability",
		//}
	})
}
