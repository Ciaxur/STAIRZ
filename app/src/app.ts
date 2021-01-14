import { app, BrowserWindow } from 'electron';

const createWindow = () => {
  const browserOptions = process.env.NODE_ENV !== 'development'
    ? { fullscreen: true }
    : {
      width: 110 * 10,
      height: 80 * 10,
    };
  
  const win = new BrowserWindow({
    ...browserOptions,
    webPreferences: {
      nodeIntegration: true,
    },
  });

  win.loadFile('dist/renderer/index.html');
};

app.whenReady()
  .then(createWindow);

app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('activate', () => {
  if (BrowserWindow.getAllWindows().length === 0) {
    createWindow();
  }
});