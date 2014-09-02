GitWebHook
----------

This repository contains sample Git Webhook implementation in Go language.
Sample usage is as follows:

```go
mux := http.NewServeMux()
gitwebhook.InstallREST("/gitwebhook/", mux)
server := &http.Server{Addr: ":8080", Handler: mux}
fmt.Println("Starting serving @", server.Addr, "...")
go server.ListenAndServe()
select {}
```

Ping call to the webhook might look like this:

```bash
curl -A "GitHub-Hookshot/github" -H "Content-Type:application/json" \
-H "X-Github-Event:ping" -d @resources/pingevent.json http://localhost:8080/gitwebhook/
```

Push call to the webhook might look like this:
```bash
curl -A "GitHub-Hookshot/github" -H "Content-Type:application/json" \
-H "X-Github-Event:push" -d @resources/pushevent.json http://localhost:8080/gitwebhook/
```
