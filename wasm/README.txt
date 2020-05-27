Set GOOS=js and GOARCH=wasm environment variables to compile for WebAssembly:

$ GOOS=js GOARCH=wasm go build -o main.wasm


To execute main.wasm in a browser, we’ll also need a JavaScript support file, and a HTML page to connect everything together.
Copy the JavaScript support file:

$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .