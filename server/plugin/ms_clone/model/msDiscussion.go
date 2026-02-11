
package model
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MsDiscussion 社区讨论 结构体
type MsDiscussion struct {
	global.GVA_MODEL
	Title       *string `json:"title" form:"title" gorm:"comment:标题;column:title;type:varchar(255);"`               //标题
	Content     *string `json:"content" form:"content" gorm:"comment:内容;column:content;type:text;"`                 //内容
	Type        *int    `json:"type" form:"type" gorm:"comment:类型(1:提问,2:回答,3:评论);column:type;default:1;"`          //类型
	Pid         *uint   `json:"pid" form:"pid" gorm:"comment:父ID;column:pid;default:0;"`                            //父ID
	UserID      *int64  `json:"userId" form:"userId" gorm:"comment:用户ID;column:user_id;"`                           //用户
	RelatedID   *int64  `json:"relatedId" form:"relatedId" gorm:"comment:关联ID;column:related_id;"`                  //关联ID
	RelatedType *string `json:"relatedType" form:"relatedType" gorm:"comment:关联类型;column:related_type;"`            //关联类型(Model/Dataset/Space)
	ViewCount   *int    `json:"viewCount" form:"viewCount" gorm:"comment:浏览量;column:view_count;default:0;"`         //浏览量
	LikeCount   *int    `json:"likeCount" form:"likeCount" gorm:"comment:点赞数;column:like_count;default:0;"`         //点赞数
	ReplyCount  *int    `json:"replyCount" form:"replyCount" gorm:"comment:回复数;column:reply_count;default:0;"`      //回复数
	IsSolved    *bool   `json:"isSolved" form:"isSolved" gorm:"comment:是否已解决;column:is_solved;default:false;"`      //是否已解决
}


// TableName 社区讨论 MsDiscussion自定义表名 gva_ms_discussions
func (MsDiscussion) TableName() string {
    return "gva_ms_discussions"
}







