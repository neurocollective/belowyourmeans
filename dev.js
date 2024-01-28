const { exec, spawn } = require('child_process');

const serverProcess = exec("make serve/local"); // { detached: true });
const uiProcess = exec("make serve/ui"); //, { detached: true });

serverProcess.stdout.on('data', data => console.log(`[server stdout] ${data}`));
serverProcess.stderr.on('data', data => console.log(`[server stderr] ${data}`));

uiProcess.stdout.on('data', data => console.log(`[ui stdout] ${data}`));
uiProcess.stderr.on('data', data => console.log(`[ui stderr] ${data}`));
