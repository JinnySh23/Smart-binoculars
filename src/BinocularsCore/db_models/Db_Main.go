// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

//
// ----------------------------------------------------------------------------------
//
// 										DB MAIN
//
// ----------------------------------------------------------------------------------
//

package db_models

import (
	// Internal project packages
	"rr/BinocularsCore/config"

	// Third-party libraries
	"github.com/jinzhu/gorm"

	// System Packages
	"crypto/sha1"
	"fmt"
)

// ----------------------------------------------
//
// 				(Base) General functionality
//
// ----------------------------------------------

// DATABASE initialization
func DB_Init() {
	db := db_Database()

	// Migration(setup)
	db.AutoMigrate(&Device{})
	db.AutoMigrate(&SystemError{})
}

// Output a debugging message To the CONSOLE if we are debugging
func db_LOG(message string) {
	if config.CONFIG_IS_DEBUG {
		fmt.Println("[DB]: " + message)
	}
}

// Database connection function
func db_Database() *gorm.DB {

	db, err := gorm.Open("sqlite3", config.CONFIG_DB_FILE)
	// db_credentials := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",config.CONFIG_DB_HOST, config.CONFIG_DB_PORT, config.CONFIG_DB_USER, config.CONFIG_DB_NAME, config.CONFIG_DB_PASSWORD)
	// db, err := gorm.Open("postgres", db_credentials)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}

//
// Other functions
//

// Get a SHA1 hash for passwords
func getSHA1Hash(input_string string) string {

	hash := sha1.New()
	hash.Write([]byte(input_string))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)

}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for i := range a {
		if x == a[i] {
			return true
		}
	}
	return false
}
