package gcp

import (
	baseModel "DesignSphere_Server/src/resource"
	"fmt"
)

type Model_GCP_Bucket_Properties struct {
	Name                     string            `yaml:"name,omitempty"`
	Bucket                   string            `yaml:"bucket,omitempty"`
	Source                   map[string]string `yaml:"source,omitempty"`
	Role                     string            `yaml:"role,omitempty"`
	Members                  []string          `yaml:"members,omitempty"`
	Location                 string            `yaml:"location,omitempty"`
	Website                  map[string]string `yaml:"website,omitempty"`
	UniformBucketLevelAccess bool              `yaml:"uniformBucketLevelAccess,omitempty"`
}

func CreateBucketModel(projectName string, bucketName string, members []string, role string, location string) baseModel.ResourceModel {
	// TODO: make this model creation logic completely dynamic, will have to refer pulumi documentation
	resourceModel := baseModel.ResourceModel{
		Name:        projectName,
		Runtime:     "yaml",
		Description: "GCP cloud storage bucket pulumi config",
		Outputs: map[string]string{
			"bucketURL": "${StorageBucketResource.url}",
		},
		Resources: map[string]baseModel.Yaml_Resource{
			"StorageBucketResource": {
				Type: "gcp:storage:Bucket",
				Properties: Model_GCP_Bucket_Properties{
					Name:     bucketName,
					Location: location,
					/* Website: map[string]string{
						"mainPageSuffix": "index.html",
					}, */
					UniformBucketLevelAccess: true,
				},
			},
			/* bucketName + "object": {
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

	return resourceModel
}
