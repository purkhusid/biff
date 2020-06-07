#! /usr/bin/env bash

go get "${1}"
./bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=godeps_macro.bzl%go_repositories