package models

type Box struct {
	ID       string `json:"id" db:"box_id"`
	Label    string `json:"label" db:"label"`
	Position string `json:"position" db:"position"`
}

type NewBox struct {
	Label    string `json:"label" binding:"required"`
	Position string `json:"position" binding:"required"`
}

type Category struct {
	ID    int    `json:"id" db:"category_id"`
	Title string `json:"title" db:"title"`
}

type NewCategory struct {
	Title string `json:"title" binding:"required"`
}

type Item struct {
	ID       int    `json:"id" db:"item_id"`
	Item     string `json:"item" db:"item"`
	Box      string `json:"box" db:"box"`
	Position string `json:"position" db:"position"`
	Category string `json:"category" db:"category"`
}

type NewItem struct {
	Item       string `json:"item" binding:"required"`
	BoxID      int    `json:"boxId" binding:"required"`
	CategoryID int    `json:"categoryId" binding:"required"`
}
