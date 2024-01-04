package gcp

import (
	baseModel "DesignSphere_Server/src/resource"
	"fmt"
)

type Model_GCP_VPC_Properties struct {
	AutoCreateSubNetworks       bool   `yaml:"autoCreateSubnetworks,omitempty"`
	MTU                         int    `yaml:"mtu,omitempty"`
	RoutingMode                 string `yaml:"routingMode,omitempty"`
	DeleteDefaultRoutesOnCreate bool   `yaml:"deleteDefaultRoutesOnCreate,omitempty"`
}

func CreateVPCModel(projectName string, networkName string, mtu int, routingMode string) baseModel.ResourceModel {
	// TODO: make this model creation logic completely dynamic, will have to refer pulumi documentation
	resourceModel := baseModel.ResourceModel{
		Name:        projectName,
		Runtime:     "yaml",
		Description: "GCP cloud VPC network pulumi config",
		Outputs: map[string]string{
			"networkIpv4": fmt.Sprintf("${%s.gatewayIpv4}", networkName),
		},
		Resources: map[string]baseModel.Yaml_Resource{
			networkName: {
				Type: "gcp:compute:Network",
				Properties: Model_GCP_VPC_Properties{
					AutoCreateSubNetworks:       true, //autoCreateSubnetworks,
					MTU:                         mtu,
					RoutingMode:                 routingMode,
					DeleteDefaultRoutesOnCreate: false,
				},
			},
		},
	}

	return resourceModel
}
