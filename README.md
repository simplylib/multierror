# multierror

[![Go Reference](https://pkg.go.dev/badge/github.com/simplylib/multierror.svg)](https://pkg.go.dev/github.com/simplylib/multierror)
[![Go Report Card](https://goreportcard.com/badge/github.com/simplylib/multierror)](https://goreportcard.com/report/github.com/simplylib/multierror)


a simple package for handling horizontal errors as opposed to veritical (fmt.Errorf %w error wrapping)

Usage example from [goreinstall](https://github.com/simplylib/goreinstall/blob/9c264ba86506d030c2392ae4b99ee070ee53f15c/gobin.go#L69):

```go
func getGoBinaryInfo(ctx context.Context, path string) (info *buildinfo.BuildInfo, err error) {
  ...
  defer func() {
    if err2 := f.Close(); err2 != nil {
      err = multierror.Append(err, fmt.Errorf("file could not be closed due to error (%w)", err2))
    }
  }()
  ...
}
```
