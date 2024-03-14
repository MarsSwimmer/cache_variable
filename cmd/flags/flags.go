package flags

var(
	Port int32

	DefaultValue string // when error return this value

	Size int32 // cache size, unit MB

	ShowErrMsg bool // show error message when internal error
)