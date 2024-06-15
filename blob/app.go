package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	e "github.com/bruno5200/CSM/env"
	u "github.com/bruno5200/CSM/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var en = e.Env()

func init() { e.Init() }

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	// CORS for external resources
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
		AllowHeaders: "Cache-Control, Accept, Content-Type, Content-Length, Accept-Encoding, Authorization",
	}))

	app.Use(recover.New())
	app.Use(favicon.New())
	app.Use(logger.New())
	app.Use(requestid.New())

	app.Get("/metrics", monitor.New(monitor.Config{
		Title:   "Blob Storage API Metrics",
		Refresh: 2 * time.Second,
	}))

	app.Static("/", "./files/", fiber.Static{
		// Download:      true,
		CacheDuration: 10 * time.Second,
		MaxAge:        10,
	})

	const documents = "/documents"

	app.Static(documents, documents, fiber.Static{
		// Download:      true,
		CacheDuration: 10 * time.Second,
		MaxAge:        10,
	})

	const images = "/images"

	app.Static(images, images, fiber.Static{
		// Download:      true,
		CacheDuration: 10 * time.Second,
		MaxAge:        10,
	})

	// app.Get(documents+"/:file", func(c *fiber.Ctx) error {
	// 	file := c.Params("file")
	// 	if _, err := os.Stat(fmt.Sprintf("./files/%s", file)); os.IsNotExist(err) {
	// 		return c.Status(fiber.StatusNotFound).SendString(`<h1>Error 404, fichero no encontrado</h1>`)
	// 	}
	// 	return c.SendFile(fmt.Sprintf("./files/%s", file))
	// })

	// app.Get(documents+"/:file", func(c *fiber.Ctx) error {
	// 	file := c.Params("file")
	// 	if _, err := os.Stat(fmt.Sprintf(".%s/%s", documents, file)); os.IsNotExist(err) {
	// 		c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	// 		return c.Status(fiber.StatusNotFound).SendString(`<h1>Error 404,<br> fichero no encontrado</h1>`)
	// 	}
	// 	return c.SendFile(fmt.Sprintf(".%s/%s", documents, file))
	// })

	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := u.CheckDir("/files"); err != nil {
			return err
		}

		log.Printf("File extension: %s", filepath.Ext(file.Filename))

		fileName := fmt.Sprintf("/files/%s", file.Filename)

		// Check if file exists
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			// Remove
			if err := os.Remove(fileName); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}

		if err := c.SaveFile(file, fileName); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
	})

	app.Post(documents, func(c *fiber.Ctx) error {
		file, err := c.FormFile("document")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := u.CheckDir(documents); err != nil {
			return err
		}

		log.Printf("File extension: %s", filepath.Ext(file.Filename))

		fileName := fmt.Sprintf("%s/%s", documents, file.Filename)
		// Check if file exists
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			// Remove
			if err := os.Remove(fileName); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}

		if err := c.SaveFile(file, fileName); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
	})

	app.Post(images, func(c *fiber.Ctx) error {
		file, err := c.FormFile("image")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := u.CheckDir(images); err != nil {
			return err
		}

		log.Printf("File extension: %s", filepath.Ext(file.Filename))

		fileName := fmt.Sprintf("%s/%s", images, file.Filename)

		// Check if file exists

		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			// Remove
			if err := os.Remove(fileName); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}

		if err := c.SaveFile(file, fileName); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true})
	})

	app.Get("/*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString(`<h1>404 No Encontrado</h1>`)
	})

	port := ":" + en.GetPort()

	if en.GetSecure() {
		go func() {
			log.Printf(`Running with TLS in https://localhost%v`, port)
			if err := app.ListenTLS(
				port,
				"./certs/tickets.cert",
				"./certs/tickets.key",
			); err != nil {
				log.Panic(err)
			}
		}()
	} else {
		go func() {
			log.Printf(`Running in http://localhost%v`, port)
			if err := app.Listen(port); err != nil {
				log.Panic(err)
			}
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")
	//database.CloseDb()
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}
