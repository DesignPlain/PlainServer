package firestore

type Document struct {
	/*
	   The client-assigned document ID to use for this document during creation.


	   - - -
	*/
	DocumentId string `json:"documentId,omitempty" yaml:"documentId,omitempty"`

	// The document's [fields](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.documents) formated as a json string.
	Fields string `json:"fields,omitempty" yaml:"fields,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The collection ID, relative to database. For example: chatrooms or chatrooms/my-document/private-messages.
	Collection string `json:"collection,omitempty" yaml:"collection,omitempty"`

	// The Firestore database id. Defaults to `"(default)"`.
	Database string `json:"database,omitempty" yaml:"database,omitempty"`
}
