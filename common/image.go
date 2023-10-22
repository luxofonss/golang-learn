package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name" gorm:"-"`
	Extension string `json:"extension" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image

	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img

	return nil
}

// value return json value, implement driver.Value interface
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var imgs Images

	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}

	*j = imgs

	return nil
}

func (j *Images) Value() (driver.Value, error) {
	if j != nil {
		return nil, nil
	}

	return json.Marshal(j)
}
