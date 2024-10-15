package models

import "strconv"

type Sample struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Value        string `json:"value"`
	LatestUpdate string `json:"latest_update"`
}

func SelectAllSamples() ([]Sample, error) {
	db := Connection()
	defer db.Close()

	sql := "select * from my_table;"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	var samples []Sample
	for rows.Next() {
		var sample Sample
		rows.Scan(
			&sample.ID,
			&sample.Name,
			&sample.Value,
			&sample.LatestUpdate,
		)
		samples = append(samples, sample)
	}

	return samples, nil
}

func SelectSampleByID(id string) (Sample, error) {
	db := Connection()
	defer db.Close()

	var sample Sample
	sql := "select * from my_table where id = ?;"
	err := db.QueryRow(sql, id).Scan(
		&sample.ID,
		&sample.Name,
		&sample.Value,
		&sample.LatestUpdate,
	)

	if err != nil {
		return sample, err
	}

	return sample, nil
}

func InsertSample(data *Sample) (string, error) {
	db := Connection()
	defer db.Close()

	dbh, err := db.Prepare("insert into my_table values(0, ?, ?, now());")
	if err != nil {
		return "", err
	}
	result, err := dbh.Exec(data.Name, data.Value)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	return strconv.FormatInt(id, 10), nil
}

func UpdateSample(data *Sample) error {
	db := Connection()
	defer db.Close()

	dbh, err := db.Prepare("update my_table set name = ?, value = ?, latest_update = now() where id = ?;")
	if err != nil {
		return err
	}

	_, err = dbh.Exec(data.Name, data.Value, data.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteSample(id string) error {
	db := Connection()
	defer db.Close()

	dbh, err := db.Prepare("delete from my_table where id = ?;")
	if err != nil {
		return err
	}

	_, err = dbh.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
