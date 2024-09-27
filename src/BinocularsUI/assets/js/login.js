$(document).ready(function () {

    $("#box-error").hide();

    // Click on the login button
    $("#button-user-login").click(function () {

        // We collect information from the email and password fields
        let login = $("#input-user-login").val();
        let password = $("#input-user-password").val();

        if(login == "" || password == "") {
            $("#box-error").show();
            $("#box-error p").text("Все поля должны быть заполнены!");
            return false;
        };

        let credentials = {
            "login": login,
            "password": password,
        };
        loginRequest = ajax_JSON(CONFIG_APP_URL_BASE + "login", "POST", credentials, {});
        handler_sendLoginRequest(loginRequest);
        return false;
    });
});



// -----------------------------------
// 
//          AJAX HANDLERS
// 
// -----------------------------------

function handler_sendLoginRequest(request) {
    console.log(request);
    request.always(function () {
        switch (request.status) {
            case 200:
                $("#box-error").hide();
                window.location.replace("/adm-panel");
                break;

            case 401:
                $("#box-error").show();
                $("#box-error p").text("Неверный логин или пароль!");
                console_RequestError("Invalid auth!", request);
                break;

            case 404:
                $("#box-error").show();
                $("#box-error p").text("Пользователя с таким логином не существует!");
                console_RequestError("Invalid auth!", request);
                break;

            default:
                $("#box-error").show();
                $("#box-error p").text("Неизвестная ошибка!");
                console_RequestError("Error!", request);
                break;
        }
    });
}
