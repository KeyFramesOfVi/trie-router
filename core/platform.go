package core

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// This anonymous function allows me to connect to Postgres
	_ "github.com/lib/pq"

	"github.com/victor-cabrera/fe-calc/server/env"
	"github.com/victor-cabrera/fe-calc/server/logger"
	"github.com/victor-cabrera/fe-calc/server/router"
)

// Platform is the basjs of where the server initializes
type Platform struct {
	*logger.Logger
	config *env.Config
	db     *sql.DB
}

// NewPlatform returns a new Platform for the server
func NewPlatform() *Platform {
	p := &Platform{
		Logger: logger.NewLogger(os.Stdout),
		config: env.NewConfig(),
	}

	p.Connect()

	// p.Setup()

	// p.Api = NewAPI(p.config, p.db)

	if p.config.DevEnv {
		p.Info("Server running in Development mode")
	} else {
		p.Info("Server running in Production mode")
	}

	return p
}

// Connect links our server to our DB
func (p *Platform) Connect() {
	var (
		err    error
		dbInfo string
	)

	dbInfo = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		p.config.DB.Host,
		p.config.DB.Username,
		p.config.DB.Password,
		p.config.DB.Name,
		"disable",
	)

	fmt.Println(dbInfo)

	p.db, err = sql.Open(env.DbDriver, dbInfo)
	if err != nil {
		panic(err)
	}

	err = p.db.Ping()
	if err != nil {
		panic(err)
	}

	p.Success("Successful Connection -> %s", p.config.DB.Host)
}

// Start begins running the server.
func (p *Platform) Start() {
	mux := router.NewRouter()
	// mux.ServeDir("../static/build")
	mux.Get("/", indexHandler)
	mux.Post("/post", postHandler)
	mux.Put("/put", putHandler)
	mux.Patch("/patch", patchHandler)
	mux.Delete("/delete", deleteHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))

	// mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(buildDir, "static")))))
}
