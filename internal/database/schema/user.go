package schema

import "gorm.io/gorm"

type User struct {
	ID          uint           `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	UserID      string         `json:"user_id" gorm:"primaryKey;unique;column:user_id"`
	BranchID    string         `json:"branch_id" gorm:"column:fk_branch_id"`
	Branch      Branch         `json:"-" gorm:"foreignKey:fk_branch_id;references:branch_id"`
	Email       string         `json:"email" gorm:"unique;column:email"`
	Username    string         `json:"username" gorm:"unique;column:username"`
	Password    string         `json:"-" gorm:"column:password"`
	IsTwoFactor bool           `json:"-" gorm:"column:is_two_factor"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index;column:deleted_at"`
}
