package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type NodeTemplate struct {
	// Name of the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

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
	   The server binding policy for nodes using this template. Determines
	   where the nodes should restart following a maintenance event.
	   Structure is documented below.
	*/
	ServerBinding types.Compute_NodeTemplateServerBinding `json:"serverBinding,omitempty" yaml:"serverBinding,omitempty"`

	/*
	   CPU overcommit.
	   Default value is `NONE`.
	   Possible values are: `ENABLED`, `NONE`.
	*/
	CpuOvercommitType string `json:"cpuOvercommitType,omitempty" yaml:"cpuOvercommitType,omitempty"`

	// An optional textual description of the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Labels to use for node affinity, which will be used in
	   instance scheduling.
	*/
	NodeAffinityLabels map[string]string `json:"nodeAffinityLabels,omitempty" yaml:"nodeAffinityLabels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Region where nodes using the node template will be created.
	   If it is not provided, the provider region is used.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
