Elos Server
===========

Start
-----

 1. `go get github.com/elos/server`
 2. `cd github.com/elos/server`
 3. `go run server.go`

Create a user
-------------

  `POST /v1/users`

Websocket
---------

  var connection = new WebSocket('ws://localhost:8000/v1/authenticate', "id-key")
  connection.onmessage = function(event) { console.log(JSON.parse(event.Data)) }
  msg = {action: "POST", data:{user: {name: "Nick Landolfi"}}} .// example
  connection.send(JSON.stringify(msg))



