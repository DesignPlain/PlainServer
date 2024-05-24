package ssmincidents

import types "DesignSphere_Server/src/resource/aws/types"

type ReplicationSet struct {
	//
	Regions []types.Ssmincidents_ReplicationSetRegion `json:"regions,omitempty" yaml:"regions,omitempty"`

	/*
	   Tags applied to the replication set.

	   For information about the maximum allowed number of Regions and tag value constraints, see [CreateReplicationSet in the -AWS Systems Manager Incident Manager API Reference-](https://docs.aws.amazon.com/incident-manager/latest/APIReference/API_CreateReplicationSet.html).
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
