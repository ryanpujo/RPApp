package infrastructure

import (
	"fmt"
	"net/http"
	"time"
)

type application struct {
	Cfg Config
}

func Application() application {

	return application{
		Cfg: LoadConfig(),
	}
}

func (app *application) Serve(mux http.Handler) error {

	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", app.Cfg.Port),
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       20 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return srv.ListenAndServe()
}
