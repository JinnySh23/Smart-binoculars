// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

package main

import (

	// Internal project packages
	"rr/BinocularsCore/config"
	"rr/BinocularsCore/db_models"
	"rr/BinocularsCore/middleware"
	"rr/BinocularsCore/routes"

	// Third-party libraries
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"

	// System Packages
	"io"
	"os"
)

//
// ----------------------------------------------------------------------------------
//
// 											MAIN
//
// ----------------------------------------------------------------------------------
//

func main() {

	// Initializing the database
	db_models.DB_Init()

	// THE LOG FILE, if we don't have debugging
	if !config.CONFIG_IS_DEBUG {
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()

		// Logging to a file.
		f, _ := os.Create("gin_server.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

	// Creating a router for processing requests
	r := gin.Default()

	// Distribution of static for the debug version
	if config.CONFIG_IS_DEBUG_SERVERLESS {
		r.Static("/assets", "./assets") // For static in debugging mode
		r.LoadHTMLGlob("assets/html/*")
	} else {
		r.LoadHTMLGlob("static/assets/html/*")
	}

	// For sessions
	store := memstore.NewStore([]byte(config.CONFIG_SECRET))
	r.Use(sessions.Sessions("data", store))

	//CORS
	r.Use(middleware.CORSMiddleware())

	//
	// Common paths
	//

	r.GET("/", routes.Handler_Index)
	// The main paths
	r.GET("/login", routes.Handler_Login)
	r.POST("/login", routes.Handler_Login)
	r.POST("/login/", routes.Handler_Login)

	r.GET("/cust-app", routes.Handler_CustApp)
	r.GET("/rates-app", routes.Handler_RatesApp)
	r.GET("/feedback-app", routes.Handler_FeedbackApp)
	r.GET("/payment-res", routes.Handler_PaymentRes)
	r.GET("/gate-lord", routes.Handler_GateLord)
	r.GET("/adm-panel", routes.Handler_AdmPanel)
	r.GET("/device-emulator", routes.Handler_DeviceEmulator)
	r.GET("/logout", routes.Handler_Logout)

	// API Group
	api := r.Group("/api")
	{
		//Devices
		devices := api.Group("/devices")
		{
			devices.GET("/", routes.Handler_API_Devices_GetList)
			devices.GET("/:uuid", routes.Handler_API_Devices_GetObject)
			devices.POST("/", routes.Handler_API_Devices_CreateObject)
			devices.POST("/:uuid/state", routes.Handler_API_Devices_State)
			devices.POST("/:uuid/command", routes.Handler_API_Devices_Command)
			devices.PUT("/", routes.Handler_API_Devices_UpdateObject)
			devices.DELETE("/:uuid", routes.Handler_API_Devices_DeleteObject)
		}
	}

	// Starting the server
	if config.CONFIG_IS_DEBUG_SERVERLESS {
		r.Run(":" + config.CONFIG_DEBUG_SERVERLESS_SERVER_PORT)
	} else {
		r.Run(":" + config.CONFIG_RELEASE_SERVER_PORT)
	}
}

//
// ----------------------------------------------------------------------------------
//
// 										/END OF	MAIN
//
// ----------------------------------------------------------------------------------
//
