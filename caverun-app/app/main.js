
const { app, BrowserWindow, Menu, MenuItem, dialog, } = require('electron');
const path = require('path');
const url = require('url');
const child_process = require('child_process');

let debugFlag = false;
for (let i in process.argv) {
	if (process.argv[i] == "-debug") {
		debugFlag = true;
	}
}

/**
 * Generate a unique ID for each window
 * @return {string}
 */
let windowIDGenerator = function () {
    let S4 = function() {
        return (((Date.now()))|0).toString(16).substring(1);
    };
    return (S4()+S4()+"-"+S4()+"-"+S4()+"-"+S4()+"-"+S4()+S4()+S4());
};


// server window, hidden, used to serve http data requests
let serverWin;

let serverProc;

let settings;
// sets up local http server, calls doWindowSetup when done
function doServerSetup() {

//	serverProc.spawn('ps', ['ax']);
app.getPath('userData');

	// special case for unpacked asar stuff
	let unpackedDirName = __dirname+"";
	if (unpackedDirName.match(/app[.]asar$/)) {
		unpackedDirName += ".unpacked";
	}
	let serverProcPath = path.join(unpackedDirName, 'caverun-srv'+(process.platform.match(/^win/)?'.exe':''));

	// looks like this exits automatically with the app, should be fine for now.
	serverProc = child_process.spawn(serverProcPath, ["-publicdir", path.join(__dirname, "public"), "-listen", "127.0.0.1:6753"]);
	serverProc.stdout.on('data', (data) => {
	  console.log(`caverun-srv (stdout): ${data}`);
	});
	serverProc.stderr.on('data', (data) => {
	  console.log(`caverun-srv (stderr): ${data}`);
	});
	serverProc.on('close', (code) => {
	  if (code !== 0) {
	    console.log(`caverun-srv process exited with code ${code}`);
	  }
	});

	doWindowSetup();

	// give it any needed settings
	// https://electron.atom.io/docs/api/exapp/#appgetpathname
	settings = {
		'userDataPath' : app.getPath('userData'),
	};

}

// global reference of all open windows
let wins = [];

function openNewBrowserWindow() {

	// Create the browser window.
	let win = new BrowserWindow({ width: 1200, height: 800 });
    win.webContents.uid = windowIDGenerator();

	// Open the DevTools in debug mode
	if (debugFlag) {
		win.webContents.openDevTools();
	}

	// Emitted when the window is closed.
	win.on('close', function () {
		// Dereference the window object, usually you would store windows
		// in an array if your app supports multi windows, this is the time
		// when you should delete the corresponding element.
        win = null;
        wins.pop(win);
        delete wins[win];
	});


	win.loadURL("http://127.0.0.1:6753/home.html");

	wins.push(win);

}



function doWindowSetup() {

	let appMenu = Menu.getApplicationMenu();
	
	const menuTemplate = [
		{
			label: 'Electron',
			submenu: [
				{
					label: 'About ...',
					click: () => {
						console.log('About Clicked');
					}
				}, {
					label: 'Quit',
					role: "quit",
					// click: () => {
					//     app.quit();
					// }
				},
			],
		},
		{
			label: 'File',
			submenu: [
				{
					label: 'New Window',
					click: () => {
						require('child_process').exec(openNewBrowserWindow());
					},
				}
			],
		},
 		{
			label: "Edit",
			submenu: [
				{ label: "Undo", accelerator: "CmdOrCtrl+Z", selector: "undo:" },
				{ label: "Redo", accelerator: "Shift+CmdOrCtrl+Z", selector: "redo:" },
				{ type: "separator" },
				{ label: "Cut", accelerator: "CmdOrCtrl+X", selector: "cut:" },
				{ label: "Copy", accelerator: "CmdOrCtrl+C", selector: "copy:" },
				{ label: "Paste", accelerator: "CmdOrCtrl+V", selector: "paste:" },
				{ label: "Select All", accelerator: "CmdOrCtrl+A", selector: "selectAll:" }
			],
		},
	];
	const menu = Menu.buildFromTemplate(menuTemplate);
	Menu.setApplicationMenu(menu);

	if (wins.length == 0) {
		openNewBrowserWindow();
		return;
	}


}


// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', doServerSetup);
// Quit when all windows are closed.
app.on('window-all-closed', () => {
	// On macOS it is common for applications and their menu bar
	// to stay active until the user quits explicitly with Cmd + Q
	//if (process.platform !== 'darwin') {

	if (serverProc) {
		serverProc.kill();
		serverProc = null;
	}

		app.quit();
	//}
});

app.on('before-quit', () => {
	if (serverProc) {
		serverProc.kill();
		serverProc = null;
	}
});

app.on('activate', () => {
	// On macOS it's common to re-create a window in the app when the
	// dock icon is clicked and there are no other windows open.
	if (wins.length == 0) {
		doWindowSetup();
	}
});
// In this file you can include the rest of your app's specific main process
// code. You can also put them in separate files and require them here.

//Only allow one instance of the app

let myWindow = null;

let shouldQuit = app.makeSingleInstance(function(commandLine, workingDirectory) {
  // Someone tried to run a second instance, we should focus our window.
  if (myWindow) {
    if (myWindow.isMinimized()) myWindow.restore();
    myWindow.focus();
  }
});

if (shouldQuit) {
  app.quit();
  return;
}


// In main process.
const {ipcMain} = require('electron')


/**
 * Opens a new window
 */
ipcMain.on('open-new-window', function(){
     openNewBrowserWindow();
});

