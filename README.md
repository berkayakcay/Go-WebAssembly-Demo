# Go-WebAssembly-Demo

### Build

*Get wasm.js*

```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

edit goland build options to access "syscall/js"

```bash
GOOS=js GOARCH=wasm go build -o main.wasm
```

### Launch

You can use live-server or http-server to launch

```bash
npm install -g live-server 
```

