package main

import (
	"errors"
	"fmt"
	"go.kuoruan.net/v8go-polyfills/fetch"
	"rogchap.com/v8go"
	"time"
)

func main() {
	fmt.Println("hello world")

	//ctx := v8.NewContext()                                  // creates a new V8 context with a new Isolate aka VM
	//ctx.RunScript("const add = (a, b) => a + b", "math.js") // executes a script on the global context
	//ctx.RunScript("const result = add(3, 4)", "main.js")    // any functions previously added to the context can be called
	//val, _ := ctx.RunScript("result", "value.js")           // return a value in JavaScript back to Go
	//fmt.Printf("addition result: %s", val)
	//iso := v8.NewIsolate() // create a new VM
	//// a template that represents a JS function
	//printfn := v8.NewFunctionTemplate(iso, func(info *v8.FunctionCallbackInfo) *v8.Value {
	//	fmt.Printf("%v", info.Args()) // when the JS function is called this Go callback will execute
	//	return nil                    // you can return a value back to the JS caller if required
	//})
	//global := v8.NewObjectTemplate(iso)       // a template that represents a JS Object
	//global.Set("print", printfn)              // sets the "print" property of the Object to our function
	//ctx := v8.NewContext(iso, global)         // new Context with the global Object set to our object template
	//ctx.RunScript("print('foo')", "print.js") // will execute the Go callback with a single argunent 'foo'

	iso := v8go.NewIsolate()
	global := v8go.NewObjectTemplate(iso)

	if err := fetch.InjectTo(iso, global); err != nil {
		panic(err)
	}

	ctx := v8go.NewContext(iso, global)

	val, err := ctx.RunScript("fetch('https://www.example.com').then(res => res.text())", "fetch.js")
	if err != nil {
		panic(err)
	}

	proms, err := val.AsPromise()
	if err != nil {
		panic(err)
	}
	done := make(chan bool, 1)

	go func() {
		for proms.State() == v8go.Pending {
			continue
		}
		done <- true
	}()

	select {
	case <-time.After(time.Second * 10):
		panic(errors.New("request timeout"))
	case <-done:
		html := proms.Result().String()
		fmt.Println(html)
	}
}
