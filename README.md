# yamlchain

## Overview

`yamlchain` is a CLI tool that combines multiple YAML files into a single one, inserting a separator line (`---`) between each file. 

## Installation

To install `yamlchain`:

```shell
go install github.com/modatwork/yamlchain@latest
```

Make sure your GOPATH/bin is in your PATH to access the `yamlchain` command globally.

## Usage

To combine multiple YAML files into one, simply run:

```bash
yamlchain file1.yaml file2.yaml ... > combined.yaml
```

Or you can apply the combined YAML directly to your Kubernetes cluster:

```bash
yamlchain file1.yaml file2.yaml ... | kubectl apply -f -
```
