package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"
)

var protoPath string
var outPath string
var language string

func init() {
	// proto 文件路径
	cmdDoc.PersistentFlags().StringVar(&protoPath, "proto", ".", "proto files (default is .)")
	// doc 导出路径
	cmdDoc.PersistentFlags().StringVar(&outPath, "out", ".", "export document path (default is .)")
	// doc 导出采用的模板语言
	cmdDoc.PersistentFlags().StringVar(&language, "language", "en", "export document template language (default is en)")
}

var cmdDoc = &cobra.Command{
	Use:   "doc",
	Short: "generating documentation",
	Long:  `Parsing proto to generate markdown documents`,
	// Args:  cobra.MinimumNArgs(1),
	Run: genDoc,
}

func genDoc(cmd *cobra.Command, args []string) {
	pi := protoImp{
		inputPath:  "",
		outputPath: outPath,
		language:   language,
	}

	for _, file := range getFiles(protoPath) {
		pi.inputPath = file
		err := pi.reflectProto()
		if err != nil {
			panic(err)
		}
	}
}

func getFiles(protoPath string) []string {
	paths := []string{}
	if protoPath == "." {
		currentPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		files := walkProto(protoPath)
		for _, file := range files {
			paths = append(paths, path.Join(currentPath, file))
		}
	} else {
		paths = append(paths, protoPath)
	}
	return paths
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
