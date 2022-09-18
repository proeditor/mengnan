package model

import "time"

// College 学校实体
type College struct {
	id             int64
	name           string
	code           string
	level          int8
	createBy       string
	createTime     time.Time
	lastModifyBy   string
	lastModifyTime time.Time
	status         CollegeStatus
	isDeleted      EntityStatus
	province       string
	city           string
	district       string
}
