#!/usr/bin/env bash
git add .
git commit -m 'release logrus-client'
git push origin master
git tag v3.0.3
git push origin v3.0.3