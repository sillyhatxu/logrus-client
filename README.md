# logrus-client


## Initialize your project

```
dep init
```

## Adding a dependency

```
dep ensure -add github.com/foo/bar github.com/foo/baz...

dep ensure -add github.com/foo/bar@1.0.0 github.com/foo/baz@master
```

## Updating dependencies

```
dep ensure -update github.com/sillyhatxu/convenient-utils

dep ensure -update
```

# Release Template

### Feature

* [NEW] Support for Go Modules [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

---

### Bug fix

* [FIX] Truncate Latency precision in long running request [#17](https://github.com/sillyhatxu/convenient-utils/issues/17)

###

```
git tag v1.0.2
git push origin v1.0.2
```
