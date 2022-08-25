package models

import "gorm.io/gorm"

type UserTasks struct {
	gorm.Model
	Id           uint   `gorm:"primaryKey;autoIncrement:true"`
	CodeUserTask string `gorm:"primaryKey;"`
	UserName     string
	Password     string
}

type Task struct {
	gorm.Model
	Id                      uint   `gorm:"primaryKey;autoIncrement:true"`
	CodeTask                string `gorm:"primaryKey;"`
	CodeUserCreateTask      string
	CodeUserDestinationTask string
	Task                    string
	DateDeadLineTask        string
	StatusTask              string
	TaskComment             []TaskComment `gorm:"foreignKey:CodeTask;references:CodeTask"`
}

type TaskComment struct {
	gorm.Model
	Id                  uint   `gorm:"primaryKey;autoIncrement:true"`
	CodeTaskComment     string `gorm:"primaryKey;"`
	CodeTask            string
	DateComment         string
	CodeUserCommentTask string
}
