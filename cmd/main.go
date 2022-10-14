package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/kalifun/proto-doc/repo/export"
	"github.com/kalifun/proto-doc/repo/proto"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "proto-doc",
	Short: "proto-doc is a tool for proto files",
	Long:  `proto-doc is a tool for generating documentation and annotations from proto files`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}

func main() {
	rootCmd.AddCommand(cmdDoc)
	rootCmd.Execute()
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
