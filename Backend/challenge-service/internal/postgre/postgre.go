package postgre

import (
	"database/sql"
	"fmt"
	"strconv"

	"backend/challenge-service/config"
)

func NewDBConnection(configdata *config.Config) (*sql.DB, error) {
	dbInfo := "host=" + configdata.Database.Host +
		" port=" + strconv.Itoa(configdata.Database.Port) +
		" user=" + configdata.Database.User +
		" password=" + configdata.Database.Password +
		" dbname=" + configdata.Database.DBName +
		" sslmode=" + configdata.Database.SSLMode

	var err error
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	return db, nil
}
