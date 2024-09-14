package gkehub

type Namespace struct {
	// The name of the Scope instance.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	/*
	   Id of the scope


	   - - -
	*/
	ScopeId string `json:"scopeId,omitempty" yaml:"scopeId,omitempty"`

	// The client-provided identifier of the namespace.
	ScopeNamespaceId string `json:"scopeNamespaceId,omitempty" yaml:"scopeNamespaceId,omitempty"`

	/*
	   Labels for this Namespace.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Namespace-level cluster namespace labels. These labels are applied
	   to the related namespace of the member clusters bound to the parent
	   Scope. Scope-level labels (`namespace_labels` in the Fleet Scope
	   resource) take precedence over Namespace-level labels if they share
	   a key. Keys and values must be Kubernetes-conformant.
	*/
	NamespaceLabels map[string]string `json:"namespaceLabels,omitempty" yaml:"namespaceLabels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
