# go-angellist

go-angellist is a Go client library for accessing the [AngelList API](https://angel.co/api).

NOTE: This is not a complete implementation. I needed to extract some Startup and User
data so those are the only API calls available. If you would like to add functionality
please submit a pull request.

## Usage

```go
import "github.com/bradrydzewski/go-angellist/angellist"
```

Construct a new AngelList client, then use the various services on the client
to access different parts of the API. For example, to retrieve a User by id:

```go
client := angellist.NewGuest()
user, err := client.Users.Get(206401)
```

The above example is executing the API calls as an anonymous user. Hoewver,
some API calls require authentication. When creating a client you can provide
your OAuth2 token:

```go
client := angellist.New("62b1f3bc5dec16c8")
```

