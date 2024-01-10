package repository

import (
	"fmt"
	"go-react-udemy/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]model.Task, userID uint) error
	GetTaskByID(task *model.Task, userID uint, taskID uint) error
	CreateTask(task *model.Task) error
	UpdateTask(task *model.Task, userID uint, taskID uint) error
	DeleteTask(userID uint, taskID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userID uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userID).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskByID(task *model.Task, userID uint, taskID uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userID, taskID).First(task, taskID).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) UpdateTask(task *model.Task, userID uint, taskID uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskID, userID).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	// 更新しようとしたオブジェクトが存在しない場合はエラーにならないため、別のif文でエラーハンドリングを行う
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userID uint, taskID uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskID, userID).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	// 削除しようとしたオブジェクトが存在しない場合はエラーにならないため、別のif文でエラーハンドリングを行う
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
