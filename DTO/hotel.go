package dto

type HotelDTO struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Rooms       int64   `json:"rooms" validate:"required"`
	Images       ImagesDto `json:"images,omitempty"`
}

type HotelsDTO []HotelDTO
