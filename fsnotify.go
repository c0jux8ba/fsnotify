// Package fsnotify provides a cross-platform interface for file system notifications.
package fsnotify

import (
	"errors"
	"fmt"
	"strings"
)

// Event represents a file system notification event.
type Event struct {
	// Name is the path of the file or directory that triggered the event.
	Name string

	// Op describes the file operation that triggered the event.
	Op Op
}

// Op describes a set of file operations.
type Op uint32

// The set of operations that can trigger a notification.
const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

// String returns a human-readable description of the event.
func (e Event) String() string {
	return fmt.Sprintf("%q: %s", e.Name, e.Op)
}

// String returns a human-readable description of the operation(s).
func (op Op) String() string {
	var sb strings.Builder

	if op.Has(Create) {
		sb.WriteString("|CREATE")
	}
	if op.Has(Write) {
		sb.WriteString("|WRITE")
	}
	if op.Has(Remove) {
		sb.WriteString("|REMOVE")
	}
	if op.Has(Rename) {
		sb.WriteString("|RENAME")
	}
	if op.Has(Chmod) {
		sb.WriteString("|CHMOD")
	}

	if sb.Len() == 0 {
		return "[no events]"
	}
	return sb.String()[1:] // Trim leading "|"
}

// Has reports whether this operation includes the given Op.
func (op Op) Has(h Op) bool {
	return op&h != 0
}

// Common errors returned by the Watcher.
var (
	// ErrNonExistentWatch is returned when a watch is removed that does not exist.
	ErrNonExistentWatch = errors.New("fsnotify: can't remove non-existent watch")

	// ErrEventOverflow is returned when the kernel event buffer is full.
	// This can happen if events are not consumed fast enough.
	ErrEventOverflow = errors.New("fsnotify: queue or buffer overflow")

	// ErrClosed is returned when the watcher is closed.
	ErrClosed = errors.New("fsnotify: watcher already closed")
)
