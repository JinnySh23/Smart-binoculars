$(document).ready(function() {

    is_connect = false;
    $("#box-message").hide();
    
    sessionStorage.setItem("last_command_id",0);

    // =========================================================
    // 
    // 						 Main functions
    // 
    // ==========================================================
    
    // Clicking on the "Close" button in the error window
    $("#box-message-close").click(function(){
        $("#box-message").hide();
        return false;
    });

    // Connect/Disconnect
    $("#button-device-connect").click(function(){
        if(!is_connect){
            if((sessionStorage.getItem("uuid") == null) || (sessionStorage.getItem("access_token") == null)){
                printMessage("error","Введите uuid и AccesToken в поля ниже и нажмите Сохранить");
                return false;
            }

            let uuid = sessionStorage.getItem("uuid");
            let access_token = sessionStorage.getItem("access_token");

            timerDevice("start");
            $("#a-device-status").addClass("connect");
            $("#a-device-status").text("Connected");
            $("#span-device-uuid").text(uuid);
            $("#span-device-token").text(access_token);
            is_connect = true;
        } else{
            timerDevice("stop");
            $("#a-device-status").removeClass("connect");
            $("#a-device-status").text("No connected");
            $("#button-device-connect").text("Connect");
            $("#span-device-uuid").text("");
            $("#span-device-token").text("");
            is_connect = false;
        }
        return false;
    });

    // The Save button
    $("#button-device-save").click(function(){

        let uuid = $("#input-uuid-device").val();
        let access_token = $("#input-token-device").val();

        if(uuid == "" || access_token == ""){
            printMessage("error","Введите uuid и AccesToken в поля");
            return false;
        }

        sessionStorage.setItem("uuid", uuid);
        sessionStorage.setItem("access_token", access_token);
        printMessage("success","Устройство сохранено и готово к подключению");
        return false;
    });

});

// -----------------------------------
//                 Misc
// -----------------------------------
// Request for device status
function getStatusDevice(){
    if((sessionStorage.getItem("uuid") == null) || (sessionStorage.getItem("access_token") == null)){
        printMessage("error","uuid и AccesToken отсутствуют в сессии");
        timerDevice("stop");
        $("#a-device-status").removeClass("connect");
        $("#a-device-status").text("No connected");
        $("#button-device-connect").text("Connect");
        $("#span-device-uuid").text("");
        $("#span-device-token").text("");
        is_connect = false;
        return false;
    }

    let is_shutter_open = false;
    if(sessionStorage.getItem("is_shutter_open") != null){
        is_shutter_open = sessionStorage.getItem("is_shutter_open");
    }

    let last_command_id = 0;
    if(sessionStorage.getItem("last_command_id") != null){
        last_command_id = sessionStorage.getItem("last_command_id");
    }

    let last_command_answer = 0;
    if(sessionStorage.getItem("last_command_answer") != null){
        last_command_answer = sessionStorage.getItem("last_command_answer");
    }

    let uuid = sessionStorage.getItem("uuid");
    let access_token = sessionStorage.getItem("access_token");

    let credentials = {
        "access_token": access_token,
        "is_shutter_open": Boolean(is_shutter_open),
        "last_command_id": Number(last_command_id),
        "last_command_answer": last_command_answer,
    };

    let statusDeviceRequest = ajax_JSON(CONFIG_APP_URL_BASE+"api/devices/" + uuid + "/state","POST", credentials, {});
    handler_postRequest("status_device", statusDeviceRequest);
}

// Universal device Timer
function timerDevice(command){
    let timerDeviceId;
    switch(command){
        case "start":
            timerDeviceId = setInterval(getStatusDevice, 1000);
            sessionStorage.setItem("timerDeviceId",timerDeviceId);
            break;

        case "stop":
            timerDeviceId = sessionStorage.getItem("timerDeviceId");
            clearInterval(timerDeviceId);
            break;
    
        default:
            timerDeviceId = sessionStorage.getItem("timerDeviceId");
            clearInterval(timerDeviceId);
            break;
    }
}

