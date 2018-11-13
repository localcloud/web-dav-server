package main

import (
	"golang.org/x/net/webdav"
	"net/http"
	"log"
	"fmt"
	"github.com/localcloud/web-dav-server.git/storage"
	"strings"
)

func init() {
	log.SetPrefix("[web-dav-server] ")
}

type MiddleWareFunc func(http.ResponseWriter, *http.Request) bool
type HandlerFunc func(http.ResponseWriter, *http.Request)

func wrapTargetByMiddleWares(targetFn HandlerFunc, handlers ...MiddleWareFunc) HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var (
			runTarget bool
		)
		runTarget = true
		for _, fn := range handlers {
			if runTarget = fn(writer, request); runTarget == false {
				break
			}
		}
		if runTarget == true {
			targetFn(writer, request)
		}
	}
}

func main() {
	var (
		fileSystem webdav.FileSystem
		err        error
		handler    webdav.Handler
	)
	if fileSystem, err = storage.New(&storage.Config{MountPath: "/tmp/dav_server"}); err != nil {
		log.Fatalln(err)
	}
	handler = webdav.Handler{
		LockSystem: webdav.NewMemLS(),
		FileSystem: fileSystem,
		Logger: func(r *http.Request, e error) {
			msg := fmt.Sprintf("%s  %s  %s", r.UserAgent(), r.Method, strings.ToLower(r.RequestURI))
			if err != nil {
				msg += fmt.Sprintf(" error: %s", err)
			}
			log.Println(msg)
		},
	}
	http.HandleFunc("/", wrapTargetByMiddleWares(handler.ServeHTTP, authMiddleware))
	log.Fatalln(http.ListenAndServe(":5566", nil))
}

func authMiddleware(writer http.ResponseWriter, request *http.Request) bool {
	if request.Header.Get("Authorization") == "" {
		writer.Header().Set("WWW-Authenticate", "Basic realm=\"LocalCloud WebDav\"")
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write(nil)
		return false
	}
	return true
}
