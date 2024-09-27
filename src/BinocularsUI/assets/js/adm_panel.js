$(document).ready(function(){

    sessionStorage.setItem("is_gen", "0");
    
    // =========================================================
    // 
    // 						Main functions
    // 
    // ==========================================================
    
    // Hiding elements
    $("#box-message").hide();
    $("#devices-table").hide();
    $("#button-close-devices-list").hide();


    // Clicking on the "Close" button in the error window
    $("#box-message-close").click(function(){
        $("#box-message").hide();
        return false;
    });


    // =========================================================
    // 						  Devices
    // ==========================================================

    // Get a List of Devices button
    $("#button-get-devices-list").click(function(){
        let devicesGetListRequest = ajax_GET(CONFIG_APP_URL_BASE+"api/devices/", {}, {});
        handler_getRequest("get_devices_list", devicesGetListRequest);
        return false;
    });

    // Remove Device List button
    $("#button-close-devices-list").click(function(){
        $("#devices-table").hide();
        $("#button-close-devices-list").hide();
        return false;
    });

    // The Add New Device button
    $("#button-add-device").click(function(){

        let is_gen = sessionStorage.getItem("is_gen");

        if(is_gen == "0"){
            printMessage("error","Сгенирируйте UUID и токен!");
            return false;
        }

        let uuid = $("#input-uuid-device-add").val();
        let access_token = $("#input-accesstoken-device-add").val();
        let location = $("#input-location-device-add").val();

        if(uuid == "" || access_token == "" || location == ""){
            printMessage("error","Пустые поля!");
            return false;
        }

        let credentials = {
            "uuid": uuid,
            "access_token": access_token,
            "location": location,
        };

        let addDeviceRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/devices/","POST", credentials, {});
        handler_postRequest("add_device", addDeviceRequest);
        return false;
    });

    // Generate button (uuid and access_token) 
    $("#button-gen-token").click(function(){

        let uuid = pass_gen(5);
        let access_token = pass_gen(10);
        $("#input-uuid-device-add").val(uuid);
        $("#input-accesstoken-device-add").val(access_token);
        sessionStorage.setItem("is_gen", "1");
        return false;
    });

    // Edit Device button
    $("#button-change-device").click(function(){

        let uuid = $("#input-uuid-device-add").val();
        let access_token = $("#input-accesstoken-device-add").val();
        let location = $("#input-location-device-add").val();

        let credentials = {
            "uuid": uuid,
            "access_token": access_token,
            "location": location,
        };

        let updateDeviceRequest = ajax_PUT(CONFIG_APP_URL_BASE+"api/devices/", credentials, {});
        handler_updateRequest("update_device", updateDeviceRequest);
        return false;
    });

    // The Delete Device button
    $("#button-delete-device").click(function(){

        let uuid = $("#input-uuid-device-del").val();
        if(uuid == ""){
            printMessage("error","Поле с uuid пустое");
            return false;
        }

        let deleteDeviceRequest = ajax_DELETE(CONFIG_APP_URL_BASE+"api/devices/" + uuid, {}, {});
        handler_deleteRequest("del_device", deleteDeviceRequest);
        return false;
    });

    // The Delete All Devices button
    $("#button-delete-all-devices").click(function(){
        if(confirm("Вы действительно хотите удалить все устройства?")) {
            let deleteDevicesAllRequest = ajax_DELETE(CONFIG_APP_URL_BASE+"api/devices/0", {}, {});
            handler_deleteRequest("del_devices_all", deleteDevicesAllRequest);
        }
        return false;
    });

    // The Send Command button
    $("#button-command-send").click(function(){

        let uuid = $("#input-uuid-device-send-command").val();
        let command = $("#input-command-device-send-command").val();

        if(uuid == "" || command == ""){
            printMessage("error","Пустые поля");
            return false;
        }

        let unique_id = randomInteger(1,CONFIG_MAX_UNIQUE_ID);

        let credentials = {
            "unique_id": unique_id,
            "data": command,
        };

        let sendCommandDeviceRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/devices/" + uuid + "/command","POST", credentials, {});
        handler_postRequest("send_command_device", sendCommandDeviceRequest);
        return false;
    });

    
    // The button to open the shuttle for 5 seconds
    $("#button-command-send-open-shatter").click(function(){

        let uuid = $("#input-uuid-device-send-command").val();
        let command = "C01S00005";

        if(uuid == ""){
            printMessage("error","Пустое поле с uuid");
            return false;
        }

        let unique_id = randomInteger(1,CONFIG_MAX_UNIQUE_ID);

        let credentials = {
            "unique_id": unique_id,
            "data": command,
        };

        let sendCommandDeviceRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/devices/" + uuid + "/command","POST", credentials, {});
        handler_postRequest("send_command_device", sendCommandDeviceRequest);
        return false;
    });

    // Close the shuttle button
    $("#button-command-send-close-shatter").click(function(){

        let uuid = $("#input-uuid-device-send-command").val();
        let command = "C02";

        if(uuid == ""){
            printMessage("error","Пустое поле с uuid");
            return false;
        }

        let unique_id = randomInteger(1,CONFIG_MAX_UNIQUE_ID);

        let credentials = {
            "unique_id": unique_id,
            "data": command,
        };

        let sendCommandDeviceRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/devices/" + uuid + "/command","POST", credentials, {});
        handler_postRequest("send_command_device", sendCommandDeviceRequest);
        return false;
    });

    // The Emulator button
    $("#button-to-emulator").click(function(){
        window.location.replace("/device-emulator");
        return false;
    });
});

