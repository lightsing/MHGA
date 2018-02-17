package proxy

import (
	"github.com/elazarl/goproxy"
	"bytes"
	"io/ioutil"
	"net/http"
)

var NewHTTPProxy = goproxy.NewProxyHttpServer()

type Response struct { }

func (*Response) Redirect(req *http.Request, uri string) *http.Response {
	resp := &http.Response{}
	resp.Request = req
	resp.TransferEncoding = req.TransferEncoding
	resp.Header = make(http.Header)
	resp.Header.Add("Location", uri)
	resp.Header.Add("Content-Type", goproxy.ContentTypeHtml)
	resp.StatusCode = http.StatusMovedPermanently

	buf := bytes.NewBufferString(`<html>
<head><title>301 Moved Permanently</title></head>
<body bgcolor="white">
<center><h1>301 Moved Permanently</h1></center>
<hr><center>MHGA</center>
</body>
</html>
`)
	resp.ContentLength = int64(buf.Len())
	resp.Body = ioutil.NopCloser(buf)
	return resp
}
