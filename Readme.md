# jg-prefixwd

A command that wraps a command, prepends all file-pointer like patterns with the current dir and prints on the same file (stdout, stderr).
I use this as an external tool for Intellij IDEA to wrap my make commands.

## Usage

`jg-prefixwd make -f Makefile.wip test`
