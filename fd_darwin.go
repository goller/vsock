package vsock

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// contextID retrieves the local context ID for this system.
func contextID() (uint32, error) {
	return 0, nil

	/* Really, we want to use the socket file descriptor to get the context ID.
	cid, err := unix.IoctlGetInt(fd, unix.IOCTL_VM_SOCKETS_GET_LOCAL_CID)
	if err != nil {
		return 0, err
	}
	return uint32(cid), nil
	*/
}

// isErrno determines if an error a matches UNIX error number.
func isErrno(err error, errno int) bool {
	switch errno {
	case ebadf:
		return err == unix.EBADF
	case enotconn:
		return err == unix.ENOTCONN
	default:
		panicf("vsock: isErrno called with unhandled error number parameter: %d", errno)
		return false
	}
}

func panicf(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}
