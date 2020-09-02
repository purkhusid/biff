workspace(
    name = "biff",
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# Bazel skylib
git_repository(
    name = "bazel_skylib",
    commit = "327d61b5eaa15c11a868a1f7f3f97cdf07d31c58",
    remote = "https://github.com/bazelbuild/bazel-skylib.git",
    shallow_since = "1572441481 +0100",
)

load("@bazel_skylib//lib:versions.bzl", "versions")

versions.check(
    minimum_bazel_version = "3.2.0",
    maximum_bazel_version = "3.2.0",
)

# Protobuf dependencies
http_archive(
    name = "com_google_protobuf",
    sha256 = "761bfffc7d53cd01514fa237ca0d3aba5a3cfd8832a71808c0ccc447174fd0da",
    strip_prefix = "protobuf-3.11.1",
    urls = ["https://github.com/protocolbuffers/protobuf/releases/download/v3.11.1/protobuf-all-3.11.1.tar.gz"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# Golang rules
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "a8d6b1b354d371a646d2f7927319974e0f9e52f73a2452d2b3877118169eb6bb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz",
    ],
)

gazelle_git_hash = "b6658b6a47d9f0a06988ff15233168dcc713b536"

http_archive(
    name = "bazel_gazelle",
    sha256 = "d1e4aaa733992a1b00084dde808eb2dfe2e325c71e31ed74184376a449b2aeac",
    strip_prefix = "bazel-gazelle-%s" % gazelle_git_hash,
    type = "zip",
    url = "https://github.com/bazelbuild/bazel-gazelle/archive/%s.zip" % gazelle_git_hash,
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

gazelle_dependencies()

load("//:godeps_macro.bzl", "go_repositories")

# gazelle:repository_macro godeps_macro.bzl%go_repositories
go_repositories()
