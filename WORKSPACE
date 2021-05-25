workspace(
    name = "swiss_dasch_dsp_repository",
    managed_directories = {"@npm": ["node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

##################################
# Base                           #
##################################

#
# bazel-skylib 1.0.2 released 2019.10.09 (https://github.com/bazelbuild/bazel-skylib/releases/tag/1.0.2)
#
skylib_version = "1.0.2"

http_archive(
    name = "bazel_skylib",
    sha256 = "97e70364e9249702246c0e9444bccdc4b847bed1eb03c5a3ece4f83dfe6abc44",
    type = "tar.gz",
    url = "https://github.com/bazelbuild/bazel-skylib/releases/download/{}/bazel-skylib-{}.tar.gz".format(skylib_version, skylib_version),
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

##################################
# Javascript                     #
##################################

#
# Download the rules_nodejs repository
#
rules_nodejs_version = "3.1.0"

rules_nodejs_version_sha256 = "dd4dc46066e2ce034cba0c81aa3e862b27e8e8d95871f567359f7a534cccb666"

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = rules_nodejs_version_sha256,
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/%s/rules_nodejs-%s.tar.gz" % (rules_nodejs_version, rules_nodejs_version)],
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories", "yarn_install")

# NOTE: this rule installs nodejs, npm, and yarn, but does NOT install
# your npm dependencies into your node_modules folder.
# You must still run the package manager to do this.
node_repositories(
    node_version = "14.15.4",
    package_json = ["//:package.json"],
)

load("@build_bazel_rules_nodejs//:package.bzl", "rules_nodejs_dev_dependencies")

rules_nodejs_dev_dependencies()

yarn_install(
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

##################################
# Protobuf                       #
##################################

#
# Download Protobuf
#
http_archive(
    name = "com_google_protobuf",
    sha256 = "9748c0d90e54ea09e5e75fb7fac16edce15d2028d4356f32211cfa3c0e956564",
    strip_prefix = "protobuf-3.11.4",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.11.4.zip"],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

##################################
# Go                             #
##################################

#
# Download the Go rules
#
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7904dbecbaffd068651916dce77ff3437679f9d20e1a7956bff43826e7645fcc",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.25.1/rules_go-v0.25.1.tar.gz",
    ],
)

#
# Download Gazelle.
#
http_archive(
    name = "bazel_gazelle",
    sha256 = "222e49f034ca7a1d1231422cdb67066b885819885c356673cb1f72f748a3c9d4",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.3/bazel-gazelle-v0.22.3.tar.gz",
    ],
)

# Load macros and repository rules.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

# Declare indirect dependencies and register toolchains.
go_rules_dependencies()

go_register_toolchains(
    version = "1.16",
    # nogo = "@io_bazel_rules_go//:tools_nogo"
)

gazelle_dependencies()

##################################
# Buildifier (Bazel linter)      #
##################################

http_archive(
    name = "com_github_bazelbuild_buildtools",
    strip_prefix = "buildtools-master",
    url = "https://github.com/bazelbuild/buildtools/archive/master.zip",
)

##################################
# Packaging      (tar, rpm, etc.)#
##################################

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_pkg",
    sha256 = "038f1caa773a7e35b3663865ffb003169c6a71dc995e39bf4815792f385d837d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_pkg/releases/download/0.4.0/rules_pkg-0.4.0.tar.gz",
        "https://github.com/bazelbuild/rules_pkg/releases/download/0.4.0/rules_pkg-0.4.0.tar.gz",
    ],
)

load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()

##################################
# Docker                         #
##################################

rules_docker_version = "0.14.4"

rules_docker_version_sha256 = "4521794f0fba2e20f3bf15846ab5e01d5332e587e9ce81629c7f96c793bb7036"

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = rules_docker_version_sha256,
    strip_prefix = "rules_docker-%s" % rules_docker_version,
    url = "https://github.com/bazelbuild/rules_docker/releases/download/v%s/rules_docker-v%s.tar.gz" % (rules_docker_version, rules_docker_version),
)

load("@io_bazel_rules_docker//repositories:repositories.bzl", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load("@io_bazel_rules_docker//repositories:pip_repositories.bzl", "pip_deps")

pip_deps()

load("@io_bazel_rules_docker//go:image.bzl", go_image_repos = "repositories")

go_image_repos()

# load container_pull method
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_pull",
)

# get json-server docker image
container_pull(
    name = "json_server",
    # 'tag' is also supported, but digest is encouraged for reproducibility.
    digest = "sha256:cba2406ac6510e72cd7a0f85becb51b9f8adcb1622c509d2f69392627729b1d8",
    registry = "docker.io",
    repository = "daschswiss/json-server",
)
