package stringer

import (
	"io"
)

func Write(w io.Writer, args ...interface{}) error {
	for _, arg := range args {
		if s, ok := arg.(string); ok {
			_, err := w.Write([]byte(s))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
