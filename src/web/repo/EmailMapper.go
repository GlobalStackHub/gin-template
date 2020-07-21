package repo

import "app-test/src/pkg/datasource"

// Insert
func InsertEmail(value interface{}) error {
	if datasource.DB == nil {
		panic("database still not initialize")
	}
	return datasource.DB.Create(value).Error
}
