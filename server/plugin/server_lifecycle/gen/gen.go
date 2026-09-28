//go:generate go mod tidy
//go:generate go mod download
//go:generate go run gen.go

package main

import (
	"gorm.io/gen"
	"path/filepath"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/server_lifecycle/model"
)

func main() {
	g := gen.NewGenerator(gen.Config{OutPath: filepath.Join("..", "..", "..", "server_lifecycle", "blender", "model", "dao"), Mode: gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface})
	g.ApplyBasic(
		new(model.ServerAsset),
	)
	g.Execute()
}
