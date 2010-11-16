#!/bin/sh
git submodule update --init || exit 1
(cd vendor/log4go && ./build.sh) || exit 2
gd -I vendor/log4go/src -o tschunk-login
