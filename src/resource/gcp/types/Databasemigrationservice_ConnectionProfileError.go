package types

type Databasemigrationservice_ConnectionProfileError struct {
	/*
	   (Output)
	   The status code, which should be an enum value of google.rpc.Code.
	*/
	Code int `json:"code,omitempty" yaml:"code,omitempty"`

	/*
	   (Output)
	   A list of messages that carry the error details.
	*/
	Details []map[string]string `json:"details,omitempty" yaml:"details,omitempty"`

	/*
	   (Output)
	   Human readable message indicating details about the current status.
	*/
	Message string `json:"message,omitempty" yaml:"message,omitempty"`
}
