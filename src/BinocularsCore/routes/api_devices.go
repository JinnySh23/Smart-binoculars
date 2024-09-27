// ------------------------------------
// RR IT Crew 2021
//
// ------------------------------------
// Базовый движок для Биноклей

//
// ----------------------------------------------------------------------------------
//
// 								Devices (Пути)
//
// ----------------------------------------------------------------------------------
//

package routes

import (

	//Внутренние пакеты проекта
	"rr/BinocularsCore/config"
	"rr/BinocularsCore/db_models"

	//Сторонние библиотеки
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	//Системные пакеты
)

// ======================================================================
//								CRUD
// ======================================================================

// CREATE
func Handler_API_Devices_CreateObject(c *gin.Context) {
	json_data := new(db_models.Device_CreateJSON)
	err := c.ShouldBindJSON(&json_data)

	//Проверка, JSON пришел или шляпа
	if err != nil {
		if config.CONFIG_IS_DEBUG {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_CreateObject")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err.Error(), error_id)
		} else {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_CreateObject")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, error_id)
		}
		return
	}

	//Проверка на необходимые поля
	if (json_data.UUID == "") && (json_data.AccessToken == "") {
		error_id := generatorErrorID(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message, json_data, "Handler_API_Devices_CreateObject")
		Answer_BadRequest(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message, error_id)
		return
	}

	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id == nil {
		error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, json_data, "Handler_API_Devices_CreateObject")
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
		return
	} else {
		db_answer, device := db_models.Model_Device_CreateObject(json_data)
		switch db_answer {
		case db_models.DB_ANSWER_SUCCESS:
			Answer_SendObject(c, device)

		case db_models.DB_ANSWER_OBJECT_EXISTS:
			error_id := generatorErrorID(c, ANSWER_OBJECT_EXISTS().Code, ANSWER_OBJECT_EXISTS().Message, device, "Handler_API_Devices_CreateObject")
			Answer_BadRequest(c, ANSWER_OBJECT_EXISTS().Code, ANSWER_OBJECT_EXISTS().Message, error_id)

		default:
			error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, device, "Handler_API_Devices_CreateObject")
			Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
		}
		return
	}
}

// READ LIST
func Handler_API_Devices_GetList(c *gin.Context) {

	//Переменная для списка постов
	var devices_list []db_models.Device_ReadJSON

	// Если чужие посты
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id == nil {
		error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, nil, "Handler_API_Devices_GetList")
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
		return
	} else {
		devices_list = db_models.Model_Device_GetList()
	}

	Answer_SendObject(c, devices_list)
	return
}

// READ Object
func Handler_API_Devices_GetObject(c *gin.Context) {
	//Получаем get-параметр uuid
	device_uuid_str := c.Query("uuid")
	db_answer, device := db_models.Model_Device_GetObject_byUUID(device_uuid_str)

	switch db_answer {
	case db_models.DB_ANSWER_SUCCESS:
		Answer_SendObject(c, device)

	case db_models.DB_ANSWER_OBJECT_NOT_FOUND:
		error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, device, "Handler_API_Devices_GetObject")
		Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)

	case db_models.DB_ANSWER_PERMISSION_DENIED:
		error_id := generatorErrorID(c, ANSWER_PERMISSION_DENIED().Code, ANSWER_PERMISSION_DENIED().Message, device, "Handler_API_Devices_GetObject")
		Answer_Forbidden(c, ANSWER_PERMISSION_DENIED().Code, ANSWER_PERMISSION_DENIED().Message, error_id)

	default:
		error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, device, "Handler_API_Devices_GetObject")
		Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
	}
	// session := sessions.Default(c)
	// session_user_id := session.Get("session_user_id")
	// if session_user_id == nil {
	// 	error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, nil, "Handler_API_Devices_GetObject")
	// 	Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
	// 	return
	// } else {

	// 	//Получаем get-параметр uuid
	// 	device_uuid_str := c.Query("uuid")
	// 	db_answer, device := db_models.Model_Device_GetObject_byUUID(device_uuid_str)

	// 	switch db_answer {
	// 	case db_models.DB_ANSWER_SUCCESS:
	// 		Answer_SendObject(c, device)

	// 	case db_models.DB_ANSWER_OBJECT_NOT_FOUND:
	// 		error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, device, "Handler_API_Devices_GetObject")
	// 		Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)

	// 	case db_models.DB_ANSWER_PERMISSION_DENIED:
	// 		error_id := generatorErrorID(c, ANSWER_PERMISSION_DENIED().Code, ANSWER_PERMISSION_DENIED().Message, device, "Handler_API_Devices_GetObject")
	// 		Answer_Forbidden(c, ANSWER_PERMISSION_DENIED().Code, ANSWER_PERMISSION_DENIED().Message, error_id)

	// 	default:
	// 		error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, device, "Handler_API_Devices_GetObject")
	// 		Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
	// 	}
	// 	return
	// }
}

