package gitwebhook

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const prefix = "/testing"

func TestWrongMethod(t *testing.T) {
	server := setup()
	defer server.Close()

	resp, _ := http.Get(server.URL + prefix)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusBadRequest ||
		!strings.Contains(string(body), "method") {
		t.Errorf("Expected BadRequest , got %s: %s!", resp.Status, string(body))
	}
}

func TestWrongContentType(t *testing.T) {
	server := setup()
	defer server.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL+prefix, nil)
	req.Header.Add("Content-Type", "application/text")
	req.Header.Add("User-Agent", "GitHub-Hookshot/github")
	req.Header.Add("X-Github-Event", "ping")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusBadRequest ||
		!strings.Contains(string(body), "Content-Type") {
		t.Errorf("Excepcted BadRequest, got %s: %s!", resp.Status, string(body))
	}
}

func TestWrongUserAgent(t *testing.T) {
	server := setup()
	defer server.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL+prefix, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "go-lang")
	req.Header.Add("X-Github-Event", "ping")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusBadRequest ||
		!strings.Contains(string(body), "User-Agent") {
		t.Errorf("Excepcted BadRequest, got %s: %s!", resp.Status, string(body))
	}
}

func TestMissingGithubEvent(t *testing.T) {
	server := setup()
	defer server.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL+prefix, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "GitHub-Hookshot/github")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusBadRequest ||
		!strings.Contains(string(body), "X-GitHub-Event") {
		t.Errorf("Excepcted BadRequest, got %s: %s!", resp.Status, string(body))
	}
}

func TestWrongGithubEvent(t *testing.T) {
	server := setup()
	defer server.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("POST", server.URL+prefix, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "GitHub-Hookshot/github")
	req.Header.Add("X-GitHub-Event", "wrong")
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusBadRequest ||
		!strings.Contains(string(body), "Unknown") {
		t.Errorf("Excepcted BadRequest, got %s: %s!", resp.Status, string(body))
	}
}

func TestJsonPingEvent(t *testing.T) {
	server := setup()
	defer server.Close()

	postFile("ping", "pingevent.json", server.URL+prefix,
		http.StatusOK, t)
}

func TestJsonPushEvent(t *testing.T) {
	server := setup()
	defer server.Close()

	postFile("push", "pushevent.json", server.URL+prefix,
		http.StatusOK, t)
}

func setup() (server *httptest.Server) {
	mux := http.NewServeMux()
	InstallREST(prefix, mux)
	server = httptest.NewServer(mux)
	return
}

func postFile(event string, filename string, url string, expStatusCode int, t *testing.T) {
	data, err := ioutil.ReadFile("resources/" + filename)
	if err != nil {
		t.Errorf("Failed to open %s: %s", filename, err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		t.Errorf("Error creating POST request: %s!", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "GitHub-Hookshot/github")
	req.Header.Add("X-Github-Event", event)
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("Failed posting webhook to: %s!", url)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != expStatusCode {
		t.Errorf("Wrong response code, expecting %d, got %s: %s!",
			expStatusCode, resp.Status, string(body))
	}
}
