# go-lang--http-request-cancellation


To demo, ability to cancel in-progress http request in Go using Goroutine and Context.

Running code:

1) go run main.go
2) Goto public folder > app.js > add your server url at <go_server_url> open index.html and click 
    "Get Movies button" while go server is processing click  "Cancel get movies".
3) Notice that in server, goroutines stops processing.

Note: to cancel in-process http request from javascript I've used AbortController (https://developer.mozilla.org/en-US/docs/Web/API/AbortController) 