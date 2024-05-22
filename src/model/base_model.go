package model

type Yaml_Resource struct {
	Type       string `yaml:"type"`
	Properties any    `yaml:"properties"`
}

type ResourceModel struct {
	Name        string                   `yaml:"name"`
	Runtime     string                   `yaml:"runtime"`
	Description string                   `yaml:"description"`
	Outputs     map[string]string        `yaml:"outputs"`
	Resources   map[string]Yaml_Resource `yaml:"resources"`
}

func CreateResourceModel(project_name string, description string, resource_type ResourceType, resource_instance_name string, resource_instance interface{}) ResourceModel {
	resourceConfigModel := ResourceModel{
		Name:        project_name,
		Runtime:     "yaml",
		Description: description,
		// Outputs: map[string]string{
		// 	"InstanceId":  "${ComputeInstanceResource" + instanceName + ".instanceId}",
		// 	"cpuPlatform": "${ComputeInstanceResource" + instanceName + ".cpuPlatform}",
		// },
		Resources: map[string]Yaml_Resource{
			"resource_type.String()" + "_" + resource_instance_name: {
				Type:       "gcp:compute:Instance",
				Properties: resource_instance,
			},
		},
	}

	return resourceConfigModel
}
