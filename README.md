
# kangaroo

Runs a command if signaled or in specific time intervals

## Installation

```sh
git clone https://github.com/alexcoder04/kangaroo.git
cd kangaroo

# build a binary and run it
go build .
./kangaroo

# install to your $GOPATH
go install .
```

## Usage

### Execute on signal

```sh
kangaroo -signal 2 echo "Hello World!"
```

This will execute `echo "Hello World!"` if kangaroo gets signaled with
`SIGRTMIN+2`. If no signal number is passed, kangaroo assumes `1`. If you pass
`0` as signal number, kangaroo will not listen at all.

### Execute in time intervals

```sh
kangaroo -interval 3 echo "Hello World!"
```

This will execute `echo "Hello World!"` every 3 seconds. If no number is passed,
kangaroo will not execute your command periodically.

