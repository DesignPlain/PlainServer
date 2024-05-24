package chime

import types "DesignSphere_Server/src/resource/aws/types"

type SdkvoiceVoiceProfileDomain struct {
	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of Voice Profile Domain.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of Voice Profile Domain.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration for server side encryption.
	ServerSideEncryptionConfiguration types.Chime_SdkvoiceVoiceProfileDomainServerSideEncryptionConfiguration `json:"serverSideEncryptionConfiguration,omitempty" yaml:"serverSideEncryptionConfiguration,omitempty"`
}
