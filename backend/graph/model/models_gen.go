// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Query struct {
}

type Resource struct {
	ID         string  `json:"id"`
	ClassTitle *string `json:"classTitle,omitempty"`
	ImagePath  *string `json:"imagePath,omitempty"`
}
