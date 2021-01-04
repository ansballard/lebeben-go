## Build across platforms from powershell

```powershell
$env:GOOS='darwin';$env:GOARCH='amd64';go build -o builds/lebeben-go-macos/lebeben-go .;Remove-Item Env:/GOOS;Remove-Item Env:/GOARCH;
$env:GOOS='windows';$env:GOARCH='amd64';go build -o builds/lebeben-go-win64/lebeben-go.exe .;Remove-Item Env:/GOOS;Remove-Item Env:/GOARCH;
$env:GOOS='linux';$env:GOARCH='amd64';go build -o builds/lebeben-go-linux/lebeben-go .;Remove-Item Env:/GOOS;Remove-Item Env:/GOARCH;

# need to make the above files executable

tar -C builds -czvf lebeben-go-macos.tar.gz lebeben-go-macos
tar -C builds -czvf lebeben-go-win64.tar.gz lebeben-go-win64
tar -C builds -czvf lebeben-go-linux.tar.gz lebeben-go-linux
```

## Build across platforms in bash (includes making them executable)

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
