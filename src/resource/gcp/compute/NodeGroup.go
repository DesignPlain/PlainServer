package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type NodeGroup struct {
	// An optional textual description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The initial number of nodes in the node group. One of `initial_size` or `autoscaling_policy` must be configured on resource creation.
	InitialSize int `json:"initialSize,omitempty" yaml:"initialSize,omitempty"`

	/*
	   Specifies the frequency of planned maintenance events. Set to one of the following:
	   - AS_NEEDED: Hosts are eligible to receive infrastructure and hypervisor updates as they become available.
	   - RECURRENT: Hosts receive planned infrastructure and hypervisor updates on a periodic basis, but not more frequently than every 28 days. This minimizes the number of planned maintenance operations on individual hosts and reduces the frequency of disruptions, both live migrations and terminations, on individual VMs.
	   Possible values are: `AS_NEEDED`, `RECURRENT`.
	*/
	MaintenanceInterval string `json:"maintenanceInterval,omitempty" yaml:"maintenanceInterval,omitempty"`

	// Specifies how to handle instances when a node in the group undergoes maintenance. Set to one of: DEFAULT, RESTART_IN_PLACE, or MIGRATE_WITHIN_NODE_GROUP. The default value is DEFAULT.
	MaintenancePolicy string `json:"maintenancePolicy,omitempty" yaml:"maintenancePolicy,omitempty"`

	/*
	   Share settings for the node group.
	   Structure is documented below.
	*/
	ShareSettings types.Compute_NodeGroupShareSettings `json:"shareSettings,omitempty" yaml:"shareSettings,omitempty"`

	// Zone where this node group is located
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   If you use sole-tenant nodes for your workloads, you can use the node
	   group autoscaler to automatically manage the sizes of your node groups.
	   One of `initial_size` or `autoscaling_policy` must be configured on resource creation.
	   Structure is documented below.
	*/
	AutoscalingPolicy types.Compute_NodeGroupAutoscalingPolicy `json:"autoscalingPolicy,omitempty" yaml:"autoscalingPolicy,omitempty"`

	// Name of the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The URL of the node template to which this node group belongs.


	   - - -
	*/
	NodeTemplate string `json:"nodeTemplate,omitempty" yaml:"nodeTemplate,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   contains properties for the timeframe of maintenance
	   Structure is documented below.
	*/
	MaintenanceWindow types.Compute_NodeGroupMaintenanceWindow `json:"maintenanceWindow,omitempty" yaml:"maintenanceWindow,omitempty"`
}
