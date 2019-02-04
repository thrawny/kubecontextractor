# kubectl-extract

[![CircleCI](https://circleci.com/gh/thrawny/kubectl-extract.svg?style=svg)](https://circleci.com/gh/thrawny/kubectl-extract) 
[![Go Report Card](https://goreportcard.com/badge/github.com/thrawny/kubectl-extract)](https://goreportcard.com/report/github.com/thrawny/kubectl-extract)

Kubectl plugin to extract a kubernetes context including authinfo, cluster and namespace.
Useful to quickly share with others.

## Installation
```bash
go get -u github.com/thrawny/kubectl-extract
```

## Usage
Kubectl >= 1.12 has plugin support and the program can then be called using:
```
kubectl extract context
```

Otherwise call with:
```
kubectl-extract context
```

Example usage:
```bash
kubectl extract context --context foo > fooconfig
KUBECONFIG=./fooconfig kubectl get pods
```
