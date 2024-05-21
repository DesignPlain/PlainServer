package monitoring

import types "DesignSphere_Server/src/resource/gcp/types"

type GenericService struct {
	/*
	   Labels which have been used to annotate the service. Label keys must start
	   with a letter. Label keys and values may contain lowercase letters,
	   numbers, underscores, and dashes. Label keys and values have a maximum
	   length of 63 characters, and must be less than 128 bytes in size. Up to 64
	   label entries may be stored. For labels which do not have a semantic value,
	   the empty string may be supplied for the label value.
	*/
	UserLabels map[string]string `json:"userLabels,omitempty" yaml:"userLabels,omitempty"`

	/*
	   A well-known service type, defined by its service type and service labels.
	   Valid values of service types and services labels are described at
	   https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
	   Structure is documented below.
	*/
	BasicService types.Monitoring_GenericServiceBasicService `json:"basicService,omitempty" yaml:"basicService,omitempty"`

	// Name used for UI elements listing this Service.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   An optional service ID to use. If not given, the server will generate a
	   service ID.


	   - - -
	*/
	ServiceId string `json:"serviceId,omitempty" yaml:"serviceId,omitempty"`
}
