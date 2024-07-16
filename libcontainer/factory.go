package libcontainer

import (
	"github.com/pacoxu/runc/libcontainer/configs"
)

type Factory interface {
	// Creates a new container with the given id and starts the initial process inside it.
	// id must be a string containing only letters, digits and underscores and must contain
	// between 1 and 1024 characters, inclusive.
	//
	// The id must not already be in use by an existing container. Containers created using
	// a factory with the same path (and filesystem) must have distinct ids.
	//
	// Returns the new container with a running process.
	//
	// On error, any partially created container parts are cleaned up (the operation is atomic).
	Create(id string, config *configs.Config) (Container, error)

	// Load takes an ID for an existing container and returns the container information
	// from the state.  This presents a read only view of the container.
	Load(id string) (Container, error)

	// StartInitialization is an internal API to libcontainer used during the reexec of the
	// container.
	StartInitialization() error

	// Type returns info string about factory type (e.g. lxc, libcontainer...)
	Type() string
}
