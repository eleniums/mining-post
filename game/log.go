package game

import (
	"log/slog"
)

// Create an error attribute for logging.
func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}
