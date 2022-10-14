package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var protoPath string
var outPath string

func init() {
	// proto 文件路径
	cmdDoc.PersistentFlags().StringVar(&protoPath, "proto", ".", "proto files (default is .)")
	// doc 导出路径
	cmdDoc.PersistentFlags().StringVar(&protoPath, "out", ".", "export document path (default is .)")
}

var cmdDoc = &cobra.Command{
	Use:   "doc",
	Short: "generating documentation",
	Long:  `Parsing proto to generate markdown documents`,
	// Args:  cobra.MinimumNArgs(1),
	Run: genDoc,
}

func genDoc(cmd *cobra.Command, args []string) {
	if protoPath == "." {
		currentPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		files := walkProto(protoPath)
		for _, file := range files {
			err := reflectProto(path.Join(currentPath, file))
			if err != nil {
				panic(err)
			}
		}
	}
	if protoPath != "" {
		err := reflectProto(protoPath)
		if err != nil {
			panic(err)
		}
	} else {
		panic("未找到proto文件")
	}

}

// walkProto
//
//	@param root
//	@return []string
func walkProto(root string) []string {
	var files []string
	err := filepath.Walk(root, func(filePath string, info os.FileInfo, err error) error {
		filesuffix := path.Ext(filePath)
		if filesuffix == ".proto" {
			files = append(files, filePath)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
