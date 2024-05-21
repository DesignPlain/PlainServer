package workbench

import types "DesignSphere_Server/src/resource/gcp/types"

type Instance struct {
	// Desired state of the Workbench Instance. Set this field to `ACTIVE` to start the Instance, and `STOPPED` to stop the Instance.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`

	/*
	   Part of `parent`. See documentation of `projectsId`.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Optional. If true, the workbench instance will not register with the proxy.
	DisableProxyAccess bool `json:"disableProxyAccess,omitempty" yaml:"disableProxyAccess,omitempty"`

	/*
	   The definition of how to configure a VM instance outside of Resources and Identity.
	   Structure is documented below.
	*/
	GceSetup types.Workbench_InstanceGceSetup `json:"gceSetup,omitempty" yaml:"gceSetup,omitempty"`

	// Required. User-defined unique ID of this instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   'Optional. Input only. The owner of this instance after creation. Format:
	   `alias@example.com` Currently supports one owner only. If not specified, all of
	   the service account users of your VM instance''s service account can use the instance.'
	*/
	InstanceOwners []string `json:"instanceOwners,omitempty" yaml:"instanceOwners,omitempty"`

	/*
	   Optional. Labels to apply to this instance. These can be later modified
	   by the UpdateInstance method.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The name of this workbench instance. Format: `projects/{project_id}/locations/{location}/instances/{instance_id}`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
