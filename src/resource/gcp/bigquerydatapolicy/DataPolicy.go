package bigquerydatapolicy

import types "DesignSphere_Server/src/resource/gcp/types"

type DataPolicy struct {
	/*
	   The data masking policy that specifies the data masking rule to use.
	   Structure is documented below.
	*/
	DataMaskingPolicy types.Bigquerydatapolicy_DataPolicyDataMaskingPolicy `json:"dataMaskingPolicy,omitempty" yaml:"dataMaskingPolicy,omitempty"`

	// User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {dataPolicyId} in part of the resource name.
	DataPolicyId string `json:"dataPolicyId,omitempty" yaml:"dataPolicyId,omitempty"`

	/*
	   The enrollment level of the service.
	   Possible values are: `COLUMN_LEVEL_SECURITY_POLICY`, `DATA_MASKING_POLICY`.


	   - - -
	*/
	DataPolicyType string `json:"dataPolicyType,omitempty" yaml:"dataPolicyType,omitempty"`

	// The name of the location of the data policy.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}.
	PolicyTag string `json:"policyTag,omitempty" yaml:"policyTag,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
