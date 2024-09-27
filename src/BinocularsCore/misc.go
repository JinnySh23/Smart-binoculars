// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

//
// ----------------------------------------------------------------------------------
//
// 											Misc
//
// ----------------------------------------------------------------------------------
//

package main

import (
	// Internal project packages
	"rr/BinocularsCore/config"

	// Third-party libraries
	// "github.com/gin-gonic/gin"

	// System Packages
	"crypto/sha1"
	"fmt"
	"strconv"
)

// Get a UINT from a string
func get_uint_fromString(str string) (uint, bool) {

	id_uint64, err := strconv.ParseUint(str, 10, 0)

	if err != nil {
		LOG("INVALID ID CONVERSION!")
		return 0, false
	}

	return uint(id_uint64), true
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

// isExist_int64tells whether a contains x (for int64 array).
func isExist_int64(a []int64, x int64) int {
	for i := range a {
		if x == a[i] {
			return i
		}
	}
	return -1
}

// Output a debugging message To the CONSOLE if we are debugging
func LOG(message string) {
	if config.CONFIG_IS_DEBUG {
		fmt.Println("[DEBUG]: " + message)
	}
}

// Get the SHA1 hash
func getSHA1Hash(input_string string) string {

	hash := sha1.New()
	hash.Write([]byte(input_string))
	bs := hash.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
