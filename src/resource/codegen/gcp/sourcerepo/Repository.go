package sourcerepo

import types "libds/gcp/types"

type Repository struct {
	/*
	   Resource name of the repository, of the form `{{repo}}`.
	   The repo name may contain slashes. eg, `name/with/slash`


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   How this repository publishes a change in the repository through Cloud Pub/Sub.
	   Keyed by the topic names.
	   Structure is documented below.
	*/
	PubsubConfigs []types.Sourcerepo_RepositoryPubsubConfig `json:"pubsubConfigs,omitempty" yaml:"pubsubConfigs,omitempty"`
}
