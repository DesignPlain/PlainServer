package gcp

import (
	baseModel "DesignSphere_Server/src/resource"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

type Model_GCP_Bucket_Properties struct {
	Bucket                   string            `yaml:"bucket,omitempty"`
	Source                   map[string]string `yaml:"source,omitempty"`
	Role                     string            `yaml:"role,omitempty"`
	Members                  []string          `yaml:"members,omitempty"`
	Location                 string            `yaml:"location,omitempty"`
	Website                  map[string]string `yaml:"website,omitempty"`
	UniformBucketLevelAccess bool              `yaml:"uniformBucketLevelAccess,omitempty"`
}

type ResourceModel struct {
	Name        string                             `yaml:"name"`
	Runtime     string                             `yaml:"runtime"`
	Description string                             `yaml:"description"`
	Outputs     map[string]string                  `yaml:"outputs"`
	Resources   map[string]baseModel.Yaml_Resource `yaml:"resources"`
}

func CreateBucketModel(stackName string, bucketName string, members []string, role string, location string) string {
	c := ResourceModel{
		Name:        stackName,
		Runtime:     "yaml",
		Description: "GCP cloud storage bucket pulumi config",
		Outputs: map[string]string{
			"bucketURL": fmt.Sprintf("${%s.url}", bucketName),
		},
		Resources: map[string]baseModel.Yaml_Resource{
			bucketName: {
				Type: "gcp:storage:Bucket",
				Properties: Model_GCP_Bucket_Properties{
					Location: location,
					/* Website: map[string]string{
						"mainPageSuffix": "index.html",
					}, */
					UniformBucketLevelAccess: true,
				},
			},
			/* bucketName + "index-html": {
				Type: "gcp:storage:BucketObject",
				Properties: Model_GCP_Bucket_Properties{
					Bucket: fmt.Sprintf("${%s}", bucketName),
					Source: map[string]string{
						"fn::fileAsset": "./index.html",
					},
				},
			}, */
			bucketName + "IAMbinding": {
				Type: "gcp:storage:BucketIAMBinding",
				Properties: Model_GCP_Bucket_Properties{
					Bucket:  fmt.Sprintf("${%s.name}", bucketName),
					Role:    "roles/storage.objectViewer",
					Members: members,
				},
			},
		},
	}

	out, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	return string(out)
}
