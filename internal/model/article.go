package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          uint8  `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl uint8  `json:"cover_image_url"`
	State         uint8  `json:"state"`
}
