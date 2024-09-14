package compute

import types "libds/gcp/types"

type NodeTemplate struct {
	/*
	   CPU overcommit.
	   Default value is `NONE`.
	   Possible values are: `ENABLED`, `NONE`.
	*/
	CpuOvercommitType string `json:"cpuOvercommitType,omitempty" yaml:"cpuOvercommitType,omitempty"`

	/*
	   Labels to use for node affinity, which will be used in
	   instance scheduling.
	*/
	NodeAffinityLabels map[string]string `json:"nodeAffinityLabels,omitempty" yaml:"nodeAffinityLabels,omitempty"`

	/*
	   Node type to use for nodes group that are created from this template.
	   Only one of nodeTypeFlexibility and nodeType can be specified.
	*/
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	/*
	   Flexible properties for the desired node type. Node groups that
	   use this node template will create nodes of a type that matches
	   these properties. Only one of nodeTypeFlexibility and nodeType can
	   be specified.
	   Structure is documented below.
	*/
	NodeTypeFlexibility types.Compute_NodeTemplateNodeTypeFlexibility `json:"nodeTypeFlexibility,omitempty" yaml:"nodeTypeFlexibility,omitempty"`

	/*
	   Region where nodes using the node template will be created.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// An optional textual description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The server binding policy for nodes using this template. Determines
	   where the nodes should restart following a maintenance event.
	   Structure is documented below.
	*/
	ServerBinding types.Compute_NodeTemplateServerBinding `json:"serverBinding,omitempty" yaml:"serverBinding,omitempty"`
}
