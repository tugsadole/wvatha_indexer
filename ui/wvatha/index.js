// index.js
const socketEventEmitter = require('./src/comm/comm')

const { app, BrowserWindow } = require("electron");
const path = require("path");
app.on("ready", () => {
  const mainWindow = new BrowserWindow(
    {webPreferences: {
      nodeIntegration: true,
      contextIsolation: false,
    }}
  );
  mainWindow.loadFile(path.join(__dirname, "public/index.html"));
  mainWindow.webContents.openDevTools();
  console.log("TUGSAD")
  console.log(socketEventEmitter.usePipe)

  socketEventEmitter.on('onConnect', function (data) {
    console.log(data);
  });
  socketEventEmitter.emit('onConnect', "hello");

});



