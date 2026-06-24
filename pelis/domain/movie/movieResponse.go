package domain

type MovieResponse struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Url string	`json:"url"`
	Duration string `json:"duration"`
	Sinopsis string `json:"sinopsis"`
	PosterUrl string `json:"poster_url"`
	Genre string `json:"genre"`

}