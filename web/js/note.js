document.getElementById('note').value = document.getElementById('note').textContent;

document.getElementById('save').onclick = function() {
    const data = {
        text: document.getElementById('note').value,
    };
    const successCallback = function () {
        return false;
    };
    sendRequest(window.location.pathname+'/save', data, successCallback.bind(this, window));
}