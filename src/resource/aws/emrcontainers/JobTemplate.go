package emrcontainers

import types "DesignSphere_Server/src/resource/aws/types"

type JobTemplate struct {
	// The KMS key ARN used to encrypt the job template.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// The specified name of the job template.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The job template data which holds values of StartJobRun API request.
	JobTemplateData types.Emrcontainers_JobTemplateJobTemplateData `json:"jobTemplateData,omitempty" yaml:"jobTemplateData,omitempty"`
}
