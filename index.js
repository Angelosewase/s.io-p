var socket = new WebSocket("ws://localhost:8080/ws");
function connectToSocketServer() {
    socket.onopen = function (event) {
        console.log("WebSocket connection opened:", event);
        sendMessage(); // Send message after the connection is open
    };
    socket.onmessage = function (evMessage) {
        console.log("Message from server:", evMessage.data);
    };
    socket.onerror = function (err) {
        console.log("WebSocket error:", err);
    };
    socket.onclose = function () {
        console.log("WebSocket connection closed");
    };
}
function sendMessage() {
    if (socket.readyState === WebSocket.OPEN) {
        socket.send("hello world");
    }
    else {
        console.log("WebSocket is not open. Current state:", socket.readyState);
    }
}
connectToSocketServer();
