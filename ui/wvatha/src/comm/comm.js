const net = require('net');
const EventEmitter = require('events').EventEmitter;
const os = require('os')
const filepath = require('fs')

function initBuildPath(platform) {
    if (os.platform === 'win32') {
        this.usePipe = true
        var pipe_name = "mypipe";
        var pipe_name = "\\\\.\\pipe\\" + PIPE_NAME;
        return pipe_name;
    } else {
        this.usePipe = false;
        return '/tmp/unixSocket';
    }
}

function CreateCommFile(path) {
    
}
/*
'aix', 'darwin', 'freebsd', 'linux', 'openbsd', 'sunos', and 'win32'

*/

// On IPC comm. windows will use named pipes and other platforms
// will use unix domain sockets.
class SocketEventsEmitter extends EventEmitter {
    constructor() {
        super({captureRejections: true});
        this.init();
    }

    // Create a socket or a named pipe for ipc
    init() {
        var path = initBuildPath(os.platform())
        var server = net.createServer().listen(path, () => {
            console.log(`now listening ${path}`);
        });

        server.on("error", (error) => {
            if (error.code === 'EADDRINUSE') {
                console.log('Caught the error');
                filepath.unlink(path, () => {
                    console.log('Trying to unlink');
                    this.init();
                });
            }

        });

        server.on("connection", (socket) => {
            socket.on("data", (buffer) => {
                console.log(buffer);
            });
        });
    }

    // Retry count will be 3 and the timeout is going to be 2 sec each.
    // If conn won't return ACK for...

}


const socketEventEmitter = new SocketEventsEmitter();

socketEventEmitter.emit('onConnect', "hello")

//https://stackoverflow.com/questions/37235021/how-do-i-use-event-handlers-between-javascript-module-files-with-node-js

module.exports = socketEventEmitter;
