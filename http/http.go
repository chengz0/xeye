package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"

	"encoding/json"
	"log"
)

var m *martini.ClassicMartini

type Dto struct {
	Succ bool
	Msg  string
	Data interface{}
}

func ErrDto(message string) Dto {
	return Dto{Succ: false, Msg: message}
}

func DataDto(d interface{}) Dto {
	return Dto{Succ: true, Msg: "", Data: d}
}

func RenderErrDto(message string) string {
	dto := ErrDto(message)
	bs, err := json.Marshal(dto)
	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

func RenderDataDto(d interface{}) string {
	dto := DataDto(d)
	bs, err := json.Marshal(dto)

	if err != nil {
		return err.Error()
	} else {
		return string(bs)
	}
}

func StartHttpServer() {
	// new martini
	m = martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Funcs: []template.FuncMap{{
			"nl2br":      nl2br,
			"htmlquote":  htmlQuote,
			"str2html":   str2html,
			"dateformat": dateFormat,
		}},
	}))

	m.Get("/what", func(r render.Render) {
		log.Println("-------")
		r.HTML(200, "hello", "xeye")
	})

	m.Get("/", func(re render.Render) {
		m := make(map[string]string)
		m["version"] = "1.0.1"
		re.HTML(200, "index", m)
	})

	// m.Run()
	KernelRouter()

	SystemRouter()

	LoginRouter()

	CpuRouter()

	MemRouter()

	SdaRouter()

	DiskRouter()

	NetRouter()

	IORouter()

	http.ListenAndServe(":3000", m)
}
