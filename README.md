# logepi-goclient

Logepi Golang client

```
go get github.com/tsirolnik/logepi-goclient
```


```go
import "github.com/tsirolnik/logepi-goclient"

...

lgpclient.Use("someaddress:6080")
if err := lgpclient.Log("table", lgpclient.LogData{
  "key":"value"
}); err != nil {
  log.Fatal("Whoops! " + err.Error())
}
```