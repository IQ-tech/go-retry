# Retry a function with exponential backoff

This project provide utilities to retry fallible function calls

# Installation

To install, use `go get`

```terminal
go get github.com/IQ-tech/go-retry
```

## Usage

## Retrying a function call

```go
// We will attempt the functioncall 3 times
// with an initial waiting time of 3 seconds.
retryOptions := retry.Options{
	Attempts:                  3,
	InitialTimeBetweenRetries: 3,
}

value, err = retry.Func(retryOptions, func() (interface{}, error) {
	return someservice.GetSomeValue()
})

// [err] will be non nill if the function didn't succeed after
// the maximum number of attempts has been reached.
if err != nil {
  panic(err)
}

fmt.Println(value)
```
