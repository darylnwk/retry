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

#### func  Jitter

```go
func Jitter(backoff time.Duration) time.Duration
```
Jitter defines a full jitter strategy

#### func  NoBackoff

```go
func NoBackoff(_ uint, _ time.Duration) time.Duration
```
NoBackoff defines no backoff strategy

#### func  NoJitter

```go
func NoJitter(backoff time.Duration) time.Duration
```
NoJitter defines no jitter strategy

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
	// Attempts defines the number of retry attempts.
	// Default 1.
	Attempts uint

	// Backoff defines a backoff function that returns `time.Duration`.
	// This is applied on subsequent attempts.
	// Default no backoff.
	Backoff func(n uint, delay time.Duration) time.Duration

	// Jitter defines a jitter function that returns `time.Duration`.
	// Default no jitter.
	Jitter func(backoff time.Duration) time.Duration

	// Delay defines duration to delay.
	// Default 0.
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
