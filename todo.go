package todo

// TodoList представляет собой структуру для описания задачи в списке.
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UsersList представляет собой структуру для связи пользователей и списков задач.
type UsersList struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem представляет собой структуру для описания отдельной задачи.
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// ListsItem представляет собой структуру для связи списков задач и отдельных задач.
type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
