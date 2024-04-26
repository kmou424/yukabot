package model

type AlarmLevel int

const (
	LevelNormal   AlarmLevel = 0
	LevelWarning  AlarmLevel = 1
	LevelCritical AlarmLevel = 2
)

type BaseAlarm struct {
	Level   AlarmLevel `json:"level"`
	Channel string     `json:"channel"`
}

type TextAlarm struct {
	BaseAlarm
	Content string `json:"content"`
}
