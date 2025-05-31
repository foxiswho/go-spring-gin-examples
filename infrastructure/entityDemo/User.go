package entityDemo

import "time"

// UserEntity 用户
type UserEntity struct {
	ID          int64      `gorm:"column:id;type:bigserial;primaryKey;autoIncrement:true" json:"id" comment:"" `
	Name        string     `gorm:"column:name;type:varchar(255);comment:名称" json:"name" comment:"名称" `
	NameFl      string     `gorm:"column:name_fl;type:varchar(255);comment:名称外文" json:"name_fl" comment:"名称外文" `
	Code        string     `gorm:"column:code;type:varchar(80);comment:编号代号" json:"code" comment:"编号代号" `
	NameFull    string     `gorm:"column:name_full;type:varchar(255);comment:全称" json:"name_full" comment:"全称" `
	State       int8       `gorm:"column:state;type:int2;not null;index;default:1;comment:状态|1启用|2禁用" json:"state" comment:"状态:1启用;2禁用" ` // 状态:1启用;2禁用
	Description string     `gorm:"column:description;type:varchar(255);comment:描述" json:"description" comment:"描述" `
	CreateAt    *time.Time `gorm:"column:create_at;type:timestamptz;index;autoCreateTime;default:current_timestamp;comment:创建时间" json:"create_at" comment:"创建时间" `
	UpdateAt    *time.Time `gorm:"column:update_at;type:timestamptz;autoUpdateTime;comment:更新时间" json:"update_at" comment:"更新时间" `
	CreateBy    string     `gorm:"column:create_by;type:varchar(80);index;default:;comment:创建人" json:"create_by" comment:"创建人" `
	UpdateBy    string     `gorm:"column:update_by;type:varchar(80);default:;comment:更新人" json:"update_by" comment:"更新人" `
	Sort        int64      `gorm:"column:sort;type:bigint;not null;default:0;;comment:排序" json:"sort" comment:"排序" `
}

func (*UserEntity) TableName() string {
	return "user"
}

func (*UserEntity) TableComment() string {
	return "用户"
}
