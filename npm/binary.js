#!/usr/bin/env node

const { Binary } = require("binary-install");
const os = require('os');
const action = process.argv[2] || "run";

if(!["run", "install"].includes(action)) {
  throw new Error(`Invalid action: ${action} provided. Must be 'run' or 'install'`)
}

function getPlatform() {
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

const platform = getPlatform();
const { name, version, repository } = require("../package.json");

const url = `${repository.url}/releases/download/v${version}/${name}-${platform}.tar.gz`;

const binary = new Binary(name, url);

switch (action) {
  case "run": {
    binary.run();
    break;
  }
  case "install": {
    binary.install();
    break;
  }
}
