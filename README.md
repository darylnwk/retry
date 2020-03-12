# retry
--
    import "github.com/darylnwk/retry"


## Usage

#### func  NoBackoff

```go
func NoBackoff() time.Duration
```
NoBackoff defines no backoff strategy

#### type Errors

```go
type Errors []error
```

Errors defines a list of `error`

#### func (Errors) Error

```go
func (errs Errors) Error() string
```

#### type Retryer

```go
type Retryer struct {
	// Attempts defines the number of retry attempts
	Attempts uint

	// Backoff defines a backoff function that returns `time.Duration`
	// This is applied on subsequent attempts.
	Backoff func() time.Duration
}
```

Retryer defines a retryer

#### func (Retryer) Do

```go
func (retryer Retryer) Do(fn func() error) (bool, Errors)
```
Do executes `fn` and returns success if `fn` executed succesfully and any errors
that occurred
