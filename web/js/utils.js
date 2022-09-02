function sendRequest(url, data, successCallback, errorCallback) {
    const xmlHttp = new XMLHttpRequest();
    xmlHttp.open("POST", url);
    xmlHttp.setRequestHeader("Content-Type", "application/json");
    xmlHttp.send(JSON.stringify(data));

    xmlHttp.onload = function(event) {
        if (xmlHttp.status >= 200 && xmlHttp.status < 300) {
            successCallback(xmlHttp.response);
        } else {
            errorCallback();
        }
    };
}