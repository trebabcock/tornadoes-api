package app

import (
	"log"
	"net/http"
	"time"
	"tornadoes/app/handler"
	"tornadoes/app/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Init initializes the database
func (a *App) Init() {
	log.Println("Connecting to database...")
	db, err := gorm.Open(sqlite.Open("my.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	log.Println("Connected to database")
	log.Println("Migrating database...")
	a.DB = model.DBMigrate(db)
	log.Println("Migrated database")
	a.Router = mux.NewRouter()
	a.setRoutes()
}

func (a *App) setRoutes() {
	a.get("/api/v1/tornadoesByRange", a.getTornadoesByRange)
	a.get("/api/v1/tornadoesByDate", a.getTornadoesByDate)
}

func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) getTornadoesByRange(w http.ResponseWriter, r *http.Request) {
	handler.GetTornadoesByRange(a.DB, w, r)
}

func (a *App) getTornadoesByDate(w http.ResponseWriter, r *http.Request) {
	handler.GetTornadoesByDate(a.DB, w, r)
}

// Run starts the server
func (a *App) Run(host string) {
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Origin", "Cache-Control", "X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	server := &http.Server{
		Addr:         host,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.CORS(headers, methods, origins)(a.Router),
	}
	log.Println("Server running at:", host)
	log.Fatal(server.ListenAndServe())
}
