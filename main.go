package main

import (
	"github.com/lightsing/makehttps/rules"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/elazarl/goproxy"
	"bytes"
	"io/ioutil"
	"net/http"
)

func genResp(req *http.Request, uri *string) *http.Response {
	resp := &http.Response{}
	resp.Request = req
	resp.TransferEncoding = req.TransferEncoding
	resp.Header = make(http.Header)
	resp.Header.Add("Location", *uri)
	resp.Header.Add("Content-Type", goproxy.ContentTypeHtml)
	resp.StatusCode = http.StatusTemporaryRedirect

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

func main() {
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	start := time.Now()
	if ruleSets, err := rules.LoadRuleSets("rules/rules"); err == nil {
		log.Warnf("Load all rule in %s", time.Since(start))
		proxy := goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		proxy.OnRequest().DoFunc(
			func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
				start := time.Now()
				if result, ok := ruleSets.Apply(req.RequestURI); ok {
					log.Infof("[%s] Redirect %s to %s", time.Since(start), req.RequestURI, result)
					return req, genResp(req, result)
				}
				log.Infof("[%s] Nothing to do with %s", time.Since(start), req.RequestURI)
				return req, nil
			})
		log.Fatal(http.ListenAndServe(":8080", proxy))

	}

}
