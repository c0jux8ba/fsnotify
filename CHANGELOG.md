# Changelog

1.10.1 2026-05-04
-----------------

### Changes and fixes

- inotify: don't remove sibling watches sharing a path prefix ([#754])

- inotify, windows: don't rename sibling watches sharing a path prefix
  ([#755])


[#754]: https://github.com/fsnotify/fsnotify/pull/754
[#755]: https://github.com/fsnotify/fsnotify/pull/755


1.10.0 2026-04-30
-----------------
This version of fsnotify needs Go 1.23.

### Changes and fixes

- inotify: improve initialization error message ([#731])

- inotify: send Rename event if recursive watch is renamed ([#696])

- inotify: avoid copying event buffers when reading names ([#741])

- kqueue: skip dangling symlinks (ENOENT) in watchDirectoryFiles, so a
  bad entry no longer aborts Watcher.Add for the whole directory ([#748])

- kqueue: drop watches directly in Close() to fix a file descriptor leak
  when recycling watchers ([#740])

- windows: fix nil pointer dereference in remWatch ([#736])

- windows: lock watch field updates against concurrent WatchList to fix
  a race introduced in v1.9.0 ([#709], [#749])


[#696]: https://github.com/fsnotify/fsnotify/pull/696
[#709]: https://github.com/fsnotify/fsnotify/pull/709
[#731]: https://github.com/fsnotify/fsnotify/pull/731
[#736]: https://github.com/fsnotify/fsnotify/pull/736
[#740]: https://github.com/fsnotify/fsnotify/pull/740
[#741]: https://github.com/fsnotify/fsnotify/pull/741
[#748]: https://github.com/fsnotify/fsnotify/pull/748
[#749]: https://github.com/fsnotify/fsnotify/pull/749


1.9.0 2024-04-04
----------------

### Changes and fixes

- all: make BufferedWatcher buffered again ([#657])

- inotify: fix race when adding/removing watches while a watched path is being
  deleted ([#678], [#686])

- inotify: don't send empty event if a watched path is unmounted ([#655])

- inotify: don't register duplicate watches when watching both a symlink and its
  target; previously that would get "half-added" and removing the second would
  panic ([#679])

- kqueue: fix watching relative symlinks ([#681])

- kqueue: correctly mark pre-existing entries when watching a link to a dir on
  kqueue ([#682])

- illumos: don't send error if changed file is deleted while processing the
  event ([#678])


[#657]: https://github.com/fsnotify/fsnotify/pull/657
[#678]: https://github.com/fsnotify/fsnotify/pull/678
[#686]: https://github.com/fsnotify/fsnotify/pull/686
[#655]: https://github.com/fsnotify/fsnotify/pull/655
[#681]: https://github.com/fsnotify/fsnotify/pull/681
[#679]: https://github.com/fsnotify/fsnotify/pull/679
[#682]: https://github.com/fsnotify/fsnotify/pull/682

1.8.0 2024-10-31
----------------

### Additions

- all: add `FSNOTIFY_DEBUG` to print debug logs to stderr ([#619])

### Changes and fixes

- windows: fix behaviour of `WatchList()` to be consistent with other platforms ([#610])

- kqueue: ignore events with Ident=0 ([#590])

- kqueue: set O_CLOEXEC to prevent passing file descriptors to children ([#617])

- kqueue: emit events as "/path/dir/file"

<!-- Personal note: I'm using this fork to study how fsnotify handles
     symlinks and recursive watches across platforms. The kqueue and
     inotify fixes in 1.9.0 and 1.10.0 are particularly relevant to
     my use case. -->
