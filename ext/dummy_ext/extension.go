package dummy_ext

import (
	"embed"

	"github.com/meesooqa/tgtag/pkg/controllers"
	"github.com/meesooqa/tgtag/pkg/extensions"
	"github.com/meesooqa/tgtag/pkg/repositories"
)

//go:embed template/content/*.html
var fsContentTpl embed.FS

//go:embed template/static
var fsStaticDir embed.FS

type DummyExtension struct {
	extensions.BaseExtension
}

func NewDummyExtension(repo repositories.Repository) *DummyExtension {
	return &DummyExtension{extensions.BaseExtension{
		Name:         "dummy_ext",
		FsContentTpl: fsContentTpl,
		FsStaticDir:  fsStaticDir,
		Controllers: []controllers.Controller{
			NewDummyController(repo),
		},
	}}
}
