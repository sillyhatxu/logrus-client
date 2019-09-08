#!/usr/bin/env bash
git add .
git commit -m 'release logrus-client'
git push origin master
git tag v2.0.0
git push origin v2.0.0