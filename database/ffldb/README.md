ffldb
=====

<<<<<<< HEAD
[![Build Status](https://travis-ci.org/navcoin/navd.png?branch=master)](https://travis-ci.org/navcoin/navd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://godoc.org/github.com/navcoin/navd/database/ffldb?status.png)](http://godoc.org/github.com/navcoin/navd/database/ffldb)
=======
[![Build Status](https://travis-ci.org/navcoin/navd.png?branch=master)]
(https://travis-ci.org/navcoin/navd)
>>>>>>> c7d17841... multi: update imports to navcoin's

Package ffldb implements a driver for the database package that uses leveldb for
the backing metadata and flat files for block storage.

This driver is the recommended driver for use with navd.  It makes use leveldb
for the metadata, flat files for block storage, and checksums in key areas to
ensure data integrity.

Package ffldb is licensed under the copyfree ISC license.

## Usage

This package is a driver to the database package and provides the database type
of "ffldb".  The parameters the Open and Create functions take are the
database path as a string and the block network.

```Go
db, err := database.Open("ffldb", "path/to/database", wire.MainNet)
if err != nil {
	// Handle error
}
```

```Go
db, err := database.Create("ffldb", "path/to/database", wire.MainNet)
if err != nil {
	// Handle error
}
```

<<<<<<< HEAD
=======
## Documentation

[![GoDoc](https://godoc.org/github.com/navcoin/navd/database/ffldb?status.png)]
(http://godoc.org/github.com/navcoin/navd/database/ffldb)

Full `go doc` style documentation for the project can be viewed online without
installing this package by using the GoDoc site here:
http://godoc.org/github.com/navcoin/navd/database/ffldb

You can also view the documentation locally once the package is installed with
the `godoc` tool by running `godoc -http=":6060"` and pointing your browser to
http://localhost:6060/pkg/github.com/navcoin/navd/database/ffldb

>>>>>>> c7d17841... multi: update imports to navcoin's
## License

Package ffldb is licensed under the [copyfree](http://copyfree.org) ISC
License.
