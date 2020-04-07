// AUTO GENERATED by make_open_tests.go - DO NOT EDIT - use go generate to rebuild

package vfs

import (
	"io"
	"os"
)

// openTest describes a test of OpenFile
type openTest struct {
	flags               int
	what                string
	openNonExistentErr  error
	readNonExistentErr  error
	writeNonExistentErr error
	openExistingErr     error
	readExistingErr     error
	writeExistingErr    error
	contents            string
}

// openTests is a suite of tests for OpenFile with all possible
// combination of flags.  This obeys Unix semantics even on Windows.
var openTests = []openTest{
	{
		flags:               os.O_RDONLY,
		what:                "os.O_RDONLY",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_EXCL,
		what:                "os.O_RDONLY|os.O_EXCL",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE,
		what:                "os.O_RDONLY|os.O_CREATE",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_EXCL,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_EXCL",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND,
		what:                "os.O_RDONLY|os.O_APPEND",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_EXCL,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_EXCL",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    EBADF,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: EBADF,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  EINVAL,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     EINVAL,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY,
		what:                "os.O_WRONLY",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HELlo",
	}, {
		flags:               os.O_WRONLY | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HELlo",
	}, {
		flags:               os.O_WRONLY | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_EXCL,
		what:                "os.O_WRONLY|os.O_EXCL",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HELlo",
	}, {
		flags:               os.O_WRONLY | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_EXCL | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HELlo",
	}, {
		flags:               os.O_WRONLY | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE,
		what:                "os.O_WRONLY|os.O_CREATE",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HELlo",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HELlo",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_EXCL,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_EXCL",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_EXCL | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND,
		what:                "os.O_WRONLY|os.O_APPEND",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_EXCL,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_EXCL",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_EXCL | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     EBADF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_WRONLY|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  EBADF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR,
		what:                "os.O_RDWR",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "heHEL",
	}, {
		flags:               os.O_RDWR | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_SYNC,
		what:                "os.O_RDWR|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "heHEL",
	}, {
		flags:               os.O_RDWR | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_EXCL,
		what:                "os.O_RDWR|os.O_EXCL",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "heHEL",
	}, {
		flags:               os.O_RDWR | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDWR|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "heHEL",
	}, {
		flags:               os.O_RDWR | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_CREATE,
		what:                "os.O_RDWR|os.O_CREATE",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "heHEL",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_CREATE|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_SYNC,
		what:                "os.O_RDWR|os.O_CREATE|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "heHEL",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_CREATE|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_EXCL,
		what:                "os.O_RDWR|os.O_CREATE|os.O_EXCL",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_CREATE|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_APPEND,
		what:                "os.O_RDWR|os.O_APPEND",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_SYNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_EXCL,
		what:                "os.O_RDWR|os.O_APPEND|os.O_EXCL",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  ENOENT,
		readNonExistentErr:  nil,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_SYNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "helloHEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     nil,
		readExistingErr:     io.EOF,
		writeExistingErr:    nil,
		contents:            "HEL",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_EXCL,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_EXCL",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_SYNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	}, {
		flags:               os.O_RDWR | os.O_APPEND | os.O_CREATE | os.O_EXCL | os.O_SYNC | os.O_TRUNC,
		what:                "os.O_RDWR|os.O_APPEND|os.O_CREATE|os.O_EXCL|os.O_SYNC|os.O_TRUNC",
		openNonExistentErr:  nil,
		readNonExistentErr:  io.EOF,
		writeNonExistentErr: nil,
		openExistingErr:     EEXIST,
		readExistingErr:     nil,
		writeExistingErr:    nil,
		contents:            "hello",
	},
}
