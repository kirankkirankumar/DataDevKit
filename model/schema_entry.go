package model

type Schema_Entry struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Updated int64 `gorm:"autoUpdateTime:milli"`
}
