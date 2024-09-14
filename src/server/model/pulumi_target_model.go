package model

import (
	"reflect"
	"strings"

	base "libds/ds_base"
)

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

func CreateResourceModel(project_name string, resource_type base.ResourceType, resource_instance_name string, resource_instance interface{}) ResourceModel {
	resourceConfigModel := ResourceModel{
		Name:        project_name,
		Runtime:     "yaml",
		Description: GetString(resource_type),
		Resources: map[string]Yaml_Resource{
			resource_instance_name: {
				Type:       GetURI(resource_type),
				Properties: resource_instance,
			},
		},
		Outputs: getOutputMap(resource_instance_name, resource_instance),
	}

	return resourceConfigModel
}

func getOutputMap(resource_instance_name string, resource interface{}) map[string]string {
	t := reflect.TypeOf(resource)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	outputMap := make(map[string]string, 0)

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		fieldName := getFormattedFieldName(name)
		outputMap[name] = "${" + resource_instance_name + "." + fieldName + "}"
	}

	return outputMap
}

func getFormattedFieldName(name string) string {
	charArray := []rune(name)
	fieldName := strings.ToLower(string(charArray[0])) + string(charArray[1:])
	return fieldName
}
