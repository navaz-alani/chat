# chat

This is a real-time messaging application written in Golang. It works using the
concurrency primitives in Go. Real-time communication is done over the websocket
protocol.

There are two components to this application (so far). The main component is the
`chat` binary which is essentially the backend for the application. This backend
can be interacted with using the client component of the application. So far
this component consists of only the `client` binary which is a command-line
application. There are plans to make a web application for a more pleasing
interface.

Here are the build commands:

* `make all` - produce both the `chat` and the `client` binaries
* `make chat` - produce just the `chat` binary
* `make client` - produce just the `client` binary

Here are some maintenance commands:

* `make tidy` - reconcile the source code dependencies and format source code
  using `go fmt`
* `make clean` - delete the produced binaries
