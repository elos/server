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

[See more](https://github.com/elos/documentation/blob/master/server/testing.md)

##### Create a user

  `POST /v1/users`

###### Websocket

``` javascript
connection = new WebSocket("ws://localhost:8000/v1/authenticate", "[id]-[key]")
connection.onmessage = function(event) { console.log(JSON.parse(event.data)); }
msg = {
    action: "POST"
    data: {
        event: {
            name: "This is a new event"
        }
    }
}
connection.send(JSON.stringify(msg))
```


