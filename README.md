# retry
--
    import "github.com/darylnwk/retry"


## Usage

#### func  ConstantBackoff

```go
func ConstantBackoff(_ uint, delay time.Duration) time.Duration
```
ConstantBackoff defines constant backoff strategy

#### func  ExponentialBackoff

```go
func ExponentialBackoff(n uint, delay time.Duration) time.Duration
```
ExponentialBackoff defines exponential backoff strategy

#### func  NoBackoff

```go
func NoBackoff(_ uint, _ time.Duration) time.Duration
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
	Backoff func(n uint, delay time.Duration) time.Duration

	// Delay defines duration to delay
	Delay time.Duration
}
```

Retryer defines a retryer

#### func (Retryer) Do

```go
func (retryer Retryer) Do(fn func() error) (bool, Errors)
```
Do executes `fn` and returns success if `fn` executed succesfully and any errors
that occurred
