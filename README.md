navd
====

[![Build Status](https://travis-ci.org/aguycalled/navd.png?branch=master)](https://travis-ci.org/aguycalled/navd)

navd is an alternative full node navcoin implementation written in Go (golang).

This project is currently under active development and is in a Beta state.  It
is extremely stable and has been in production use since October 2013.

It properly downloads, validates, and serves the block chain using the exact
rules (including consensus bugs) for block acceptance as Navcoin Core.  We have
taken great care to avoid navd causing a fork to the block chain.  It includes a
full block validation testing framework which contains all of the 'official'
block acceptance tests (and some additional ones) that is run on every pull
request to help ensure it properly follows consensus.  Also, it passes all of
the JSON test data in the Navcoin Core code.

It also properly relays newly mined blocks, maintains a transaction pool, and
relays individual transactions that have not yet made it into a block.  It
ensures all individual transactions admitted to the pool follow the rules
required by the block chain and also includes more strict checks which filter
transactions based on miner requirements ("standard" transactions).

One key difference between navd and Navcoin Core is that navd does *NOT* include
wallet functionality and this was a very intentional design decision.  See the
blog entry [here](https://blog.conformal.com/navd-not-your-moms-navcoin-daemon)
for more details.  This means you can't actually make or receive payments
directly with navd.  That functionality is provided by the
[navwallet](https://github.com/aguycalled/navwallet) and
[Paymetheus](https://github.com/btcsuite/Paymetheus) (Windows-only) projects
which are both under active development.

## Requirements

[Go](http://golang.org) 1.8 or newer.

## Installation

#### Windows - MSI Available

https://github.com/aguycalled/navd/releases

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
$ go env GOROOT GOPATH
```

NOTE: The `GOROOT` and `GOPATH` above must not be the same path.  It is
recommended that `GOPATH` is set to a directory in your home directory such as
`~/goprojects` to avoid write permission issues.  It is also recommended to add
`$GOPATH/bin` to your `PATH` at this point.

- Run the following commands to obtain navd, all dependencies, and install it:

```bash
$ go get -u github.com/Masterminds/glide
$ git clone https://github.com/aguycalled/navd $GOPATH/src/github.com/aguycalled/navd
$ cd $GOPATH/src/github.com/aguycalled/navd
$ glide install
$ go install . ./cmd/...
```

- navd (and utilities) will now be installed in ```$GOPATH/bin```.  If you did
  not already add the bin directory to your system path during Go installation,
  we recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- Run the following commands to update navd, all dependencies, and install it:

```bash
$ cd $GOPATH/src/github.com/aguycalled/navd
$ git pull && glide install
$ go install . ./cmd/...
```

## Getting Started

navd has several configuration options avilable to tweak how it runs, but all
of the basic operations described in the intro section work with zero
configuration.

#### Windows (Installed from MSI)

Launch navd from your Start menu.

#### Linux/BSD/POSIX/Source

```bash
$ ./navd
```

## IRC

- irc.freenode.net
- channel #navd
- [webchat](https://webchat.freenode.net/?channels=navd)

## Issue Tracker

The [integrated github issue tracker](https://github.com/aguycalled/navd/issues)
is used for this project.

## Documentation

The documentation is a work-in-progress.  It is located in the [docs](https://github.com/aguycalled/navd/tree/master/docs) folder.

## GPG Verification Key

All official release tags are signed by Conformal so users can ensure the code
has not been tampered with and is coming from the btcsuite developers.  To
verify the signature perform the following:

- Download the Conformal public key:
  https://raw.githubusercontent.com/aguycalled/navd/master/release/GIT-GPG-KEY-conformal.txt

- Import the public key into your GPG keyring:
  ```bash
  gpg --import GIT-GPG-KEY-conformal.txt
  ```

- Verify the release tag with the following command where `TAG_NAME` is a
  placeholder for the specific tag:
  ```bash
  git tag -v TAG_NAME
  ```

## License

navd is licensed under the [copyfree](http://copyfree.org) ISC License.
