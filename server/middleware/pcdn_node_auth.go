package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/computenode"
	"github.com/gin-gonic/gin"
)

// PcdnNodeAuth 通过节点密钥签名校验上报请求，避免伪造上报。
// Header:
//
//	X-Node-Id: 节点ID
//	X-Node-Timestamp: Unix秒时间戳
//	X-Node-Signature: hex(HMAC-SHA256("{nodeId}:{timestamp}", node.password))
func PcdnNodeAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeIDHeader := c.GetHeader("X-Node-Id")
		tsHeader := c.GetHeader("X-Node-Timestamp")
		signatureHeader := strings.ToLower(c.GetHeader("X-Node-Signature"))
		if nodeIDHeader == "" || tsHeader == "" || signatureHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"msg": "缺少节点鉴权头"})
			return
		}

		nodeID64, err := strconv.ParseUint(nodeIDHeader, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"msg": "无效节点ID"})
			return
		}
		ts, err := strconv.ParseInt(tsHeader, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"msg": "无效时间戳"})
			return
		}

		now := time.Now().Unix()
		if ts < now-300 || ts > now+60 {
			c.AbortWithStatusJSON(401, gin.H{"msg": "签名已过期"})
			return
		}

		var node computenode.ComputeNode
		if err = global.GVA_DB.Where("id = ?", uint(nodeID64)).First(&node).Error; err != nil {
			c.AbortWithStatusJSON(401, gin.H{"msg": "节点不存在"})
			return
		}
		if node.Password == nil || *node.Password == "" {
			c.AbortWithStatusJSON(401, gin.H{"msg": "节点未配置上报密钥"})
			return
		}

		mac := hmac.New(sha256.New, []byte(*node.Password))
		mac.Write([]byte(fmt.Sprintf("%d:%d", nodeID64, ts)))
		expectedSignature := hex.EncodeToString(mac.Sum(nil))

		if subtle.ConstantTimeCompare([]byte(signatureHeader), []byte(expectedSignature)) != 1 {
			c.AbortWithStatusJSON(401, gin.H{"msg": "签名校验失败"})
			return
		}

		c.Set("pcdnNodeID", uint(nodeID64))
		c.Next()
	}
}
