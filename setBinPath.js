const fs = require("fs");
const os = require("os");

const isWindows = os.platform() === "win32";
const binPath = isWindows ? "./bin/sygnus.exe" : "./bin/sygnus";

const packageJsonPath = "./package.json";
const packageJson = require(packageJsonPath);

packageJson.bin = {
  sygnus: binPath,
};

fs.writeFileSync(packageJsonPath, JSON.stringify(packageJson, null, 2));
