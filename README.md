# Biff - A Bazel VCS diffing tool

Biff tries to answer a simple question in your Bazel monorepository.

_What Bazel targets were affected by my VCS changeset?_

Bazel can't answer this question because it does not have any information about your VCS setup.

## How do I use this thing

You can run the tool straight from this repo by running

`./bazel run //:biff -- <OPTIONS HERE>`

or you can build the tool and copy the binary from `bazel-bin`

`./bazel build //:biff`

`cp bazel-bin/biff_/biff <YOUR LOCATION>`

`biff <OPTIONS HERE>`

### Commands

#### Calculate

Calculates the sha256 checksum for each target in the graph and outputs it to a file.
This file can then be diffed against another output with the compare command

Usage:
  biff calculate [flags]

Flags:
      --bazel string       Location of Bazel executable. By default uses bazel from path
  -h, --help               help for calculate
      --out string         Where the output should be written to (Required)
      --workspace string   Path to the workspace root (Required)

#### Compare

Compares two outputs from the calculate command and outputs a file that can be
fed to the '--target_pattern_file' option in bazel

Usage:
  biff compare [flags]

Flags:
  -h, --help           help for compare
      --left string    The 'left' side of the comparison (Required)
      --out string     Where the output should be written to (Required)
      --right string   The 'right' side of the comparison (Required)
