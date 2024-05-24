package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type FieldLevelEncryptionConfig struct {
	// An optional comment about the Field Level Encryption Config.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Content Type Profile Config specifies when to forward content if a content type isn't recognized and profiles to use as by default in a request if a query argument doesn't specify a profile to use.
	ContentTypeProfileConfig types.Cloudfront_FieldLevelEncryptionConfigContentTypeProfileConfig `json:"contentTypeProfileConfig,omitempty" yaml:"contentTypeProfileConfig,omitempty"`

	// Query Arg Profile Config that specifies when to forward content if a profile isn't found and the profile that can be provided as a query argument in a request.
	QueryArgProfileConfig types.Cloudfront_FieldLevelEncryptionConfigQueryArgProfileConfig `json:"queryArgProfileConfig,omitempty" yaml:"queryArgProfileConfig,omitempty"`
}
