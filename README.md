# New Relic Go Agent

## Beta

This is beta software.  The Go Agent requires a beta token.  Getting a token is
easy!

1. Agree to the click-through [Beta Agreement](http://goo.gl/forms/Rcv1b10Qvt1ENLlr1)
2. Once your account is approved, we will email you a beta token, usually
   within the same business day.
3. Add the beta token to your config (see below for details).

Please join our [Go Agent Beta
Forum](https://discuss.newrelic.com/c/language-agents/go-agent-beta) to tell us
how the Go Agent works for you, what you'd like to see and how we can improve
it.  We're eager to hear your feedback!

Breaking changes may be made before release 1.0.

## Description

The New Relic Go Agent allows you to monitor your Go applications with New
Relic.  It helps you track transactions, outbound requests, database calls, and
other parts of your Go application's behavior and provides a running overview of
garbage collection, goroutine activity, and memory use.

## Requirements

Go 1.3+ is required, due to the use of http.Client's Timeout field.

Linux and OS X are supported.

## Getting Started

Here are the basic steps to instrumenting your application.  For more
information, see [GUIDE.md](GUIDE.md).

#### Step 0: Installation

Installing the Go Agent is the same as installing any other Go library.  The
simplest way is to run:

```
go get github.com/newrelic/go-agent
```

Then import the `github.com/newrelic/go-agent` package in your application.

#### Step 1: Create a Config and an Application

In your `main` function or an `init` block:

```go
config := newrelic.NewConfig("Your Application Name", "__YOUR_NEW_RELIC_LICENSE_KEY__")
config.BetaToken = "__YOUR_NEW_RELIC_BETA_TOKEN__"
app, err := newrelic.NewApplication(config)
```

[more info](GUIDE.md#config-and-application), [application.go](api/application.go),
[config.go](api/config.go)

#### Step 2: Add Transactions

Transactions time requests and background tasks.  Use `WrapHandle` and
`WrapHandleFunc` to create transactions for requests handled by the `http`
standard library package.

```go
http.HandleFunc(newrelic.WrapHandleFunc(app, "/users", usersHandler))
```

Alternatively, create transactions directly using the application's
`StartTransaction` method:

```go
txn := app.StartTransaction("myTxn", optionalResponseWriter, optionalRequest)
defer txn.End()
```

[more info](GUIDE.md#transactions), [transaction.go](api/transaction.go)

#### Step 3: Instrument Segments

Segments show you where time in your transactions is being spent.  At the
beginning of important functions, add:

```go
defer txn.EndSegment(txn.StartSegment(), "mySegmentName")
```

[more info](GUIDE.md#segments), [segments.go](api/segments.go)

## Runnable Example

[example/main.go](./example/main.go) is an example that will appear as "My Go
Application" in your New Relic applications list.  To run it:

```
env NEW_RELIC_LICENSE_KEY=__YOUR_NEW_RELIC_LICENSE_KEY__LICENSE__ \
    NEW_RELIC_BETA_TOKEN=__YOUR_NEW_RELIC_BETA_TOKEN__ \
    go run example/main.go
```

Some endpoints exposed are [http://localhost:8000/](http://localhost:8000/)
and [http://localhost:8000/notice_error](http://localhost:8000/notice_error)


## Basic Example

Before Instrumentation

```go
package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8000", nil)
}
```

After Instrumentation

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/newrelic/go-agent"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world")
}

func main() {
	// Create a config.  You need to provide the desired application name
	// and your New Relic license key.
	cfg := newrelic.NewConfig("My Go Application", "__YOUR_NEW_RELIC_LICENSE_KEY__")

	// Add the beta token emailed to you after signing:
	//   http://goo.gl/forms/Rcv1b10Qvt1ENLlr1
	cfg.BetaToken = "__YOUR_NEW_RELIC_BETA_TOKEN__"

	// Create an application.  This represents an application in the New
	// Relic UI.
	app, err := newrelic.NewApplication(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wrap helloHandler.  The performance of this handler will be recorded.
	http.HandleFunc(newrelic.WrapHandleFunc(app, "/", helloHandler))
	http.ListenAndServe(":8000", nil)
}
```
