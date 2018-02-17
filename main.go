package main

import (
	"github.com/elazarl/goproxy"
	"github.com/lightsing/makehttps/rules"
	"github.com/lightsing/makehttps/proxy"
	"github.com/lightsing/makehttps/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const name = "mhga"
var Version string

func main() {
	// Only log the warning severity or above.
	config.Init(name)
	start := time.Now()
	if ruleSets, err := rules.LoadRuleSets("rules/rules"); err == nil {
		log.Warnf("Load all rule in %s", time.Since(start))
		server := goproxy.NewProxyHttpServer()
		responseBuilder := proxy.NewResponseBuilder(Version)
		server.OnRequest().DoFunc(
			func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
				start := time.Now()
				if result, ok := ruleSets.Apply(req.RequestURI); ok {
					result := *result
					log.Infof("[%s] Redirect %s to %s", time.Since(start), req.RequestURI, result)
					var code int
					if req.Method == "GET" {
						code = http.StatusMovedPermanently
					} else {
						code = http.StatusTemporaryRedirect
					}
					return req, responseBuilder.Gen(req, result, code)
				}
				log.Infof("[%s] Nothing to do with %s", time.Since(start), req.RequestURI)
				return req, nil
			})
		log.Fatal(http.ListenAndServe(":8080", server))
	}

}
