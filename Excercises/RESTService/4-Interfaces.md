# Excercise 4: Convert beer storage to interface

This is a pretty lengthy excercise, showing the way Go implement the power of interfaces. We will first refactor the in-memory database in a way it implements the interface we will design at the start. After this is done, we will implement sqlite support in the persistence layer.

## Define and implement the interface

Let's start by defining the interface. Open beers.go and add the following interface definition. An interface definition is a list of methods the type should implement. It does not define variables.

```go
// BeerStore defines the methods for Beer Storage implementations
type BeerStore interface {
       PostBeer(nb *NewBeer) (Beer, error)
       GetBeer(i string) (Beer, error)
       GetAllBeers() ([]Beer, error)
}
```

Now we have an interface, we will rename our in memory map to be called something more descriptive.

change

```go
type Beers map[string]Beer
```

to

```go
type InMemoryBeerStore map[string]Beer
```

Now, we will change the existing PostBeer, GetBeer and GetAllBeers methods to be methods of the InMemoryBeerStore type.

```go
// OLD: func PostBeer(nb *NewBeer, db Beers) (Beer, error) {
func (db InMemoryBeerStore) PostBeer(nb *NewBeer) (Beer, error) {
```

```go
// OLD func GetBeer(i string, db Beers) (Beer, error) {
func (db InMemoryBeerStore) GetBeer(i string) (Beer, error) {
```

Let's also add the GetAllBeers method

```go
func(bs InMemoryBeerStore) GetAllBeers() ([]Beer, error) {
       var beers []Beer
       for _, value := range bs {
               beers = append(beers, value)
       }

       return beers, nil
}
```

Change the SynTappd type to look like this. It will now accept the interface instead of a map of beers

```go
// Syntappd hold the components to run this REST service
type Syntappd struct {
    l *log.Logger
    d beer.BeerStore
}
```

Change the New function to reflect this.

```go
func New(beerstore beer.BeerStore) *Syntappd {
    // Setup Syntappd object
    app := Syntappd{
        l: log.New(os.Stdout, "", 0),
        d: beerstore,
    }
    return &app
}
```

In the PostBeer handler, change this:

```go
// OLD  rb, err := beer.PostBeer(&nb, app.d)
       rb, err := app.d.PostBeer(&nb)
```

In the GetBeer handler, change this:

```go
// OLD b, err := beer.GetBeer(vars["id"], app.d)
    b, err := app.d.GetBeer(vars["id"])
``` 

Remove the for loop from GetAllBeers, we will not use it anymore due to the new GetAllBeers function within the interface. (Which is a different function!)

```go
// REMOVE ALL OG THIS
        var beers []beer.Beer

       // couple of lines down

       for _, value := range app.d {
               beers = append(beers, value)
       }
```

In the same function, change the object written by the json encoder.

```go
// OLD err := json.NewEncoder(w).Encode(beers)
err := json.NewEncoder(w).Encode(app.d)
```

This will change the output format of the getallbeers functions a little bit. For this excercise that does not really matter.

Last change is in main.go. Pass an InMemoryBeerStore for now, since we renamed the object.

```go
// OLD app := handlers.New(beer.Beers{})
       app := handlers.New(beer.InMemoryBeerStore{})
```

Now is a good time to start the app, and check if adding and retrieving beers still work. Use the curl statements from previous statements.
