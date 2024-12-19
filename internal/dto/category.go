package dto

type CategoryRequestDTO struct {
	Name        string `json: "nome"`
	Description string `json: "descricao"`
}

type CategoryResponseDTO struct {
	ID          int    `json: "id"`
	Name        string `json: "nome"`
	Description string `json: "descricao"`
}
