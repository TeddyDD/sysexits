// package sysexits is port of sysexits.h for Go
// See sysexits.h(3)
package sysexits

import "os"

type statusCode struct {
	code int
	msg  string
}

func (s statusCode) Error() string {
	return s.msg
}

func (s statusCode) StatusCode() int {
	return s.code
}

var (
	// UsageError means that command was used incorrectly, e.g., with the
	// wrong number of arguments, a bad flag, a bad syntax in a parameter,
	// or whatever.
	UsageError error = statusCode{code: 64, msg: "command was used incorrectly"}
	// DataError means the input data was incorrect in some way.
	// This should only be used for user's data and not system files.
	DataError error = statusCode{code: 65, msg: "input data was incorrect"}

	// NoInputError means an input file (not a system file) did not exist or
	// was not readable.
	NoInputError error = statusCode{code: 66, msg: "input file does not exist or is not readable"}

	// NoUserError means the user specified did not exist. This might be used for
	// mail addresses or remote logins.
	NoUserError error = statusCode{code: 67, msg: "specified user does not exist"}

	// NoHostError means the host specified did not exist. This is used in mail addresses
	// or network requests.
	NoHostError error = statusCode{code: 68, msg: "specified host does not exist"}

	// UnavailableError means a service is unavailable. This can occur if a support program
	// or file does not exist. It can also be used as a catchall message when something you
	// wanted to do doesn't work, but you don't know why.
	UnavailableError error = statusCode{code: 69, msg: "service is unavailable"}

	// SoftwareError means an internal software error has been detected. This should be
	// limited to non-operating system related errors as possible.
	SoftwareError error = statusCode{code: 70, msg: "internal software error"}

	// OSError means an operating system error has been detected. This is intended to be used
	// for things like "cannot fork", "cannot create pipe", or similar errors. It includes
	// cases like getuid returning a user that does not exist in the passwd file.
	OSError error = statusCode{code: 71, msg: "operating system error"}

	// OSFileError means some system file (e.g., /etc/passwd, /var/run/utmp, etc.) does not exist,
	// cannot be opened, or has some sort of error (e.g., syntax error).
	OSFileError error = statusCode{code: 72, msg: "system file error"}

	// CantCreateError means a (user specified) output file cannot be created.
	CantCreateError error = statusCode{code: 73, msg: "output file cannot be created"}

	// IOError means an error occurred while doing I/O on some file.
	IOError error = statusCode{code: 74, msg: "I/O error"}

	// TempFailError means temporary failure, indicating something that is not really an error.
	// In sendmail, this means that a mailer (e.g.) could not create a connection, and the request
	// should be reattempted later.
	TempFailError error = statusCode{code: 75, msg: "temporary failure"}

	// ProtocolError means the remote system returned something that was "not possible" during a
	// protocol exchange.
	ProtocolError error = statusCode{code: 76, msg: "protocol error"}

	// NoPermError means you did not have sufficient permission to perform the operation. This is
	// not intended for file system problems, which should use NoInputError or CantCreateError,
	// but rather for higher-level permissions.
	NoPermError error = statusCode{code: 77, msg: "insufficient permission"}

	// ConfigError means something was found in an unconfigured or misconfigured state.
	ConfigError error = statusCode{code: 78, msg: "unconfigured or misconfigured state"}
)

// StatusCoder is an interface that can be implemented by error to specify
// status code.
type StatusCoder interface {
	// StatusCode returns exit status code for given error.
	// For portability, the status code should be in the range [0, 125].
	StatusCode() int
}

// Exit with error. If error is a [StatusCoder] then the result of [StatusCode]
// method is used. Otherwise 1 is used as status code.
func Exit(err error) {
	if v, ok := err.(StatusCoder); ok {
		os.Exit(v.StatusCode())
	}
	os.Exit(1)
}
