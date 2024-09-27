// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

package db_models

import (

	// Third-party libraries
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	// System Packages
	"time"
)

// ----------------------------------------------
//
//	Structures
//
// ----------------------------------------------
type SystemError struct {
	gorm.Model
	ClientID     uint   `json:"client_id"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	ObjectDump   string `json:"object_dump"`
	Function     string `json:"function"`
}

type SystemError_CreateJSON struct {
	ClientID     uint   `json:"client_id"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	ObjectDump   string `json:"object_dump"`
	Function     string `json:"function"`
}

type SystemError_ReadJSON struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	ClientID     uint      `json:"client_id"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	ObjectDump   string    `json:"object_dump"`
	Function     string    `json:"function"`
}

// Initialize the error
func Model_SystemError_CreateObject(error_to_add *SystemError_CreateJSON) (int, *SystemError) {

	db := db_Database()
	defer db.Close()

	var system_error SystemError

	system_error = SystemError{
		ClientID:     error_to_add.ClientID,
		ErrorCode:    error_to_add.ErrorCode,
		ErrorMessage: error_to_add.ErrorMessage,
		ObjectDump:   error_to_add.ObjectDump,
		Function:     error_to_add.Function,
	}

	db.Save(&system_error)
	return DB_ANSWER_SUCCESS, &system_error
}

// Getting errors
func Model_SystemError_GetList() []SystemError_ReadJSON {

	db := db_Database()
	defer db.Close()

	var errors []SystemError
	db.Find(&errors)

	errors_list := make([]SystemError_ReadJSON, 0)

	if len(errors) <= 0 {
		return errors_list
	}

	for i := range errors {

		current_error := SystemError_ReadJSON{
			ID:           errors[i].ID,
			CreatedAt:    errors[i].CreatedAt,
			ClientID:     errors[i].ClientID,
			ErrorCode:    errors[i].ErrorCode,
			ErrorMessage: errors[i].ErrorMessage,
			ObjectDump:   errors[i].ObjectDump,
			Function:     errors[i].Function,
		}
		errors_list = append(errors_list, current_error)
	}
	return errors_list
}

// Getting error by id
func Model_SystemError_GetObject_byID(error_id uint) (int, *SystemError_ReadJSON) {

	db := db_Database()
	defer db.Close()

	var err SystemError
	db.Where("id = ?", error_id).First(&err)
	if err.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND, nil
	}

	err_to_show := SystemError_ReadJSON{
		ID:           err.ID,
		CreatedAt:    err.CreatedAt,
		ClientID:     err.ClientID,
		ErrorCode:    err.ErrorCode,
		ErrorMessage: err.ErrorMessage,
		ObjectDump:   err.ObjectDump,
		Function:     err.Function,
	}

	return DB_ANSWER_SUCCESS, &err_to_show
}

// Removing errors by id
func Model_SystemError_DeleteObject_byID(error_id uint) int {

	db := db_Database()
	defer db.Close()

	var err SystemError
	db.Where("id = ?", error_id).First(&err)
	if err.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND
	}

	db.Unscoped().Delete(&err)
	return DB_ANSWER_SUCCESS
}
