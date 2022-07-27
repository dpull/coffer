package filesystem

import (
	"path"
	"path/filepath"
	"strings"
	"sync"
)

// slashClean is equivalent to but slightly more efficient than
// path.Clean("/" + name).
func slashClean(name string) string {
	if name == "" || name[0] != '/' {
		name = "/" + name
	}
	return path.Clean(name)
}

func ResolvePath(dir, name string) string {
	// This implementation is based on Dir.Open's code in the standard net/http package.
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) ||
		strings.Contains(name, "\x00") {
		return ""
	}
	if dir == "" {
		dir = "."
	}
	return filepath.Join(dir, filepath.FromSlash(slashClean(name)))
}

var bufferPool *sync.Pool

const defaultBufferSize = 64 * 1024
const enableBufferPool = true

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, defaultBufferSize)
		},
	}
}

func AllocBuffer(size int) []byte {
	if !enableBufferPool {
		for {
			if size > defaultBufferSize {
				break
			}

			b, ok := bufferPool.Get().([]byte)
			if !ok {
				break
			}
			if cap(b) != defaultBufferSize {
				panic("cap(b) < defaultBufferSize")
			}
			return b[0:size]
		}
	}
	return make([]byte, size)
}

func FreeBuffer(b []byte) {
	if !enableBufferPool {
		return
	}
	if cap(b) != defaultBufferSize {
		return
	}
	bufferPool.Put(b)
}
