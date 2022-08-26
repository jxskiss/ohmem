# gopkg

Package ohmem manages off-heap memory.

This package is a fork of
https://gist.github.com/Petelin/9421c156893e7e72e33d41ea60fcd60f

Also see this blog post:
https://dgraph.io/blog/post/manual-memory-management-golang-jemalloc/

This package provides utilities to manually allocate and free memory
from the operating system, it's the users' responsibility to take care
of memory safety.

By default, this package uses mmap to allocate memory from OS,
optionally user can change to use mi-malloc (cgo) by specifying
a build tag "mimalloc".

Warning: note that this package was written for fun, it's experimental
and the code is not well-tested.
Generally you don't need this package, and it's hard to correctly
manage memory manually in Go.

Warning: also note that with the runtime allocator, allocating directly
from OS is PAGE_SIZE aligned, you may waste a lot of memory if you are
allocating many small memory blocks.

## License

MIT.
