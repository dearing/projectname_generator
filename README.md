# projectname_generator
because project names are hard yo

hosted live: https://dearing.github.io/projectname_generator/

_This project is silly in reality but it's a demo of how to use wasm with go._

## Usage

This function is a wrapper of sorts to take an optional arg from javacript to size the array of randomized names.  Note we could accept and call javascript functions here as well.  Before being returned, the values are converted to javascript equivalents for the export function described next.


```go
func generateFunction(this js.Value, args []js.Value) interface{} {
	count := 1
	if len(args) != 0 {
		count = args[0].Int()
	}
	var names []interface{}
	for i := 0; i < count; i++ {
		names = append(names, GenerateProjectName())
	}

	return js.ValueOf(names[:])
}
```

Main is loaded in this setup so we export the generate function while keeping the running binary stuck in an inifinite loop by waiting on a goroutine that never gets content.

```go
func main() {
	println("app loaded")
	rand.Seed(time.Now().Unix())
	c := make(chan struct{}, 0)
	js.Global().Set("generate", js.FuncOf(generateFunction))
	<-c
}
```

Over in the frontend world we have an html document with some supporting javascript to interact with the wasm binary.  The wasm binary is loaded along with a javascript package provided by the go team.  With this, the generate function is availiable to javascript to call with cleaned up results that we simply print out to the a div with some styling.

```html
  <script src="wasm_exec.js"></script>
  <script>
    const go = new Go();
    let inst;

    WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject).then(
      (result) => {
        go.run(result.instance);
        document.getElementById("content").innerHTML +=
          "<pre><code>" + JSON.stringify(generate(21), null, " ") + "</code></pre>";
      }
    );
  </script>
```

## Chrome Dev Tools

`console.time('generate');generate(n);console.timeEnd('generate')`
|n|time|
|---|----|
|10000| 109.079833984375 ms|
|100000|  1508.261962890625 ms|
|1000000|  9836.356201171875 ms|
|10000000|  149811.64013671875 ms|
