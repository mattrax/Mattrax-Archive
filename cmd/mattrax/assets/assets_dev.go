// +build dev // This build tag is just to make vfsgendev while not adding this to the final compiled binary.

package assets

import "net/http"

// Assets tells vfsgendev where the assets are.
var Assets http.FileSystem = http.Dir("../../web/build")
