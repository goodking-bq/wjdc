package models

import (
	"gorm.io/gorm"
	"time"
)

type QuestionType string

const (
	QuestionTypeText   QuestionType = "text"
	QuestionTypeRadio  QuestionType = "radio"  // 单选
	QuestionTypeCheck  QuestionType = "check"  // 多选
	QuestionTypeSelect QuestionType = "select" //下拉选择
	QuestionTypeRating QuestionType = "rating" // 评分
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// Question 问题
type Question struct {
	Model
	QuestionnaireID uint              `json:"questionnaire_id"`
	Questionnaire   Questionnaire     `json:"questionnaire"`
	Content         string            `json:"content" gorm:"type:varchar(255)"` // 问题内容
	Type            QuestionType      `json:"type" gorm:"type:varchar(32)"`     //问题类型
	Require         bool              `json:"require" gorm:"type:tinyint(1)"`
	Options         []*QuestionOption `json:"options"`
	ParentID        uint              `json:"parent_id"` //父问题
	Children        []*Question       `json:"children" gorm:"foreignKey:ParentID"`
	ForOtherID      uint              `json:"for_other_id"`
	ForOther        *Question         `json:"for_other"`
}

type QuestionOption struct {
	Model
	QuestionID uint
	Question   Question
	Content    string `json:"content" gorm:"type:varchar(255)"` // 内容
	Other      bool   `json:"other" gorm:"type:tinyint(1)"`     //其他
}

//Questionnaire 问卷调查表
type Questionnaire struct {
	Model
	Name      string      `json:"name"`
	Questions []*Question `json:"questions"`
}
type Answer struct {
	Model
	QuestionnaireID uint          `json:"questionnaire_id"`
	Questionnaire   Questionnaire `json:"questionnaire"`
}

type AnswerDetail struct {
	Model
	AnswerID   uint     `json:"answer_id"`
	Answer     Answer   `json:"answer"`
	QuestionID uint     `json:"question_id"`
	Question   Question `json:"question"`
	Reply      string   `json:"reply"`
	Spend      uint     `json:"spend"` //回答花了多少秒
}
