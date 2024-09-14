package sql

type Database struct {
	/*
	   The deletion policy for the database. Setting ABANDON allows the resource
	   to be abandoned rather than deleted. This is useful for Postgres, where databases cannot be
	   deleted from the API if there are users other than cloudsqlsuperuser with access. Possible
	   values are: "ABANDON", "DELETE". Defaults to "DELETE".
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   The name of the Cloud SQL instance. This does not include the project
	   ID.


	   - - -
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   The name of the database in the Cloud SQL instance.
	   This does not include the project ID or instance name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The charset value. See MySQL's
	   [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	   and Postgres' [Character Set Support](https://www.postgresql.org/docs/9.6/static/multibyte.html)
	   for more details and supported values. Postgres databases only support
	   a value of `UTF8` at creation time.
	*/
	Charset string `json:"charset,omitempty" yaml:"charset,omitempty"`

	/*
	   The collation value. See MySQL's
	   [Supported Character Sets and Collations](https://dev.mysql.com/doc/refman/5.7/en/charset-charsets.html)
	   and Postgres' [Collation Support](https://www.postgresql.org/docs/9.6/static/collation.html)
	   for more details and supported values. Postgres databases only support
	   a value of `en_US.UTF8` at creation time.
	*/
	Collation string `json:"collation,omitempty" yaml:"collation,omitempty"`
}
