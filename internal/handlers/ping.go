package handlers

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/http"
)

func (h *URLHandler) Ping(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", h.config.GetDataBaseDSN())
	if err != nil {
		h.log.Error("error connecting to the database", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		h.log.Error("error ping", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
