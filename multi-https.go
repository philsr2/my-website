package main

import (
	"crypto/tls"
	"log"
	"time"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "${host} ${ip} - - [${time}] \"${method} ${url} ${protocol}\" ${status} ${bytesSent} \"${referer}\" \"${ua}\"\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "America/Chicago",
	}))
//

	app.Get("/cert-status", func(c *fiber.Ctx) error {
		host := c.Hostname()
		conn, err := tls.Dial("tcp", host+":443", &tls.Config{
			InsecureSkipVerify: true, 
		})
		if err != nil {
			log.Println("TLS dial error:", err)
			return c.Status(500).SendString("Failed to connect to TLS endpoint.")
		}
		defer conn.Close()

		state := conn.ConnectionState()
		if len(state.PeerCertificates) == 0 {
			return c.Status(500).SendString("No certificate found in TLS handshake.")
		}

		cert := state.PeerCertificates[0]
		return c.JSON(fiber.Map{
			"subject":       cert.Subject.CommonName,
			"issuer":        cert.Issuer.CommonName,
			"not_before":    cert.NotBefore,
			"not_after":     cert.NotAfter,
			"expires_in":    time.Until(cert.NotAfter).String(),
			"serial_number": cert.SerialNumber.String(),
		})
	})

//
	//app.Get("/secure", func(c *fiber.Ctx) error { return c.SendString("This is a secure server 👮") })

//	app.Static("/", "/var/www/secure")
// ----------------------------------------------

	app.Use(func(c *fiber.Ctx) error {
		host := c.Hostname()

		switch host {
		case "spora.us.to":
			return c.SendFile("./static/spora.us.to" + c.Path())
		case "zpoc.soon.it":
			return c.SendFile("./static/zpoc.soon.it" + c.Path())
		case "golang.soon.it":
			return c.SendFile("./static/golang.soon.it" + c.Path())
		case "plh.fr.to":
			return c.SendFile("./static/plh.fr.to" + c.Path())
		default:
			return c.SendFile("./static/" + c.Path())
		}
	})
// ----------------------------------------------
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("spora.us.to","zpoc.soon.it","golang.soon.it","plh.fr.to"),
		Cache: autocert.DirCache("./certs"),
	}

	cfg := &tls.Config{ GetCertificate: m.GetCertificate, NextProtos: []string{ "http/1.1", "acme-tls/1", }, }
	ln, err := tls.Listen("tcp", ":443", cfg)
	if err != nil { panic(err) }

	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(404) 
	// => 404 "Not Found" 
	})

	log.Fatal(app.Listener(ln))
}
