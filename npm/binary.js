const { Binary } = require("binary-install");
const os = require('os');

const getPlatform = () => {
  const arch = os.arch();

  if(arch !== "x64") {
    throw new Error("64 bit required");
  }
  switch (os.type()) {
    case "Windows_NT":
      return "win64";
    case "Linux":
      return "linux";
    case "Darwin":
      return "macos";
  }
}

const getBinary = () => {
  const platform = getPlatform();
  const { name, version } = require("../package.json");

  const url = `https://github.com/ansballard/${name}/releases/download/v${version}/${name}-${platform}.tar.gz`;

  return new Binary(name, url);
};

const run = () => {
  const binary = getBinary();
  binary.run();
};

const install = () => {
  const binary = getBinary();
  binary.install();
};

module.exports = {
  install,
  run
};
