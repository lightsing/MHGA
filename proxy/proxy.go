package proxy

import (
	"bytes"
	"github.com/elazarl/goproxy"
	"html/template"
	"io/ioutil"
	"net/http"
)

type ResponseBuilder struct {
	version  string
	template *template.Template
}

func (rsb *ResponseBuilder) mustRender(code int) *bytes.Buffer {
	var buf bytes.Buffer
	if err := rsb.template.Execute(&buf, template.FuncMap{
		"StatusCode": code,
		"StatusText": http.StatusText(code),
		"Version":    rsb.version,
	}); err != nil {
		panic(err)
	}
	return &buf
}

func (rsb *ResponseBuilder) Gen(req *http.Request, uri string, code int) *http.Response {
	resp := &http.Response{}
	resp.Request = req
	resp.TransferEncoding = req.TransferEncoding
	resp.Header = make(http.Header)
	resp.Header.Add("Location", uri)
	resp.Header.Add("Server", rsb.version)
	resp.Header.Add("Content-Type", goproxy.ContentTypeHtml)
	resp.StatusCode = code
	buf := rsb.mustRender(code)
	resp.ContentLength = int64(buf.Len())
	resp.Body = ioutil.NopCloser(buf)
	return resp
}

func NewResponseBuilder(version string) *ResponseBuilder {
	rsb := &ResponseBuilder{
		version: "MHGA - " + version,
		template: template.Must(template.New("MHGA").Parse(`<html>
<head><title>{{.StatusCode}} {{.StatusText}}</title></head>
<body bgcolor="white">
<center><h1>{{.StatusCode}} {{.StatusText}}</h1></center>
<hr><center>{{.Version}}</center>
</body>
</html>
`)),
	}
	return rsb
}
