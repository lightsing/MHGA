package main

import (
	"fmt"
	"github.com/lightsing/makehttps/rules"
)

func main() {
	/*proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	pattern := regexp.MustCompile(`(?i)^http:`)
	proxy.OnRequest(goproxy.DstHostIs("store.steampowered.com")).DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			resp := &http.Response{}
			resp.Request = req
			resp.TransferEncoding = req.TransferEncoding
			resp.Header = make(http.Header)
			resp.Header.Add("Location", pattern.ReplaceAllString(req.RequestURI, "https:"))
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
			return req, resp
		})
	log.Fatal(http.ListenAndServe(":8080", proxy))*/
	ruleSet, err := rules.LoadRuleSet("rules/rules/rules/Google.xml")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%v\n", ruleSet)
		for _, target := range ruleSet.Targets {
			fmt.Println(target.Is("http://www.google.com.hk/test"))
		}
	}
}