package gcp

import (
	baseModel "DesignSphere_Server/src/resource"
	"fmt"
)

type Model_GCP_ComputeInstance_Properties struct {
	Name                  string           `yaml:"name,omitempty"`
	MachineType           string           `yaml:"machineType,omitempty"`
	Zone                  string           `yaml:"zone,omitempty"`
	MetadataStartupScript string           `yaml:"metadataStartupScript,omitempty"`
	BootDiskImage         map[string]any   `yaml:"bootDisk,omitempty"`
	NetworkInterfaces     []map[string]any `yaml:"networkInterfaces,omitempty"`
	ServiceAccount        map[string]any   `yaml:"serviceAccount,omitempty"`
}

func CreateComputeInstanceModel(projectName string, instanceName string, network string, account string) baseModel.ResourceModel {
	// TODO: make this model creation logic completely dynamic, will have to refer pulumi documentation
	resourceModel := baseModel.ResourceModel{
		Name:        projectName,
		Runtime:     "yaml",
		Description: "GCP cloud VPC network pulumi config",
		Outputs: map[string]string{
			"InstanceId":  "${ComputeInstanceResource.instanceId}",
			"cpuPlatform": "${ComputeInstanceResource.cpuPlatform}",
		},
		Resources: map[string]baseModel.Yaml_Resource{
			"ComputeInstanceResource": {
				Type: "gcp:compute:Instance",
				Properties: Model_GCP_ComputeInstance_Properties{
					Name:        instanceName,
					MachineType: "e2-micro", //autoCreateSubnetworks,
					Zone:        "us-central1-c",
					BootDiskImage: map[string]any{
						"initializeParams": map[string]any{
							"image": "debian-cloud/debian-11",
						},
					},
					NetworkInterfaces: append(make([]map[string]any, 0), map[string]any{
						"network": network, //default
					}),
					ServiceAccount: map[string]any{
						"email":  fmt.Sprintf("${%s.email}", account),
						"scopes": []string{"cloud-platform"},
					},
				},
			},
			account: {
				Type: "gcp:serviceAccount:Account",
				Properties: map[string]any{
					"accountId": "designsphere-managed-" + account,
				},
			},
		},
	}

	return resourceModel
}
