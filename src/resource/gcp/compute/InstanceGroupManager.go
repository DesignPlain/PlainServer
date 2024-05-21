package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type InstanceGroupManager struct {
	/*
	   Properties to set on all instances in the group. After setting
	   allInstancesConfig on the group, you must update the group's instances to
	   apply the configuration.
	*/
	AllInstancesConfig types.Compute_InstanceGroupManagerAllInstancesConfig `json:"allInstancesConfig,omitempty" yaml:"allInstancesConfig,omitempty"`

	/*
	   Whether to wait for all instances to be created/updated before
	   returning. Note that if this is set to true and the operation does not succeed, this provider will
	   continue trying until it times out.
	*/
	WaitForInstances bool `json:"waitForInstances,omitempty" yaml:"waitForInstances,omitempty"`

	/*
	   The zone that instances in this group should be created
	   in.

	   - - -
	*/
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   An optional textual description of the instance
	   group manager.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// External network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
	StatefulExternalIps []types.Compute_InstanceGroupManagerStatefulExternalIp `json:"statefulExternalIps,omitempty" yaml:"statefulExternalIps,omitempty"`

	/*
	   The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)

	   - - -
	*/
	UpdatePolicy types.Compute_InstanceGroupManagerUpdatePolicy `json:"updatePolicy,omitempty" yaml:"updatePolicy,omitempty"`

	/*
	   Application versions managed by this instance group. Each
	   version deals with a specific instance template, allowing canary release scenarios.
	   Structure is documented below.
	*/
	Versions []types.Compute_InstanceGroupManagerVersion `json:"versions,omitempty" yaml:"versions,omitempty"`

	/*
	   When used with `wait_for_instances` it specifies the status to wait for.
	   When `STABLE` is specified this resource will wait until the instances are stable before returning. When `UPDATED` is
	   set, it will wait for the version target to be reached and any per instance configs to be effective as well as all
	   instances to be stable before returning. The possible values are `STABLE` and `UPDATED`
	*/
	WaitForInstancesStatus string `json:"waitForInstancesStatus,omitempty" yaml:"waitForInstancesStatus,omitempty"`

	/*
	   The base instance name to use for
	   instances in this group. The value must be a valid
	   [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	   are lowercase letters, numbers, and hyphens (-). Instances are named by
	   appending a hyphen and a random four-character string to the base instance
	   name.
	*/
	BaseInstanceName string `json:"baseInstanceName,omitempty" yaml:"baseInstanceName,omitempty"`

	/*
	   The named port configuration. See the section below
	   for details on configuration.
	*/
	NamedPorts []types.Compute_InstanceGroupManagerNamedPort `json:"namedPorts,omitempty" yaml:"namedPorts,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Internal network IPs assigned to the instances that will be preserved on instance delete, update, etc. This map is keyed with the network interface name. Structure is documented below.
	StatefulInternalIps []types.Compute_InstanceGroupManagerStatefulInternalIp `json:"statefulInternalIps,omitempty" yaml:"statefulInternalIps,omitempty"`

	/*
	   The target number of running instances for this managed instance group. This value should always be explicitly set
	   unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
	*/
	TargetSize int `json:"targetSize,omitempty" yaml:"targetSize,omitempty"`

	/*
	   The autohealing policies for this managed instance
	   group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	*/
	AutoHealingPolicies types.Compute_InstanceGroupManagerAutoHealingPolicies `json:"autoHealingPolicies,omitempty" yaml:"autoHealingPolicies,omitempty"`

	// The instance lifecycle policy for this managed instance group.
	InstanceLifecyclePolicy types.Compute_InstanceGroupManagerInstanceLifecyclePolicy `json:"instanceLifecyclePolicy,omitempty" yaml:"instanceLifecyclePolicy,omitempty"`

	/*
	   Pagination behavior of the `listManagedInstances` API
	   method for this managed instance group. Valid values are: `PAGELESS`, `PAGINATED`.
	   If `PAGELESS` (default), Pagination is disabled for the group's `listManagedInstances` API method.
	   `maxResults` and `pageToken` query parameters are ignored and all instances are returned in a single
	   response. If `PAGINATED`, pagination is enabled, `maxResults` and `pageToken` query parameters are
	   respected.
	*/
	ListManagedInstancesResults string `json:"listManagedInstancesResults,omitempty" yaml:"listManagedInstancesResults,omitempty"`

	/*
	   The name of the instance group manager. Must be 1-63
	   characters long and comply with
	   [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
	   include lowercase letters, numbers, and hyphens.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []types.Compute_InstanceGroupManagerStatefulDisk `json:"statefulDisks,omitempty" yaml:"statefulDisks,omitempty"`

	/*
	   The full URL of all target pools to which new
	   instances in the group are added. Updating the target pools attribute does
	   not affect existing instances.
	*/
	TargetPools []string `json:"targetPools,omitempty" yaml:"targetPools,omitempty"`
}
