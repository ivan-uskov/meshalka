package config

import (
	"net/http"
)

type UseArgs struct {
	Use bool
	Con map[string]string
}

type Page struct {
	Id   string
	Path string
	Args UseArgs
}

func (p *Page) Copy() Page {
	newP := Page{
		Id:   p.Id,
		Path: p.Path,
		Args: UseArgs{
			Use: p.Args.Use,
			Con: make(map[string]string),
		},
	}
	for k, v := range p.Args.Con {
		newP.Args.Con[k] = v
	}
	return newP
}

type Handler func(http.ResponseWriter, *http.Request, Page)

type Route struct {
	P Page
	H Handler
}

type PageInfo struct {
	Styles  []string
	Scrypts []string
	Title   string
	TplName string
	Content interface{}
}

func NewPageInfo(pageId string, title string, tplName string, content interface{}) PageInfo {
	return PageInfo{
		Title:   title,
		TplName: tplName,
		Styles:  GetPageStyles(pageId),
		Scrypts: GetPageScripts(pageId),
		Content: content,
	}
}

type FormInfo struct {
	Styles  []string
	Scrypts []string
	TplName string
	Content interface{}
}

func NewFormInfo(pageId string, tplName string, content interface{}) FormInfo {
	return FormInfo{
		TplName: tplName,
		Styles:  GetPageStyles(pageId),
		Scrypts: GetPageScripts(pageId),
		Content: content,
	}
}
