package must

import (
	"log/slog"
)

// Must literally just forces zero errors,  yeah.
func Must(err error) {
	if err != nil {
		slog.Error(err.Error())
	}
}
