# Lebeben
_lebeben-go_

Lebeben is a single CLI to transpile, bundle, watch, and serve your javascript SPA.

## Usage

### **CLI Options**

- All option are prefixed with a single `-`
- boolean options must be passed `=1`
  - `watch=1 serve=1`, etc

|Name|Description|Type|Default|
|---|---|---|---|
|**help**|display the help message|boolean|false|
|**jsxFactory**|jsx render function to use for nodes|string|h|
|**jsxFragment**|jsx render function to use for fragments|string|Fragment|
|**minify**|minify all output|boolean|false|
|**nomodule**|generate the es2015 fallback bundle|boolean|false|
|**port**|the port to serve|number|5000|
|**public**|the directory to serve|string|public|
|**serve**|serve app locally|boolean|false|
|**watch**|one or many directories to watch for changes|boolean|false|

### **Examples**

```sh
# Build your app from src/index.js
lebeben-go src/index.js
```

```sh
# Build/Minify your React app from src/index.jsx
lebeben-go -minify=1 --jsxFactory="React.createElement" -jsxFragment="React.Fragment" src/index.jsx
```

```sh
# Watch/Serve a preact app from src/index.tsx
lebeben-go -watch=1 -serve=1 src/index.tsx
```

## Build across platforms in bash (handled by Github Actions)

```sh
env GOOS='darwin' GOARCH='amd64' go build -o builds/lebeben-go-macos/lebeben-go .
env GOOS='windows' GOARCH='amd64' go build -o builds/lebeben-go-win64/lebeben-go.exe .
env GOOS='linux' GOARCH='amd64' go build -o builds/lebeben-go-linux/lebeben-go .

chmod +x builds/lebeben-go-macos/lebeben-go
chmod +x builds/lebeben-go-win64/lebeben-go.exe
chmod +x builds/lebeben-go-linux/lebeben-go

tar -C builds -czvf lebeben-go-macos.tar.gz lebeben-go-macos
tar -C builds -czvf lebeben-go-win64.tar.gz lebeben-go-win64
tar -C builds -czvf lebeben-go-linux.tar.gz lebeben-go-linux
```

## Development/Release Process

- `npm i`
- make changes, rebuilt, test, etc
- bump the version in `package.json` and `package-lock.json`
- `git add .`
- `git commit -m "messages"`
  - precommit hooks will format go and js files
- `git tag vX.Y.Z`
  - make note of the leading `v`
  - should match the version in `package.json`
- `git push origin master --tags`
- wait for github actions to
  - build the go binaries for mac, linux, and windows
  - create a release on Github
  - publish a new version to npm, pointing to that release

## Performance

Comparison between the native javascript version, version pulled from `go get` and the version installed via `npm`. 30 rebuilds triggered by saving a single watched file, saving roughly every 250ms. The average of the 30 builds is shown below.

|Javascript|go get|npm i|
|---|---|---|
|~18s|~6s|~7s|
