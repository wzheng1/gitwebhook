Sample Git Webhook in Go
------------------

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
