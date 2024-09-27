// ------------------------------------
// RR IT Crew 2021
//
// ------------------------------------
// Базовый движок для Биноклей

//
// ----------------------------------------------------------------------------------
//
// 								Routes (Пути)
//
// ----------------------------------------------------------------------------------
//

package routes

import (
	//Внутренние пакеты проекта
	"rr/BinocularsCore/config"

	// "../modules/rr_randstr"

	//Сторонние библиотеки
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	//Системные пакеты
	"fmt"
	"net/http"
	"strconv"
)

// ----------------------------------------------
//
// 				Структуры
//
// ----------------------------------------------
type Admin_LoginJSON struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// ----------------------------------------------
//
// 				Root requests
//
// ----------------------------------------------
//HTML-пути

//	/
func Handler_Index(c *gin.Context) {
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {
		c.HTML(http.StatusOK, "my_profile.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

//	/adm-panel
func Handler_AdmPanel(c *gin.Context) {
	// Проверяем сессию, если она есть - перебрасываем на страницу, если нет - на логин
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {
		c.HTML(http.StatusOK, "adm_panel.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

//	/gate-lord
func Handler_GateLord(c *gin.Context) {
	// Проверяем сессию, если она есть - перебрасываем на страницу, если нет - на логин
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id != nil {
		c.HTML(http.StatusOK, "gate_lord.html", gin.H{})
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

//	/cust-app
func Handler_CustApp(c *gin.Context) {
	c.HTML(http.StatusOK, "cust_app.html", gin.H{})
}

//	/rates-app
func Handler_RatesApp(c *gin.Context) {
	c.HTML(http.StatusOK, "rates_app.html", gin.H{})
}

//	/feedback-app
func Handler_FeedbackApp(c *gin.Context) {
	c.HTML(http.StatusOK, "feedback_app.html", gin.H{})
}

//	/payment-res
func Handler_PaymentRes(c *gin.Context) {
	c.HTML(http.StatusOK, "payment_res.html", gin.H{})
}

//	/device-emulator
func Handler_DeviceEmulator(c *gin.Context) {
	c.HTML(http.StatusOK, "device_emulator.html", gin.H{})
}

//	/login
func Handler_Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", gin.H{})
		return
	}

	if c.Request.Method == "POST" {
		json_data := new(Admin_LoginJSON)
		err := c.ShouldBindJSON(&json_data)

		//Проверка, JSON пришел или шляпа
		if err != nil {
			if config.CONFIG_IS_DEBUG {
				error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_Login")
				Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message+" Error: "+err.Error(), error_id)
			} else {
				error_id := generatorErrorID(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, json_data, "Handler_Login")
				Answer_BadRequest(c, ANSWER_INVALID_JSON().Code, ANSWER_INVALID_JSON().Message, error_id)
			}
			return
		}

		if json_data.Login == config.CONFIG_DEFAULT_LOGIN && json_data.Password == config.CONFIG_DEFAULT_PASSWORD {
			session := sessions.Default(c)
			session.Set("session_user_id", 1)
			session.Save()
			Answer_OK(c)
		} else {
			error_id := generatorErrorID(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, json_data, "Handler_Login")
			Answer_NotFound(c, ANSWER_OBJECT_NOT_FOUND().Code, ANSWER_OBJECT_NOT_FOUND().Message, error_id)
		}
		return
	}
}

// Выход
func Handler_Logout(c *gin.Context) {
	session := sessions.Default(c)
	session_user_id := session.Get("session_user_id")
	if session_user_id == nil {
		error_id := generatorErrorID(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, nil, "Handler_Logout")
		Answer_Unauthorized(c, ANSWER_INVALID_SESSION().Code, ANSWER_INVALID_SESSION().Message, error_id)
		return
	} else {
		session.Clear()
		Answer_OK(c)
		return
	}
}

func get_uint_fromString(str string) (uint, bool) {
	//Вот такая вот странная передача данных
	id_uint64, err := strconv.ParseUint(str, 10, 0)

	//Неверная трансформация строки в число
	if err != nil {
		LOG("INVALID ID CONVERSION!")
		return 0, false
	}

	return uint(id_uint64), true
}

//
// Вывод отладочного сообщения В КОНСОЛЬ, если у нас отладка
//
func LOG(message string) {
	if config.CONFIG_IS_DEBUG {
		fmt.Println("[DEBUG]: " + message)
	}
}
