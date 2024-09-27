// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

//
// ----------------------------------------------------------------------------------
//
// 								DB Answer Codes
//
// ----------------------------------------------------------------------------------
//

package db_models

const (
	// Standard answers
	DB_ANSWER_SUCCESS                 = 0 // Successful request
	DB_ANSWER_OBJECT_EXISTS           = 1 // The requested object exists
	DB_ANSWER_OBJECT_NOT_FOUND        = 2 // The requested object does NOT exist
	DB_ANSWER_INVALID_CREDENTIALS     = 3 // Invalid access parameters (login, password, token)
	DB_ANSWER_PERMISSION_DENIED       = 4 // Access denied
	DB_ANSWER_UNEXPECTED_ERROR        = 5 // An undefined error
	DB_ANSWER_INVALID_JSON_CONVERSION = 6 // Error in converting from JSON to string

	// Custom answers (start with 500)
	DB_ANSWER_INVALID_COMMAND = 501 // Incorrect command

)
