package model

import (
	"fmt"
	"meshalka/database"
)

func getBoolResult(res int, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	return res != 0, nil
}

func getIntFunctionResult(db database.Database, q querier) (int, error) {
	con, err := db.Connection()
	if err != nil {
		return 0, err
	}

	rows, err := q(con)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var result int
	if rows.Next() {
		if rows.Scan(&result) == nil {
			return result, nil
		}
	}

	return 0, fmt.Errorf("function error")
}