// UPDATE
func Handler_API_Devices_UpdateObject(c *gin.Context) {

	json_data := new(db_models.Device_UpdateJSON)
	err := c.ShouldBindJSON(&json_data)

	//Проверка, JSON пришел или шляпа
	if err != nil {
		if config.CONFIG_IS_DEBUG {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_UpdateObject")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err.Error(), error_id)
		} else {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_UpdateObject")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, error_id)
		}
		return
	}

	//Проверка на необходимые поля
	if (json_data.UUID == "") && (json_data.AccessToken == "") {
		error_id := generatorErrorID(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message, json_data, "Handler_API_Devices_UpdateObject")
		Answer_BadRequest(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message, error_id)
		return
	}

	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id == nil {
		error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, json_data, "Handler_API_Devices_UpdateObject")
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
		return
	} else {
		db_answer, new_device := db_models.Model_Device_UpdateObject(json_data)
		switch db_answer {
		case db_models.DB_ANSWER_SUCCESS:
			Answer_SendObject(c, new_device)

		case db_models.DB_ANSWER_OBJECT_NOT_FOUND:
			error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, new_device, "Handler_API_Devices_UpdateObject")
			Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)

		default:
			error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, new_device, "Handler_API_Devices_UpdateObject")
			Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
		}
		return
	}
}

// DELETE
func Handler_API_Devices_DeleteObject(c *gin.Context) {
	device_id_str := c.Param("uuid")

	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")

	if session_user_id == nil {
		error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, nil, "Handler_API_Devices_DeleteObject")
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
		return
	} else {

		if device_id_str == "0" {
			db_answer := db_models.Model_Device_DeleteObject_All()

			switch db_answer {
			case db_models.DB_ANSWER_SUCCESS:
				Answer_OK(c)

			default:
				error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, nil, "Handler_API_Devices_DeleteObject")
				Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
			}
			return
		} else {
			db_answer := db_models.Model_Device_DeleteObject_byUUID(device_id_str)

			switch db_answer {
			case db_models.DB_ANSWER_SUCCESS:
				Answer_OK(c)

			case db_models.DB_ANSWER_OBJECT_NOT_FOUND:
				error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, nil, "Handler_API_Devices_DeleteObject")
				Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)

			default:
				error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, nil, "Handler_API_Devices_DeleteObject")
				Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
			}
		}
	}
	return
}

