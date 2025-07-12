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

Todo #1 Fix the single static files folder.
Todo #2 Find a better place to serve the static folders
Todo #3 Run as a non-root user
