package baseModel

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
