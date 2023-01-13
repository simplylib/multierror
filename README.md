# multierror

[![Go Reference](https://pkg.go.dev/badge/github.com/simplylib/multierror.svg)](https://pkg.go.dev/github.com/simplylib/multierror)
[![Go Report Card](https://goreportcard.com/badge/github.com/simplylib/multierror)](https://goreportcard.com/report/github.com/simplylib/multierror)


a simple package for handling horizontal errors as opposed to veritical (fmt.Errorf %w error wrapping)

Usage example from [goreinstall](https://github.com/simplylib/goreinstall/blob/3d4ada803ef2305420450fb52782f6d515b840cc/gobin.go#L69):

```go
func getGoBinaryInfo(ctx context.Context, path string) (info *buildinfo.BuildInfo, err error) {
  ...
  defer func() {
    if err2 := f.Close(); err2 != nil {
      err = multierror.Append(err, err2)
    }
  }()
  ...
}
```

Note: this example really should have 
```go
err = multierror.Append(err, fmt.Errorf("file (%v) could not be closed due to (%w)"))
``` 
instead of
```go
err = multierror.Append(err, err2)
```
