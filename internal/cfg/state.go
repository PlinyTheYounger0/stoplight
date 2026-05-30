package cfg

import (
	"database/sql"

	"github.com/PlinyTheYounger0/stoplight/internal/database"
)

type State struct {
	Queries *database.Queries
	DB *sql.DB
}
