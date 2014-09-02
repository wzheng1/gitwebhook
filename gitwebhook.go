package gitwebhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type GitWebHook struct{}

// Install webhook REST services with prefix inside mux server.
func InstallREST(prefix string, mux *http.ServeMux) {
	mux.Handle(prefix, http.StripPrefix(prefix, &GitWebHook{}))
}

// Main HTTP service method.
func (*GitWebHook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if ok := verifyRequest(w, req); !ok {
		return
	}
	var err error
	method := req.Header.Get("X-GitHub-Event")
	switch method {
	case "ping":
		err = decodePing(req)
	case "push":
		err = decodePush(req)
	default:
		err = errors.New("Unknown method " + method + "!")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func verifyRequest(w http.ResponseWriter, req *http.Request) bool {
	if req.Method != "POST" {
		http.Error(w, "Wrong method!", http.StatusBadRequest)
		return false
	}
	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Wrong Content-Type!", http.StatusBadRequest)
		return false
	}
	if !strings.HasPrefix(req.Header.Get("User-Agent"), "GitHub-Hookshot/") {
		http.Error(w, "Wrong User-Agent!", http.StatusBadRequest)
		return false
	}

	if req.Header.Get("X-GitHub-Event") == "" {
		http.Error(w, "Missing X-GitHub-Event!", http.StatusBadRequest)
		return false
	}

	return true
}

// decode ping event
func decodePing(req *http.Request) (err error) {
	ping := &Ping{}
	// var ping map[string]interface{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &ping)
	if err != nil {
		return
	}
	fmt.Println("Received PING event:\n", ping, "\n")
	return
}

// decode puh event
func decodePush(req *http.Request) (err error) {
	push := &Push{}
	// var ping map[string]interface{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &push)
	if err != nil {
		return
	}
	fmt.Println("Received PUSH event:\n", push, "\n")
	return
}
