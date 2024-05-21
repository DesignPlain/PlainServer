package resourcemanager

type Lien struct {
	/*
	   A stable, user-visible/meaningful string identifying the origin
	   of the Lien, intended to be inspected programmatically. Maximum length of
	   200 characters.
	*/
	Origin string `json:"origin,omitempty" yaml:"origin,omitempty"`

	/*
	   A reference to the resource this Lien is attached to.
	   The server will validate the parent against those for which Liens are supported.
	   Since a variety of objects can have Liens against them, you must provide the type
	   prefix (e.g. "projects/my-project-name").
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Concise user-visible strings indicating why an action cannot be performed
	   on a resource. Maximum length of 200 characters.
	*/
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`

	/*
	   The types of operations which should be blocked as a result of this Lien.
	   Each value should correspond to an IAM permission. The server will validate
	   the permissions against those for which Liens are supported.  An empty
	   list is meaningless and will be rejected.
	   e.g. ['resourcemanager.projects.delete']


	   - - -
	*/
	Restrictions []string `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`
}
