package types

type Biglake_DatabaseHiveOptions struct {
	// Cloud Storage folder URI where the database data is stored, starting with "gs://".
	LocationUri string `json:"locationUri,omitempty" yaml:"locationUri,omitempty"`

	/*
	   Stores user supplied Hive database parameters. An object containing a
	   list of"key": value pairs.
	   Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   - - -
	*/
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
