let raxx = require('child_process').spawn(__dirname+'/raxx')
const {app, BrowserWindow} = require('electron')
const path = require('path')
const url = require('url')

// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let win
//Keys
const {globalShortcut} = require('electron')
app.on('ready', () => {
  const ret = globalShortcut.register('CommandOrControl+X', () => {
    app.quit();
  })
  if (!ret) {
    console.log('registration failed')
  }
  //console.log(globalShortcut.isRegistered('CommandOrControl+X'))
})

app.on('will-quit', () => {
  globalShortcut.unregister('CommandOrControl+X')
  globalShortcut.unregisterAll()
})
//End keys

function createWindow () {
  // Create the browser window.
  win = new BrowserWindow({width: 220,height: 30,darkTheme: true,transparent: true,resizable: false,frame: false, x: 0,y: 708})
  win.setAlwaysOnTop(false);
  win.setMenuBarVisibility(false);
  win.setSkipTaskbar(true);
  win.setAlwaysOnTop(true);
  // and load the index.html of the app.
  win.loadURL('file://'+__dirname+'/index.html');
  // Open the DevTools.
  //win.webContents.openDevTools()

  // Emitted when the window is closed.
  win.on('closed', () => {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    raxx.kill('SIGINT');
    win = null
  })
}

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow)

// Quit when all windows are closed.
app.on('window-all-closed', () => {
  // On macOS it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', () => {
  // On macOS it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (win === null) {
    createWindow()
  }
})

// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.