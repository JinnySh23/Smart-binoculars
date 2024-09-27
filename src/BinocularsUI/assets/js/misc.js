// Page reload function
function reloadPage(){
	location.reload();
}

// Exiting the session
function sessionExit(){
    outSessionRequest = ajax_GET(CONFIG_APP_URL_BASE+"logout",{},{});
    handler_outSessionRequest(outSessionRequest);
}

function pass_gen(len) {
    chrs = 'abdehkmnpswxzABDEFGHKMNPQRSTWXZ123456789';
    let str = '';
    for(let i = 0; i < len; i++) {
        let pos = Math.floor(Math.random() * chrs.length);
        str += chrs.substring(pos,pos+1);
    }
    return str;
}