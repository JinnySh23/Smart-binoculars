// ------------------------------------
// RR IT 2021
//
// ------------------------------------
// Basic engine for Binoculars

package db_models

import (

	// Third-party libraries
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	// System Packages
	"encoding/json"
	"time"
)

// ----------------------------------------------
//
//	Structures
//
// ----------------------------------------------
type Device struct {
	gorm.Model
	UUID          string    `json:"uuid"`           // Device ID
	AccessToken   string    `json:"access_token"`   // Device Token
	BatteryCharge string    `json:"battery_charge"` // Battery charge
	Location      string    `json:"location"`       // Device Location
	State         string    `json:"state"`          // Device status
	LastSync      time.Time `json:"last_sync"`      // The label of the last successful synchronization with the device
	Command       string    `json:"command"`        // Execution command
}

type Device_CreateJSON struct {
	UUID        string `json:"uuid"`
	AccessToken string `json:"access_token"`
	Location    string `json:"location"`
}

type Device_ReadJSON struct {
	ID            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UUID          string    `json:"uuid"`
	AccessToken   string    `json:"access_token"`
	Location      string    `json:"location"`
	BatteryCharge string    `json:"battery_charge"`
	State         string    `json:"state"`
	LastSync      time.Time `json:"last_sync"`
	Command       string    `json:"command"`
}

type Device_Sync_RequestJSON struct {
	AccessToken       string `json:"access_token"`
	BatteryCharge     string `json:"battery_charge"`
	IsShutterOpen     bool   `json:"is_shutter_open"`
	LastCommandID     int    `json:"last_command_id"`
	LastCommandAnswer int    `json:"last_command_answer"`
}

type Device_UpdateJSON struct {
	UUID        string `json:"uuid"`
	AccessToken string `json:"access_token"`
	Location    string `json:"location"`
}

type Device_StateJSON struct {
	IsShutterOpen     bool `json:"is_shutter_open"`
	LastCommandID     int  `json:"last_command_id"`
	LastCommandStatus int  `json:"last_command_status"`
}

type Device_CommandJSON struct {
	UniqueID    int    `json:"unique_id"`
	Data        string `json:"data"`
	RequesterID int    `json:"requester_id"`
}

type Server_Device_AnswerJSON struct {
	CommandID   int    `json:"command_id"`
	CommandData string `json:"command_data"`
}

type Device_Admin_CommandJSON struct {
	Command *Device_CommandJSON `json:"command"`
}

// Add a device
func Model_Device_CreateObject(device_to_add *Device_CreateJSON) (int, *Device) {

	db := db_Database()
	defer db.Close()

	// Check if there is such a device (by uuid)
	var device Device
	db.Where("uuid = ?", device_to_add.UUID).First(&device)
	if device.ID != 0 {
		return DB_ANSWER_OBJECT_EXISTS, nil
	}

	device = Device{
		UUID:        device_to_add.UUID,
		AccessToken: device_to_add.AccessToken,
		Location:    device_to_add.Location,
	}

	db.Save(&device)
	return DB_ANSWER_SUCCESS, &device
}

// Get a list of all devices
func Model_Device_GetList() []Device_ReadJSON {

	db := db_Database()
	defer db.Close()

	var devices []Device
	db.Find(&devices)
	devices_list := make([]Device_ReadJSON, 0)

	// If the list is empty, you do not need to take data
	if len(devices) <= 0 {
		return devices_list
	}

	for i := range devices {

		current_device := Device_ReadJSON{
			ID:            devices[i].ID,
			CreatedAt:     devices[i].CreatedAt,
			UUID:          devices[i].UUID,
			AccessToken:   devices[i].AccessToken,
			Location:      devices[i].Location,
			State:         devices[i].State,
			BatteryCharge: devices[i].BatteryCharge,
			LastSync:      devices[i].LastSync,
			Command:       devices[i].Command,
		}
		devices_list = append(devices_list, current_device)
	}

	return devices_list
}

