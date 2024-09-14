package ssmincidents

import types "libds/aws/types"

type ReplicationSet struct {
	//
	Regions []types.Ssmincidents_ReplicationSetRegion `json:"regions,omitempty" yaml:"regions,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
