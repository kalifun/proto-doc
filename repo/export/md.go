package export

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/kalifun/proto-doc/consts/doc"
	"github.com/kalifun/proto-doc/entity/api"
)

// markdown
type markdown struct {
	template string // 配置新模板
	outPath  string
	language string
}

type option func(*markdown)

// SetTemplate 设置模板
//
//	@param template
//	@return option
func SetTemplate(template string) option {
	return func(m *markdown) {
		m.template = template
	}
}

// SetOutPath 设置输出路径
//
//	@param outPath
//	@return option
func SetOutPath(outPath string) option {
	return func(m *markdown) {
		m.outPath = outPath
	}
}

// SetLanguage
//
//	@param lan
//	@return option
func SetLanguage(lan string) option {
	return func(m *markdown) {
		m.language = lan
	}
}

// NewMarkdown
//
//	@param opts
//	@return *markdown
func NewMarkdown(opts ...option) *markdown {
	m := markdown{}
	for _, f := range opts {
		f(&m)
	}
	return &m
}

// GenMarkdown
//
//	@receiver m
//	@param apis
//	@return error
func (m *markdown) GenMarkdown(apis []api.ExportApis) error {
	m.setTemplate()
	for _, api := range apis {
		for _, apiInfo := range api.Apis {
			fileName := fmt.Sprintf("%s_%s.md", apiInfo.Name, m.language)
			f, err := os.Create(path.Join(m.getOutPath(), fileName))
			if err != nil {
				return err
			}
			tl, err := template.New("doc").Parse(m.template)
			if err != nil {
				return err
			}
			w := bufio.NewWriter(f)
			err = tl.Execute(w, apiInfo)
			if err != nil {
				return err
			}
			w.Flush()
		}
	}
	return nil
}

// setTemplate
//
//	@receiver m
func (m *markdown) setTemplate() {
	if m.template == "" {
		// 判断用什么模板返回
		switch m.language {
		case "zh":
			m.template = doc.MdZh
		case "en":
			m.template = doc.MdEn
		default:
			m.template = doc.MdEn
		}
	}
}

// getOutPath
//
//	@receiver m
//	@return string
func (m *markdown) getOutPath() string {
	if m.outPath == "" {
		return "./"
	}
	return m.outPath
}
