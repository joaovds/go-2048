package logic

import (
	"encoding/json"
	"os"
	"time"
)

const filename = "data.json"

type Record struct {
	Score    int           `json:"score"`
	Time     time.Duration `json:"time"`
	Datetime *time.Time    `json:"datetime"`
}

type Data struct {
	Record *Record `json:"record"`
}

func NewData() *Data {
	return &Data{Record: new(Record)}
}

func (d *Data) GetGameData() error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		defaultData := &Data{
			Record: &Record{
				Score:    0,
				Time:     0,
				Datetime: nil,
			},
		}
		file, err := json.MarshalIndent(defaultData, "", "  ")
		if err != nil {
			return err
		}

		err = os.WriteFile(filename, file, 0644)
		if err != nil {
			return err
		}
	}

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	data := new(Data)
	json.Unmarshal(fileBytes, data)

	d.Record = data.Record

	return nil
}

func (d *Data) SaveGameRecord(newRecord *Data) error {
	data, err := json.MarshalIndent(&newRecord, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) IsNewRecord(newRecordScore int, newRecordTime time.Duration) bool {
	if newRecordScore >= d.Record.Score {
		if newRecordTime <= d.Record.Time {
			return true
		}
	}

	return false
}
