# my-website
Go, Fiber, and Autocert to serve my four freebie 3rd level domains.

Clone, add a certs/ directory, change the domains whitelisted to your domain or domains.
Figure out where you want static files to live.
/var/www/secure it probably not idiomatic for Go

Then:

go mod init autocert

go mod tidy

go build autocert.go

I run ./autocert > log

https://domain/cert-status will return the expiry of the current cert in json format.

(COMPLETE) Todo #1 Fix the single static files folder.

(COMPLETE) Todo #2 Find a better place to serve the static folders

(COMPLETE) Todo #3 Run as a non-root user - note running as a non-priviledged user required fixing the executable
every time I recompiled it with this command: "setcap cap_net_bind_service+ep ~multi-https/src/multi-https"

I chose multi-https as the user and executable name.