// ======================================================================
//								Additionally
// ======================================================================
func Handler_API_Devices_State(c *gin.Context) {

	device_uuid_str := c.Param("uuid")

	json_data := new(db_models.Device_Sync_RequestJSON)
	err := c.ShouldBindJSON(&json_data)

	//Проверка, JSON пришел или шляпа
	if err != nil {
		if config.CONFIG_IS_DEBUG {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_State")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err.Error(), error_id)
		} else {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_State")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, error_id)
		}
		return
	}

	//Проверка на необходимые поля
	if json_data.AccessToken == "" {
		error_id := generatorErrorID(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message, json_data, "Handler_API_Devices_State")
		Answer_BadRequest(c, ANSWER_EMPTY_FIELDS().Code, ANSWER_EMPTY_FIELDS().Message, error_id)
		return
	}

	db_answer, device_status := db_models.Model_Device_State(device_uuid_str, json_data)
	switch db_answer {
	case db_models.DB_ANSWER_SUCCESS:
		Answer_SendObject(c, device_status)

	case db_models.DB_ANSWER_OBJECT_NOT_FOUND:
		error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, device_status, "Handler_API_Devices_State")
		Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)

	case db_models.DB_ANSWER_INVALID_CREDENTIALS:
		error_id := generatorErrorID(c, ANSWER_INVALID_CREDENTIALS().Code, ANSWER_INVALID_CREDENTIALS().Message, device_status, "Handler_API_Devices_State")
		Answer_BadRequest(c, ANSWER_INVALID_CREDENTIALS().Code, ANSWER_INVALID_CREDENTIALS().Message, error_id)

	case db_models.DB_ANSWER_INVALID_JSON_CONVERSION:
		error_id := generatorErrorID(c, ANSWER_INVALID_JSON_CONVERSION().Code, ANSWER_INVALID_JSON_CONVERSION().Message, device_status, "Handler_API_Devices_State")
		Answer_BadRequest(c, ANSWER_INVALID_JSON_CONVERSION().Code, ANSWER_INVALID_JSON_CONVERSION().Message, error_id)

	default:
		error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, device_status, "Handler_API_Devices_State")
		Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
	}
	return
}

func Handler_API_Devices_Command(c *gin.Context) {
	device_uuid_str := c.Param("uuid")

	json_data := new(db_models.Device_CommandJSON)
	err := c.ShouldBindJSON(&json_data)
	//Проверка, JSON пришел или шляпа
	if err != nil {
		if config.CONFIG_IS_DEBUG {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_Command")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err.Error(), error_id)
		} else {
			error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_API_Devices_Command")
			Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, error_id)
		}
		return
	}

	// //Проверка на необходимые поля
	// if json_data.ServerAccessToken != config.CONFIG_SERVER_ACCESS_TOKEN {
	// 	Answer_BadRequest(c, ANSWER_PERMISSION_DENIED().Code, ANSWER_PERMISSION_DENIED().Message)
	// 	return
	// }

	// session := sessions.Default(c)
	// session_user_id := session.Get("session_user_id")
	// if session_user_id == nil {
	// 	error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, json_data, "Handler_API_Devices_Command")
	// 	Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
	// 	return
	// } else {
	db_answer := db_models.Model_Device_Command(device_uuid_str, json_data)
	switch db_answer {
	case db_models.DB_ANSWER_SUCCESS:
		Answer_OK(c)

	case db_models.DB_ANSWER_OBJECT_NOT_FOUND:
		error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, nil, "Handler_API_Devices_Command")
		Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)

	case db_models.DB_ANSWER_INVALID_COMMAND:
		error_id := generatorErrorID(c, ANSWER_INVALID_COMMAND().Code, ANSWER_INVALID_COMMAND().Message, nil, "Handler_API_Devices_Command")
		Answer_BadRequest(c, ANSWER_INVALID_COMMAND().Code, ANSWER_INVALID_COMMAND().Message, error_id)

	case db_models.DB_ANSWER_INVALID_JSON_CONVERSION:
		error_id := generatorErrorID(c, ANSWER_INVALID_JSON_CONVERSION().Code, ANSWER_INVALID_JSON_CONVERSION().Message, nil, "Handler_API_Devices_Command")
		Answer_BadRequest(c, ANSWER_INVALID_JSON_CONVERSION().Code, ANSWER_INVALID_JSON_CONVERSION().Message, error_id)

	default:
		error_id := generatorErrorID(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, nil, "Handler_API_Devices_Command")
		Answer_BadRequest(c, ANSWER_UNEXPECTED_ERROR().Code, ANSWER_UNEXPECTED_ERROR().Message, error_id)
	}
	return
	// }
}
