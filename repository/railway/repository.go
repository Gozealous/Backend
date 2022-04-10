package railway

import (
	"database/sql"
	"gozealous/anomaly"
	"gozealous/code"
	"gozealous/repository/common"
)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		database: db,
	}
}

type Repository struct {
	database *sql.DB
}

func (r *Repository) Stations() ([]Station, *anomaly.ServiceError) {
	stations := []Station{}

	rows, queryErr := r.database.Query("select name, code, number from station order by name, code, number")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get stations.", queryErr)
		return stations, err.Trace()
	}
	defer rows.Close()

	for rows.Next() {
		var (
			stationName   string
			stationCode   string
			stationNumber int
		)

		if scanErr := rows.Scan(&stationName, &stationCode, &stationNumber); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a station.", scanErr)
			return stations, err.Trace()
		}

		station := &Station{
			Name:   stationName,
			Code:   stationCode,
			Number: stationNumber,
		}

		stations = append(stations, *station)
	}

	if rowErr := rows.Err(); rowErr != nil {
		err := common.ParseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return stations, err.Trace()
	}

	return stations, nil
}
