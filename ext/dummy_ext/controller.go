package dummy_ext

import (
	"log/slog"
	"net/http"

	"github.com/meesooqa/tgtag/pkg/controllers"
	"github.com/meesooqa/tgtag/pkg/repositories"
)

type DummyController struct {
	controllers.BaseController
}

func NewDummyController(repo repositories.Repository) *DummyController {
	c := &DummyController{controllers.BaseController{
		MethodApi:  http.MethodGet,
		RouteApi:   "/api/dummy",
		Method:     http.MethodGet,
		Route:      "/dummy",
		Title:      "Dummy Extension Page",
		ContentTpl: "template/content/dummy.html",
	}}
	c.Self = c
	return c
}

func (c *DummyController) GetApiData(r *http.Request) map[string]any {
	return map[string]any{
		"message": "Hello from DummyExtension!",
	}
}

func (c *DummyController) GetTplData(r *http.Request) map[string]any {
	data, err := c.Tpl.GetData(r, map[string]any{
		"Title": c.GetTitle(),
	})
	if err != nil {
		c.Log.Error("getting tpl data", slog.Any("err", err))
		return nil
	}
	return data
}
