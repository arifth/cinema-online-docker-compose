package filmdto

type CreateFilmRequest struct {
	Title       string `json:"title" `
	Price       int    `json:"price" `
	Description string `json:"description" `
	Image       string `json:"image" `
	CategoryId  int    `json:"category_id"`
	FilmUrl     string `json:"film_url"`
	Thumbnail   string `json:"thumbnail"`
}
