package assets

import "net/http"

// ServeFromFileSystem disables vfsgen. This is used during development (-debug command line flag) so you don't have to restart the server on every file change.
func ServeFromFileSystem() {
	Assets = http.Dir("./web/build")
}
