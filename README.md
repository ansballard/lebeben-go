## Build across platforms from powershell

```powershell
$env:GOOS='darwin';$env:GOARCH='amd64';go build -o builds/lebeben-go-macos/lebeben-go .;Remove-Item Env:/GOOS;Remove-Item Env:/GOARCH;
$env:GOOS='windows';$env:GOARCH='amd64';go build -o builds/lebeben-go-win64/lebeben-go.exe .;Remove-Item Env:/GOOS;Remove-Item Env:/GOARCH;
$env:GOOS='linux';$env:GOARCH='amd64';go build -o builds/lebeben-go-linux/lebeben-go .;Remove-Item Env:/GOOS;Remove-Item Env:/GOARCH;

tar -C builds -czvf lebeben-go-macos.tar.gz lebeben-go-macos
tar -C builds -czvf lebeben-go-win64.tar.gz lebeben-go-win64
tar -C builds -czvf lebeben-go-linux.tar.gz lebeben-go-linux
```

