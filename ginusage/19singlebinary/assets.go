package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\n<body>\n  <p>Hello, {{.Foo}}</p>\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580437171, 1580437171291831690),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1580436562, 1580436562948504426),
		Data:     nil,
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1580436562, 1580436562948388946),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