// Get a device by UUID
func Model_Device_GetObject_byUUID(device_uuid string) (int, *Device_ReadJSON) {

	db := db_Database()
	defer db.Close()

	device := new(Device)
	db.Where("uuid = ?", device_uuid).First(&device)
	if device.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND, nil
	}

	device_read := Device_ReadJSON{
		ID:            device.ID,
		CreatedAt:     device.CreatedAt,
		UUID:          device.UUID,
		BatteryCharge: device.BatteryCharge,
		Location:      device.Location,
		State:         device.State,
		LastSync:      device.LastSync,
		Command:       device.Command,
	}

	return DB_ANSWER_SUCCESS, &device_read
}

// Update device data
func Model_Device_UpdateObject(update_json *Device_UpdateJSON) (int, *Device) {

	db := db_Database()
	defer db.Close()

	var device Device
	db.Where("uuid = ?", update_json.UUID).First(&device)
	if device.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND, nil
	}

	//UUID
	if update_json.UUID != "" {
		device.UUID = update_json.UUID
	}

	//AccessToken
	if update_json.AccessToken != "" {
		device.AccessToken = update_json.AccessToken
	}

	//Location
	if update_json.Location != "" {
		device.Location = update_json.Location
	}

	db.Save(&device)
	return DB_ANSWER_SUCCESS, &device
}

// Delete a device by UUID
func Model_Device_DeleteObject_byUUID(device_uuid string) int {

	db := db_Database()
	defer db.Close()

	var device Device
	db.Where("uuid = ?", device_uuid).First(&device)
	if device.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND
	}

	db.Unscoped().Delete(&device)
	return DB_ANSWER_SUCCESS
}

// Delete all devices
func Model_Device_DeleteObject_All() int {

	db := db_Database()
	defer db.Close()

	db.Unscoped().Delete(Device{})
	return DB_ANSWER_SUCCESS
}

// Get the device status
func Model_Device_State(uuid string, device_data *Device_Sync_RequestJSON) (int, *Server_Device_AnswerJSON) {

	db := db_Database()
	defer db.Close()

	var device Device
	var server_answer Server_Device_AnswerJSON

	db.Where("uuid = ?", uuid).First(&device)
	if device.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND, nil
	}

	if device.AccessToken != device_data.AccessToken {
		return DB_ANSWER_INVALID_CREDENTIALS, nil
	}

	state_raw_json, err_json := json.Marshal(device_data)
	if err_json != nil {
		db_LOG("Error parsing specialist. Error: " + err_json.Error())
		return DB_ANSWER_INVALID_JSON_CONVERSION, nil
	}

	state_str := string(state_raw_json)
	device.State = state_str
	device.BatteryCharge = device_data.BatteryCharge

	if device.Command != "" {

		var commandJSON Device_CommandJSON
		json.Unmarshal([]byte(device.Command), &commandJSON)

		server_answer = Server_Device_AnswerJSON{
			CommandID:   commandJSON.UniqueID,
			CommandData: commandJSON.Data,
		}
	}

	db.Save(&device)
	return DB_ANSWER_SUCCESS, &server_answer
}

// Receiving a command to the device
func Model_Device_Command(uuid string, command_data *Device_CommandJSON) int {

	db := db_Database()
	defer db.Close()

	var device Device
	var command_str string

	db.Where("uuid = ?", uuid).First(&device)
	if device.ID == 0 {
		return DB_ANSWER_OBJECT_NOT_FOUND
	}

	if (command_data.Data != "") && (command_data.UniqueID != 0) {
		command_raw_json, err_json := json.Marshal(command_data)
		if err_json != nil {
			db_LOG("Error parsing specialist. Error: " + err_json.Error())
			return DB_ANSWER_INVALID_JSON_CONVERSION
		}
		command_str = string(command_raw_json)
		device.Command = command_str

		db.Save(&device)
	} else {
		return DB_ANSWER_INVALID_COMMAND
	}
	return DB_ANSWER_SUCCESS
}
