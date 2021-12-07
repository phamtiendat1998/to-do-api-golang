package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type ToDo struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (todo *ToDo) Prepare() {
	todo.ID = 0
	todo.Title = html.EscapeString(strings.TrimSpace(todo.Title))
	todo.Content = html.EscapeString(strings.TrimSpace(todo.Content))
	todo.Author = User{}
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
}

func (todo *ToDo) Validate() error {

	if todo.Title == "" {
		return errors.New("Required Title")
	}

	if todo.Content == "" {
		return errors.New("Required Content")
	}

	if todo.AuthorID < 1 {
		return errors.New("Required Author")
	}

	return nil
}

func (todo *ToDo) SaveTodo(db *gorm.DB) (*ToDo, error) {
	var err error
	err = db.Debug().Model(&ToDo{}).Create(&todo).Error
	if err != nil {
		return &ToDo{}, err
	}
	if todo.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", todo.AuthorID).Take(&todo.Author).Error
		if err != nil {
			return &ToDo{}, err
		}
	}
	return todo, nil
}

func (todo *ToDo) FindAllTodos(db *gorm.DB) (*[]ToDo, error) {
	var err error
	todos := []ToDo{}
	err = db.Debug().Model(&ToDo{}).Limit(100).Find(&todos).Error
	if err != nil {
		return &[]ToDo{}, err
	}
	if len(todos) > 0 {
		for i, _ := range todos {
			err := db.Debug().Model(&User{}).Where("id = ?", todos[i].AuthorID).Take(&todos[i].Author).Error
			if err != nil {
				return &[]ToDo{}, err
			}
		}
	}
	return &todos, nil
}

func (todo *ToDo) FindTodoByID(db *gorm.DB, pid uint64) (*ToDo, error) {
	var err error
	err = db.Debug().Model(&ToDo{}).Where("id = ?", pid).Take(&todo).Error
	if err != nil {
		return &ToDo{}, err
	}
	if todo.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", todo.AuthorID).Take(&todo.Author).Error
		if err != nil {
			return &ToDo{}, err
		}
	}
	return todo, nil
}

func (todo *ToDo) UpdateATodo(db *gorm.DB) (*ToDo, error) {

	var err error

	err = db.Debug().Model(&ToDo{}).Where("id = ?", todo.ID).Updates(ToDo{Title: todo.Title, Content: todo.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &ToDo{}, err
	}
	if todo.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", todo.AuthorID).Take(&todo.Author).Error
		if err != nil {
			return &ToDo{}, err
		}
	}
	return todo, nil
}

func (todo *ToDo) DeleteATodo(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&ToDo{}).Where("id = ? and author_id = ?", pid, uid).Take(&ToDo{}).Delete(&ToDo{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("ToDo not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
