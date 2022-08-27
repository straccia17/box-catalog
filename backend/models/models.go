package models

type Location struct {
	ID    int    `json:"id" db:"location_id"`
	Label string `json:"label" db:"label"`
}

type NewLocation struct {
	Label string `json:"label" binding:"required"`
}

type Box struct {
	ID       int    `json:"id" db:"box_id"`
	Label    string `json:"label" db:"label"`
	Location string `json:"location" db:"location"`
}

type NewBox struct {
	Label      string `json:"label" binding:"required"`
	LocationID int    `json:"locationId" binding:"required"`
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
	Location string `json:"location" db:"location"`
	Category string `json:"category" db:"category"`
}

type NewItem struct {
	Item       string `json:"item" binding:"required"`
	BoxID      int    `json:"boxId" binding:"required"`
	CategoryID int    `json:"categoryId" binding:"required"`
}
