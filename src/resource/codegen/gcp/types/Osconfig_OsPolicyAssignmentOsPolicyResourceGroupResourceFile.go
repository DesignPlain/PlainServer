package types

type Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceFile struct {
	/*
	   A a file with this content. The size of the content
	   is limited to 1024 characters.
	*/
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	/*
	   A remote or local source. Structure is
	   documented below.
	*/
	File Osconfig_OsPolicyAssignmentOsPolicyResourceGroupResourceFileFile `json:"file,omitempty" yaml:"file,omitempty"`

	// The absolute path of the file within the VM.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Consists of three octal digits which represent, in
	   order, the permissions of the owner, group, and other users for the file
	   (similarly to the numeric mode used in the linux chmod utility). Each digit
	   represents a three bit number with the 4 bit corresponding to the read
	   permissions, the 2 bit corresponds to the write bit, and the one bit
	   corresponds to the execute permission. Default behavior is 755. Below are
	   some examples of permissions and their associated values: read, write, and
	   execute: 7 read and execute: 5 read and write: 6 read only: 4
	*/
	Permissions string `json:"permissions,omitempty" yaml:"permissions,omitempty"`

	/*
	   Desired state of the file. Possible values are:
	   `DESIRED_STATE_UNSPECIFIED`, `PRESENT`, `ABSENT`, `CONTENTS_MATCH`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
