package rekognition

import types "DesignSphere_Server/src/resource/aws/types"

type Project struct {
	// Specify if automatic retraining should occur. Valid values are `ENABLED` or `DISABLED`. Defaults to `DISABLED`
	AutoUpdate string `json:"autoUpdate,omitempty" yaml:"autoUpdate,omitempty"`

	// Specify the feature being customized. Valid values are `CONTENT_MODERATION` or `CUSTOM_LABELS`. Defaults to `CUSTOM_LABELS`
	Feature string `json:"feature,omitempty" yaml:"feature,omitempty"`

	/*
	   Desired name of the project

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Timeouts types.Rekognition_ProjectTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
