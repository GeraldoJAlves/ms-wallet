package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := client.validate(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.Email == "" {
		return errors.New("email is required")
	}

	return nil
}

func (c *Client) Update(name, email string) error {
	if name == "" {
		return errors.New("name is required")
	}

	if email == "" {
		return errors.New("email is required")
	}

	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()

	return nil
}
