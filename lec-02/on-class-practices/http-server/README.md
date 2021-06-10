# Project's Info

This project is a demo for running a simple http server. 
It will listen on a fixed port 8090 and serve 2 patterns along with 2 handlers on that:
- Hello handler: Just return "Hello guys" string when browse to http://localhost:8090/hello
- Hi with simple param reader: it read query param named your_name and return string like "Hi, [[your_name]]" when your_name param is filled, *Hi guys. I don't know your name because you don't enter the your_name query param* otherwise.

## Project's tree
```
├── handlers
│   ├── hello.go
│   └── hi_with_simple_param.go
├── server
│   └── server.go
├── README.md
└── main.go
```