// -----------------------------------
//                 Misc
// -----------------------------------
function randomInteger(min, max) {
    let rand = min + Math.random() * (max + 1 - min);
    return Math.floor(rand);
}

// -----------------------------------
//              Views
// -----------------------------------


// -----------------------------------
//              Requests
// -----------------------------------
function showTableDevicesList(data){

    let table_header = ` <table class="devices-list-table">
                        <tbody>
                        <tr>
                            <th>Номер</th>
                            <th>ID в базе</th>
                            <th>Дата создания</th>
                            <th>UUID</th>
                            <th>Местоположение</th>
                            <th>Статус</th>
                            <th>Дата последней синхронизации</th>
                            <th>Команда</th>
                        </tr>`;

    let table_footer = `</tbody>
                        </table>`;

    let devices_list = "";
    let x = 0;
    const element = data;
    for(i=0; i < element.length; i++){
        x = i+1;

        devices_list += `<tr>
                            <td>${x}</td>
                            <td>${element[i].id}</td>
                            <td>${element[i].created_at}</td>
                            <td>${element[i].uuid}</td>
                            <td>${element[i].location}</td>
                            <td>${element[i].state}</td>
                            <td>${element[i].last_sync}</td>
                            <td>${element[i].command}</td>
                        </tr>`;
    }
    $('#devices-table').html(table_header + devices_list + table_footer);
    $('#devices-table').show();
    $('#button-close-devices-list').show();
}

// -----------------------------------
//              Handlers
// -----------------------------------

// GET
function handler_getRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            case 200:
                switch(request_type){
                    case "get_devices_list":
                        if(request.responseJSON.data.length != 0){
                            showTableDevicesList(request.responseJSON.data);
                        } else{
                            printMessage("error","Ни одно устройство не зарегистрированно");
                        }
                        
                        break;
                }     
                break;

            case 401:
                printMessage("error","Администратор не авторизован");
                console_RequestError("Invalid auth!",request);
                break;

            case 404:
                printMessage("error","Администратор не найден");
                console_RequestError("User not found!",request);
                break;
            
            //В ином случае
            default:
                printMessage("error","Неизвестная ошибка");
                console_RequestError("Error!", request);
                break;
        }
    });
}

// POST
function handler_postRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            case 200:
                switch(request_type){
                    case "add_device":
                        sessionStorage.setItem("is_gen", "0");
                        printMessage("success","Новое устройство успешно добавлено");
                        break;

                    case "send_command_device":
                        printMessage("success","Команда успешно отправлена");
                        break;
                }     
                break;

            case 401:
                printMessage("error","Администратор не авторизован");
                console_RequestError("Invalid auth!",request);
                break;

            case 404:
                switch(request_type){
                    case "send_command_device":
                        printMessage("error","Устройство не найдено");
                        break;
                }
                console_RequestError("Device not found!",request);
                break;

            case 400:
                switch(request.responseJSON.status.code){
                    case 1:
                        switch(request_type){
                            case "add_device":
                                printMessage("error","Устройство с таким uuid уже существует");
                                console_RequestError("Device exists!",request);
                                break;
                        }
                        
                        break;

                    case 3:
                        printMessage("error","Неверные данные в запросе");
                        console_RequestError("Incorrect data in the query!",request);
                        break;

                    case 4:
                        errorMessageText("Пустые поля!");
                        console_RequestError("Empty fields!",request);
                        break;

                    default:
                        errorMessageText("Неизвестная ошибка!");
                        console_RequestError("Error!", request);
                        break;
                }
                break;
            
            //В ином случае
            default:
                errorMessageText("Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}

// PUT
function handler_updateRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            case 200:
                switch(request_type){
                    case "update_device":
                        printMessage("success","Данные устройства успешно обновлены");
                        break;
                }     
                break;

            case 401:
                errorMessageText("Администратор не авторизован!");
                console_RequestError("Invalid auth!",request);
                break;

            case 404:
                switch(request_type){
                    case "update_device":
                        printMessage("error","Устройство с таким uuid не найдено");
                        console_RequestError("Device not found!",request);
                        break;
                }
                break;
            
            case 400:
                switch(request.responseJSON.status.code){
                    case 3:
                        printMessage("error","Неверные данные в запросе");
                        console_RequestError("Incorrect data in the query!",request);
                        break;

                    case 4:
                        errorMessageText("Пустые поля!");
                        console_RequestError("Empty fields!",request);
                        break;

                    default:
                        errorMessageText("Неизвестная ошибка!");
                        console_RequestError("Error!", request);
                        break;
                }
                break;

            //В ином случае
            default:
                errorMessageText("Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}

// DELETE
function handler_deleteRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            case 200:
                switch(request_type){
                    case "del_device":
                        printMessage("success","Устройство успешно удалено");
                        break;
                    
                    case "del_devices_all":
                        printMessage("success","Все устройства успешно удалены");
                        break;
                }     
                break;

            case 401:
                errorMessageText("Администратор не авторизован!");
                console_RequestError("Invalid auth!",request);
                break;

            case 404:
                switch(request_type){
                    case "del_device":
                        printMessage("error","Устройство с таким uuid не найдено");
                        console_RequestError("Device not found!",request);
                        break;
                }
                break;

            case 400:
                switch(request.responseJSON.status.code){
                    case 3:
                        printMessage("error","Неверные данные в запросе");
                        console_RequestError("Incorrect data in the query!",request);
                        break;

                    default:
                        errorMessageText("Неизвестная ошибка!");
                        console_RequestError("Error!", request);
                        break;
                }
                break;
            
            //В ином случае
            default:
                errorMessageText("Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}