package gkehub

import types "libds/gcp/types"

type Membership struct {
	/*
	   If this Membership is a Kubernetes API server hosted on GKE, this is a self link to its GCP resource.
	   Structure is documented below.
	*/
	Endpoint types.Gkehub_MembershipEndpoint `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	/*
	   Labels to apply to this membership.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Location of the membership.
	   The default value is `global`.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The client-provided identifier of the membership.


	   - - -
	*/
	MembershipId string `json:"membershipId,omitempty" yaml:"membershipId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Authority encodes how Google will recognize identities from this Membership.
	   See the workload identity documentation for more details:
	   https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	   Structure is documented below.
	*/
	Authority types.Gkehub_MembershipAuthority `json:"authority,omitempty" yaml:"authority,omitempty"`

	/*
	   The name of this entity type to be displayed on the console. This field is unavailable in v1 of the API.

	   > --Warning:-- `description` is deprecated and will be removed in a future major release.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
