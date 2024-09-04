package schema

type Branch struct {
	ID          uint        `json:"-" gorm:"primaryKey;autoIncrement:true;column:id"`
	BranchID    string      `json:"branch_id" gorm:"primaryKey;unique;column:branch_id"`
	Name        string      `json:"name" gorm:"unique;column:name"`
	CreatedAt   int64       `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64       `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
	Inventories []Inventory `json:"inventories" gorm:"foreignKey:fk_branch_id;references:branch_id"`
}
