# Excercise 3: Add some struct validation

This excercise will show a basic input validation setup. With the application as it is now, you can basically input anything and Go will not complain and just persist it. There is a lot possible in this area. We choose to use a well know 3rd party package called [go-validate](github.com/go-playground/validator)

We'll start by telling Go which specific fields have validation rules. Open beer.go again and change the Beer structs Name and Brewery fields to enforce that these field should not be empty.

```go
       Name    string    `json:"name" validate:"required"`
       Brewery string    `json:"brewery" validate:"required"`
```

Now we need to edit PostBeer. Add the following logic right after the creation of the beer variable. (After line 46 in beer.go)

```go
	v := validator.New()
	err := v.Struct(beer)
	if err != nil {
		return beer, fmt.Errorf("error validating input: %+v", err)
	}
```

Also change the method signature

```go
func PostBeer(nb *NewBeer, db Beers) (Beer, error) {
```

Note we need to return an error now. If all is well, we can return a nil object to indicate all is ok. Add this to return statement at the end of the function.

```go
return beer, nil
```

We need to make a modification to the handlers package (which is in handlers.go) to make it handle the error PostBeer now returns. In the PostBeer method, change line 79 to look like the following.

```go
       rb, err := beer.PostBeer(&nb, app.d)
       if err != nil {
               w.WriteHeader(http.StatusBadRequest)
               w.Header().Set("Content-Type", "application/json")
               w.Write([]byte(`{ "error": "invalid input json" }`))
               return
       }
```

## Test the validation logic

Run the app

```bash
go run main.go
```

it might be you run into the following error:

```bash
bas@DESKTOP-RFVONSL: /mnt/c/data/golang-bits-n-bytes/Excercises/RESTService/Scaffold $ go run main.go
# github.com/ninckblokje/golang-bits-n-bytes/RESTService/beer
beer/beer.go:48:7: undefined: validator
zsh: exit 2     go run main.go
```

This is due to a missing import in beer.go. Fix it by adding the package to the import statement on the top of the file. Another option is typing the line out manually, VSCode with the Go plugin will then automatically pick-up the import.

```go
import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)
```

After this, retry running you app. It should work now.

If you now do a curl with an empty name or brewery, the creation of the beer will fail.

```bash
curl --location --request POST 'http://localhost:8080/beer' \
--header 'Content-Type: application/json' \
--data-raw '{ "brewery": "SynBeer Lab", "name": "" }'
```