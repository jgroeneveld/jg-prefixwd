# jg-prefixwd

A command that wraps a command, prepends all file-pointer like patterns with the current dir and prints on the same file (stdout, stderr).
I use this as an external tool for Intellij IDEA to wrap my make commands.

Turns:
`package/my_model.go:37:19: expected ';', found 'IDENT' MyModel`

into 

`/Users/jgroeneveld/go/src/github.com/jgroeneveld/project/package/my_model.go:37:19: expected ';', found 'IDENT' MyModel`

## Usage

`jg-prefixwd make -f Makefile.wip test`
