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
	buf := rsb.mustRender(code)
	return &http.Response{
		Request: req,
		TransferEncoding: req.TransferEncoding,
		Header: http.Header{
			"Location": []string{uri},
			"Server": []string{rsb.version},
			"Content-Type": []string{goproxy.ContentTypeHtml},
		},
		StatusCode: code,
		ContentLength: int64(buf.Len()),
		Body: ioutil.NopCloser(buf),
	}
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
