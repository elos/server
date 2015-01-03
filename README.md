Elos Server
-----------

##### Start

Requires mongod.

```bash
    go get github.com/elos/server
    cd github.com/elos/server
    go run server.go
```

##### Test

```bash
    ginkgo -r
```

See more(link needed)

##### Create a user

  `POST /v1/users`

###### Websocket

``` coffeescript
connection = new WebSocket "ws://localhost:8000/v1/authenticate", "#{id}-#{key}"
connection.onmessage = (event) -> console.log JSON.parse event.data
msg =
  action: "POST"
  data:
    user:
      name: "Nick Landolfi"
connection.send JSON.stringify msg
```


