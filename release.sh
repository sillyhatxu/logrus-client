#!/usr/bin/env bash
#git checkout -b go-mod-release-new-version v2.0.0
#git tag -a -m "Tag version 2.0.0, go mod release" v2.0.0
git add .
git commit -m 'release logrus-client'
git push origin master
git tag v2.0.0
git push origin v2.0.0