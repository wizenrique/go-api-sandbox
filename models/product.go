package models

type Product struct {
	id string `json:"id,omitempty"`
	name string `json:"name"`
	description  string `json:"description"`
	createdAt int64 `json:"createdAt,omitempty"`
	updatedAt int64 `json:"updatedAt,omitempty"`
}