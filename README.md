# my-website
Go, Fiber, and Autocert to serve my four freebie 3rd level domains.

Clone, modify multi-https.go for the domains you want it to serve.

Then:

go mod init multi-https

go mod tidy

go build multi-https.go

I wanted to run it as a user named multi-https, so on Fedora42, I had to do this:
  setcap cap_net_bind_service+ep ~multi-https/src/multi-https


I run ./multi-https > log

https://domain/cert-status will return the expiry of the current cert in json format.

(COMPLETE) Todo #1 Fix the single static files folder.

(COMPLETE) Todo #2 Find a better place to serve the static folders

(COMPLETE) Todo #3 Run as a non-root user - note running as a non-priviledged user required fixing the executable
every time I recompiled it with this command: "setcap cap_net_bind_service+ep ~multi-https/src/multi-https"

I chose multi-https as the user and executable name.

