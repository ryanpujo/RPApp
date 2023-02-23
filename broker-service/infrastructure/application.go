package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type application struct {
	Cfg config
}

func Application() application {
	return application{
		Cfg: config{
			Port: os.Getenv("PORT"),
		},
	}
}

func (app *application) Serve(mux http.Handler) error {

	srv := http.Server{
		Addr:              fmt.Sprintf(":%s", app.Cfg.Port),
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       20 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return srv.ListenAndServe()
}
