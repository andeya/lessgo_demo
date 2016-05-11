var wsUri = "ws://localhost:8080/websocket";
var ws = new WebSocket(wsUri);
ws.onopen = function() {
    console.log("connected to " + wsUri);
};
ws.onclose = function(e) {
    console.log("connection closed (" + wsUri + " : " + e.code + "," + e.reason + ")");
}
ws.onerror = function(e) {
    for (var p in e) {
        console.log(p + "=" + e[p]);
    }
};
ws.onmessage = function(m) {
    console.log("receive:", m.data);
};
window.onload = function() {
    var now = new Date();
    ws.send(JSON.stringify({
        "testsend": now
    }));
    console.log("send:", now);
}
window.onbeforeunload = function() {
    ws.close();
    console.log("关闭连接");
    return
}
