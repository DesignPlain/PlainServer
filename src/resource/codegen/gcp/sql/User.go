package sql

import types "libds/gcp/types"

type User struct {
	/*
	   The user type. It determines the method to authenticate the
	   user during login. The default is the database's built-in user type. Flags
	   include "BUILT_IN", "CLOUD_IAM_USER", "CLOUD_IAM_GROUP" or "CLOUD_IAM_SERVICE_ACCOUNT".
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   The deletion policy for the user.
	   Setting `ABANDON` allows the resource to be abandoned rather than deleted. This is useful
	   for Postgres, where users cannot be deleted from the API if they have been granted SQL roles.

	   Possible values are: `ABANDON`.

	   - - -
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	/*
	   The host the user can connect from. This is only supported
	   for BUILT_IN users in MySQL instances. Don't set this field for PostgreSQL and SQL Server instances.
	   Can be an IP address. Changing this forces a new resource to be created.
	*/
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   The name of the Cloud SQL instance. Changing this
	   forces a new resource to be created.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   The name of the user. Changing this forces a new resource
	   to be created.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The password for the user. Can be updated. For Postgres
	   instances this is a Required field, unless type is set to either CLOUD_IAM_USER
	   or CLOUD_IAM_SERVICE_ACCOUNT. Don't set this field for CLOUD_IAM_USER
	   and CLOUD_IAM_SERVICE_ACCOUNT user types for any Cloud SQL instance.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	//
	PasswordPolicy types.Sql_UserPasswordPolicy `json:"passwordPolicy,omitempty" yaml:"passwordPolicy,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
