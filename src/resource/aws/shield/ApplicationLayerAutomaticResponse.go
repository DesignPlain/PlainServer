package shield

import types "DesignSphere_Server/src/resource/aws/types"

type ApplicationLayerAutomaticResponse struct {
	// ARN of the resource to protect (Cloudfront Distributions and ALBs only at this time).
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	//
	Timeouts types.Shield_ApplicationLayerAutomaticResponseTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// One of `COUNT` or `BLOCK`
	Action string `json:"action,omitempty" yaml:"action,omitempty"`
}
