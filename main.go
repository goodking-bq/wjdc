package main

import (
	"embed"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/goodking-bq/wjdc/utils"
	"github.com/goodking-bq/wjdc/web"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Embed a directory
////go:embed ui/*.html
//go:embed ui/static/bootstrap-4.6.1-dist/css/bootstrap.min.css
//go:embed ui/static/fontawesome-free-5.15.4-web/css/all.min.css
//go:embed ui/static/bootstrap-table-1.19.1/bootstrap-table.min.css
//go:embed ui/static/jquery-3.6.0.min.js
//go:embed ui/static/bootstrap-table-1.19.1/bootstrap-table.min.js
//go:embed ui/static/bootstrap-4.6.1-dist/js/bootstrap.bundle.min.js
//go:embed ui/static/bootstrap-table-1.19.1/locals/bootstrap-table-zh-CN.min.js
//go:embed ui/static/select2-4.0.13/js/select2.full.min.js
//go:embed ui/static/select2-4.0.13/css/select2.min.css
//go:embed ui/static/select2-4.0.13/css/select2-bootstrap4.min.css
//go:embed ui/static/bootstrap-4.6.1-dist/css/bootstrap.min.css.map
//go:embed ui/static/bootstrap-4.6.1-dist/js/bootstrap.min.js.map
var embedDirStatic embed.FS

//go:embed ui/*.tpl
var indexTpl embed.FS

func main() {
	utils.InitDB(utils.MysqlConfig{
		Host:     utils.GetEnvString("DB_HOST", "172.0.4.101"),
		Port:     "3306",
		DB:       utils.GetEnvString("DB_NAME", "wjdc"),
		User:     utils.GetEnvString("DB_USER", "root"),
		Password: utils.GetEnvString("DB_PASSWORD", "root"),
		Options:  "charset=utf8&parseTime=true",
	})

	engine := html.NewFileSystem(http.FS(indexTpl), ".tpl")
	app := fiber.New(
		fiber.Config{Views: engine})
	app.Use(
		logger.New(logger.Config{
			Format:     "[${time}] ${status} - ${latency} ${method} ${url}\n",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Shanghai",
		}),
		compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	h := new(web.Handlers)
	app.Get("/home", h.Index)
	app.Post("/home", h.SaveAnswer)
	app.Get("/csv", h.Csv)
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "ui",
		Browse:     true,
	}))
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Redirect("/home")
		})
	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
