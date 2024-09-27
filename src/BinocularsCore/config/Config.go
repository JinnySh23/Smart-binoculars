// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

package config

const (

	//
	// General configs
	//

	// Basic real URL
	CONFIG_URL_BASE = "YOUR_URL"

	// The URL for the redirect after successful verification
	CONFIG_VERIFY_REDIRECT_URL = "YOUR_URL/thankyou"

	// Ports for HTTP
	CONFIG_RELEASE_SERVER_PORT          = "80.."
	CONFIG_DEBUG_SERVERLESS_SERVER_PORT = "80.."

	// Debugging level
	CONFIG_DEBUG_LEVEL = 1

	// Debugging mode
	CONFIG_IS_DEBUG = true

	// Using an internal server (for debugging)
	CONFIG_IS_DEBUG_SERVERLESS = true

	// The key for encrypting tokens (secret)
	CONFIG_SECRET = "YOUR_SECRET_KEY"

	CONFIG_SERVER_ACCESS_TOKEN = "YOUR_ACCESS_TOKEN"

	CONFIG_DEFAULT_LOGIN    = "YOUR_LOGIN"
	CONFIG_DEFAULT_PASSWORD = "YOUR_PASSWORD"

	//
	// SMTP
	//

	// Mail Settings
	CONFIG_SMTP_HOST      = "YOUR_SMTP_HOST"
	CONFIG_SMTP_PORT      = "YOUR_SMTP_PORT"
	CONFIG_SMTP_USER      = "YOUR_SMTP_USER"
	CONFIG_SMTP_USER_NAME = "YOUR_SMTP_USER_NAME"
	CONFIG_SMTP_PASSWORD  = "YOUR_SMTP_PASSWORD"
	CONFIG_SMTP_TITLE     = "YOUR_SMTP_TITLE"

	// Message Template
	// CONFIG_SMTP_TEMPLATE_ORDER_MADE = "YOUR_PATH_TO_THE_MESSAGE_TEMPLATE"

	//
	// Database Settings
	//

	// DB Type
	CONFIG_DB_TYPE = "sqlite3"
	CONFIG_DB_FILE = "app.db"

	//
	// File Storage
	//
	FILE_STORAGE_ROOT = "content"
)
