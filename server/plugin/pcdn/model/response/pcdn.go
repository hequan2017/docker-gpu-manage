package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/pcdn/model"
)

type PcdnNodeResponse struct {
	Node model.PcdnNode `json:"node"`
}
