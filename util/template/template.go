package template

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

type (
	//NavMap The nap map for the book
	NavMap struct {
		Title  string `xml:"docTitle>text"`
		Author string `xml:"docAuthor>text"`
		Navs   []nav  `xml:"navMap>navPoint"`
	}

	nav struct {
		Title   string `xml:"navLabel>text"`
		Src     src    `xml:"content"`
		SubNavs []nav  `xml:"navPoint"`
	}

	src struct {
		URL string `xml:"src,attr"`
	}
)

var (
	navMap NavMap
)

//Build 构造页面框架
func Build(workingPath string, tocPath string) error {
	err := parseToc(tocPath)
	if err != nil {
		return err
	}

	// Back original "index.html" to "index.back.html" if it exists.
	indexPath := path.Join(workingPath, "index.html")
	indexBackPath := path.Join(workingPath, "index.back.html")
	os.Rename(indexPath, indexBackPath)

	err = parseTemplate(indexPath)
	if err != nil {
		return err
	}

	return nil
}

func parseTemplate(file string) error {
	t := template.New("template")
	t = t.Funcs(template.FuncMap{"getFirstSrc": getFirstSrc})
	t, err := t.Parse(htmlTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, navMap)
	if err != nil {
		return err
	}
	return nil
}

func getFirstSrc(navs []nav) string {
	if navs[0].Src.URL == "index.html" {
		return "index.bak.html"
	}
	return navs[0].Src.URL
}

//parseToc Parse table of contents
func parseToc(tocPath string) error {
	file, err := os.Open(tocPath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = xml.Unmarshal(data, &navMap)
	if err != nil {
		return err
	}
	return nil
}
