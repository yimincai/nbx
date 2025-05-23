package utils

import (
	"os/user"

	"github.com/yimincai/nbx/pkg/clog"
)

// GetUsernameOrPanic gets the current system username
func GetUsernameOrPanic() string {
	u, err := user.Current()
	if err != nil {
		clog.Panic("failed to get current user: %v", err)
	}

	return u.Username
}
