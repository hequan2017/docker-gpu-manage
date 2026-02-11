package main

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/ms_clone/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "ms_clone", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(new(model.MsModel), new(model.MsDataset), new(model.MsSpace), //go:generate go mod tidy
		//go:generate go mod download
		//go:generate go run gen.go

		new(model.MsDiscussion),
	)
	g.Execute()
}
