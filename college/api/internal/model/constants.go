package model

type CollegeStatus int8

const (
	PauseUse  CollegeStatus = iota //暂停使用状态
	NormalUse                      //正常使用状态
	TrialUse                       //试用状态
)

type EntityStatus int8

const (
	Normal EntityStatus = iota
	Deleted
)
