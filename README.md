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

Drifter check returns either an exit code of `0` (pass)  or `2` (fail)

Usage Demo

[![asciicast](https://asciinema.org/a/SHFqgQMyAFifMsAMoVBm58sxD.svg)](https://asciinema.org/a/SHFqgQMyAFifMsAMoVBm58sxD)

### Drifter Schema:

Drifter yaml is easy to configure to check the expectations
Currently supports 8 different type of validators

- helm
- namespaces
- deployments
- daemonsets
- statefulsets
- secrets
- ingress classes
- storage classes

check [examples](examples) directory for each different type of validator

```yaml
helm:
  components:
    - name: ingress-nginx
      version: 4.2.3
      appVersion: 1.2.0
    - name: external-secrets-operator
      version: 0.6.8 # just chart version check not checking for appVersion here 

kubernetes:
  namespaces:
    - name: kube-system
    - name: es

  daemonsets:
    - namespace: kube-system
      names:
        - anetd
        - nvidia-gpu-device-plugin
    - namespace: gmp-public
      names:
        - node-exporter

  deployments:
    - namespace: kube-system
      names:
        - kube-dns

  statefulsets:
    - namespace: gkebackup
      names:
        - gkebackup-agent

  storage:
    classes:
      - filestore-premium-rwx
      - filestore-standard-rwx

  ingress:
    classes:
      - nginx
```

