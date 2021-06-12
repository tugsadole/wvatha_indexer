const net = require('net');
const EventEmitter = require('events');

class SocketEventsEmitter extends EventEmitter {}
const socketEventEmitter = new SocketEventsEmitter();

//https://stackoverflow.com/questions/37235021/how-do-i-use-event-handlers-between-javascript-module-files-with-node-js

module.exports = socketEventEmitter;

func

constructor({
    onOpen = () => {},
    onClose = () => {},
    onDisconnected = () => {},
    onReconnected = () => {},
})