// The command to open the shuttle
function openShutter(){
    sessionStorage.setItem("is_shutter_open",true);
    $("#span-device-is-shutter").text("OPEN");
}

// The command to close the shuttle
function closeShutter(){
    $("#span-device-is-shutter").text("CLOSED");
    sessionStorage.setItem("is_shutter_open",false);
}

// -----------------------------------
//              Views
// -----------------------------------


// -----------------------------------
//            Requests
// -----------------------------------


// -----------------------------------
//            Handlers
// -----------------------------------
// POST
function handler_postRequest(request_type, request){
    request.always(function(){
    
        switch(request.status){
            case 200:
                switch(request_type){
                    case "status_device":
                        $("#a-device-server-answer").addClass("connect");
                        $("#a-device-server-answer").text("200 (OK)");
                        let element = request.responseJSON.data;
                        let last_command_id = sessionStorage.getItem("last_command_id");
                        if(element.command_data != ""){
                            if(last_command_id != element.command_id){
                                command_data_C = element.command_data.split('S')[0];
                                command_data_S = element.command_data.split('M')[0].split('S')[1];
                                command_data_M = element.command_data.split('M')[1];
                                switch(command_data_C){
                                    case "C01":
                                        sessionStorage.setItem("last_command_id",element.command_id);
                                        openShutter();
                                        setTimeout(closeShutter, Number(command_data_S) * 1000);
                                        break;

                                    case "C02":
                                        closeShutter();
                                        break;
                                
                                    default:
                                        break;
                                }
                            }
                        }
                        break;
                }     
                break;

            case 401:
                $("#a-device-server-answer").removeClass("connect");
                $("#a-device-server-answer").text("401 (ERROR)");
                printMessage("error","Администратор не авторизован");
                console_RequestError("Invalid auth!",request);
                timerDevice("stop");
                break;

            case 404:
                switch(request_type){
                    case "status_device":
                        $("#a-device-server-answer").removeClass("connect");
                        $("#a-device-server-answer").text("404 (NOT FOUND)");
                        printMessage("error","Устройство не найдено");
                        timerDevice("stop");
                        break;
                }
                console_RequestError("Device not found!",request);
                timerDevice("stop");
                break;

            case 400:
                switch(request.responseJSON.status.code){
                    case 3:
                        $("#a-device-server-answer").removeClass("connect");
                        $("#a-device-server-answer").text("400 - 3 (ERROR)");
                        printMessage("error","Неверные данные в запросе");
                        console_RequestError("Incorrect data in the query!",request);
                        timerDevice("stop");
                        break;

                    case 4:
                        $("#a-device-server-answer").removeClass("connect");
                        $("#a-device-server-answer").text("400 - 4 (ERROR)");
                        printMessage("error","Пустые поля!");
                        console_RequestError("Empty fields!",request);
                        timerDevice("stop");
                        break;

                    case 6:
                        $("#a-device-server-answer").removeClass("connect");
                        $("#a-device-server-answer").text("400 - 6 (ERROR)");
                        printMessage("error","Токены не совпадают!");
                        console_RequestError("Invalid AccessToken!",request);
                        timerDevice("stop");
                        break;

                    case 503:
                        $("#a-device-server-answer").removeClass("connect");
                        $("#a-device-server-answer").text("400 - 503 (ERROR)");
                        printMessage("error","Конвертация в JSON  - ошибка поля!");
                        console_RequestError("Invalid JSON conversion!",request);
                        timerDevice("stop");
                        break;

                    default:
                        $("#a-device-server-answer").removeClass("connect");
                        $("#a-device-server-answer").text("400 - ? (ERROR)");
                        printMessage("error","Неизвестная ошибка!");
                        console_RequestError("Error!", request);
                        timerDevice("stop");
                        break;
                }
                break;
            
            default:
                $("#a-device-server-answer").removeClass("connect");
                $("#a-device-server-answer").text("BAD ERROR");
                printMessage("error","Неизвестная ошибка!");
                console_RequestError("Error!", request);
                timerDevice("stop");
                break;
        }
    });
}