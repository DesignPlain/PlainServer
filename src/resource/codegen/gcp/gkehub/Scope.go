package gkehub

type Scope struct {
	/*
	   Labels for this Scope.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Scope-level cluster namespace labels. For the member clusters bound
	   to the Scope, these labels are applied to each namespace under the
	   Scope. Scope-level labels take precedence over Namespace-level
	   labels (`namespace_labels` in the Fleet Namespace resource) if they
	   share a key. Keys and values must be Kubernetes-conformant.
	*/
	NamespaceLabels map[string]string `json:"namespaceLabels,omitempty" yaml:"namespaceLabels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The client-provided identifier of the scope.


	   - - -
	*/
	ScopeId string `json:"scopeId,omitempty" yaml:"scopeId,omitempty"`
}
