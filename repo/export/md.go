package export

import (
	"bufio"
	"fmt"
	"os"
	"text/template"

	"github.com/kalifun/proto-doc/entity/api"
)

// GenMd
//
//	@param apis
func GenMarkdown(apis []api.ExportApis) {
	for _, v := range apis {
		for _, val := range v.Apis {
			f, _ := os.Create(fmt.Sprintf("%s.md", val.Name))
			defer f.Close()
			tl, err := template.ParseFiles("../api.template")
			if err != nil {
				fmt.Println(err)
				continue
			}
			w := bufio.NewWriter(f)
			tl.Execute(w, val)
			w.Flush()
		}
	}
}
