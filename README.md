<div align="center" class="no-border">
  <img src="/img/drift.jpg" alt="Drift" width="300" height="300"/>
  <br>

<b>Find configuration drifts on Kubernetes resources or Helm charts running in your cluster.</b>

  <a href="https://github.com/ajayk/drifter/releases">
    <img src="https://img.shields.io/github/v/release/ajayk/drifter">
  </a>
  <a href="https://goreportcard.com/report/github.com/ajayk/drifter">
    <img src="https://goreportcard.com/badge/github.com/ajayk/drifter">
  </a>

</div>

Drifter scans your cluster for installed kubernetes components ,
installed Helm charts, then cross-checks them against
the passed expectation file .

# Installing

Using Drifter is easy. First, use `go get` to install the latest version
of the library.

```
go get -u github.com/ajayk/drifter@latest
```

Next, include Drifter in your application:

```go
import "github.com/ajayk/drifter"
```

# Usage

```bash
drifter check -k /Users/drifter/.kube/config -c  examples/gcp-gke-check.yaml
```


