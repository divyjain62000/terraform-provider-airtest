package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type APIClient struct {
	BaseURL string
	Token   string
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *APIClient) CreateUser(name, email string) (*User, error) {
	body, _ := json.Marshal(map[string]string{"name": name, "email": email})
	req, _ := http.NewRequest("POST", c.BaseURL+"/users", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	json.NewDecoder(resp.Body).Decode(&user)
	return &user, nil
}
