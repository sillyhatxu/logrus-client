# logrus-client

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![Go-Version](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/sillyhatxu/logurs-client)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/sillyhatxu/logurs-client)](https://pkg.go.dev/github.com/sillyhatxu/logurs-client)
[![Build and Test](https://github.com/sillyhatxu/logurs-client/workflows/Build%20and%20Test/badge.svg?branch=master&event=push)](https://github.com/sillyhatxu/logurs-client/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sillyhatxu/logurs-client)](https://goreportcard.com/report/github.com/sillyhatxu/logurs-client)
[![codecov](https://codecov.io/gh/sillyhatxu/logurs-client/branch/master/graph/badge.svg)](https://codecov.io/gh/sillyhatxu/logurs-client)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://choosealicense.com/licenses/mit/)
[![Release](https://img.shields.io/github/release/sillyhatxu/logurs-client.svg?style=flat-square)](https://github.com/sillyhatxu/logurs-client/releases)


## Initialize your project

```
go mod init github.com/sillyhatxu/logrus-client
```

## Updating dependencies

```
go mod vendor
```

## verify dependencies

```
go mod verify
```

## remove dependencies that is not used

```
go mod tidy
```

## print dependence diagram

```
go mod graph
```

## download dependencies

```
go mod download
```

# Release Template

### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)
