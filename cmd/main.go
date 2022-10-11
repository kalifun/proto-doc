package main

import (
	"errors"
	"strings"
	"sync"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/kalifun/proto-doc/repo/export"
	"github.com/kalifun/proto-doc/repo/proto"
)

func main() {

}

func reflectProto(path string) error {
	// from google api linter
	var errorsWithPos []protoparse.ErrorWithPos
	var lock sync.Mutex
	p := protoparse.Parser{
		IncludeSourceCodeInfo: true,
		ErrorReporter: func(errorWithPos protoparse.ErrorWithPos) error {
			// Protoparse isn't concurrent right now but just to be safe for the future.
			lock.Lock()
			errorsWithPos = append(errorsWithPos, errorWithPos)
			lock.Unlock()
			// Continue parsing. The error returned will be protoparse.ErrInvalidSource.
			return nil
		},
	}

	fd, err := p.ParseFiles(path)
	if err != nil {
		if err == protoparse.ErrInvalidSource {
			if len(errorsWithPos) == 0 {
				return errors.New("got protoparse.ErrInvalidSource but no ErrorWithPos errors")
			}
			// TODO: There's multiple ways to deal with this but this prints all the errors at least
			errStrings := make([]string, len(errorsWithPos))
			for i, errorWithPos := range errorsWithPos {
				errStrings[i] = errorWithPos.Error()
			}
			return errors.New(strings.Join(errStrings, "\n"))
		}
	}
	apis := proto.ReflectProtos(fd...)
	export.GenMarkdown(apis)
	return nil
}
