//      RZHEVSKY ROBOTICS
//  Web-development Division
// 
//  core.js - Базоваые функции и конфигурация проекта
//

const CONFIG_APP_URL_BASE = "YOUR_URL";

const CONFIG_DATE_TIME_FORMAT_SHORT = new Intl.DateTimeFormat('ru', {
	year: 'numeric',
	month: 'numeric',
	day: 'numeric',
});

const CONFIG_DIALOG_REFRESH_INTERVAL = 5000; //ms

const CONFIG_MAX_UNIQUE_ID = 2147483648;

//Обработчик ошибки запроса
function console_RequestError(message, request){
    console.error("------\n"
    +"Error! Message: "+message +"\n"
    +"Status code: "+request.status +"\n"
    +"Answer: "+ request.responseText +"\n"
    +"------");
}
//Показать объект запроса
function console_RequestShowObject(object){
    console.log("%c ------",'color: green');
    console.log("%c Ok! Answer: ",'color: green');
    console.log(object);
    console.log("%c ------",'color: green');
}


//AJAX
function ajax_JSON(url,request_type, data,custom_headers){
	return $.ajax({
		url : url,
		type : request_type,
		headers: custom_headers,
		data:	JSON.stringify(data),
		contentType: 'application/json; charset=utf-8',
		dataType: 'json',
	});
}

function ajax_PUT(url, data,custom_headers){
	return $.ajax({
		url : url,
		type : "PUT",
		headers: custom_headers,
		data:	JSON.stringify(data),
		contentType: 'application/json; charset=utf-8',
		dataType: 'json',
	});
}

function ajax_GET(url,data,custom_headers){
	return $.ajax({
		url : url,
		type : "GET",
		headers: custom_headers,
		data:	data,
	});
}

function ajax_DELETE(url,data,custom_headers){
	return $.ajax({
		url : url,
		type : "DELETE",
		headers: custom_headers,
		data:	data,
	});
}

function ajax_SendFile(url,formData,custom_headers){
	console.log(url);
	return $.ajax({
		url : url,
		type : "POST",
		headers: custom_headers,
		timeout: 60000,
		contentType: false,
		processData: false,
		data: formData,
	});
}

// Getting GET parameters
function getUrlParameter(sParam) {
    let sPageURL = window.location.search.substring(1),
        sURLVariables = sPageURL.split('&'),
        sParameterName,
        i;

    for (i = 0; i < sURLVariables.length; i++) {
        sParameterName = sURLVariables[i].split('=');

        if (sParameterName[0] === sParam) {
            return sParameterName[1] === undefined ? true : decodeURIComponent(sParameterName[1]);
        }
    }
};

// Exiting the session
function sessionExit(){
    outSessionRequest = ajax_GET(CONFIG_APP_URL_BASE+"logout",{},{});
    handler_outSessionRequest(outSessionRequest);
}

// Page reload function
function reloadPage(){
	location.reload();
}

// Exiting the session
function handler_outSessionRequest(request){
    request.always(function(){
        switch(request.status){

            case 200:
                window.location.replace("/login");
                break;

            case 404:
                console_RequestError("Error!",request); 
                break;
            
            default:
                console_RequestError("Error!",request); 
                break;  
        }
    });
}

function printMessage(type,text){
	if(type == "success"){
		$("#box-message").addClass("success");
	} else{
		$("#box-message").removeClass("success");
	}
	$("#box-message").show();
	$("#box-message-text").text(text);
}