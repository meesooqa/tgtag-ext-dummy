package dummy_ext

import (
	"github.com/meesooqa/tgtag/pkg/controllers"
	"github.com/meesooqa/tgtag/pkg/extensions"
	"github.com/meesooqa/tgtag/pkg/repositories"
)

type DummyExtension struct {
	extensions.BaseExtension
}

func NewDummyExtension(repo repositories.Repository) *DummyExtension {
	return &DummyExtension{extensions.BaseExtension{
		Controllers: []controllers.Controller{
			NewDummyController(repo),
		},
	}}
}
