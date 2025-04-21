package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func isEmptyOrMalformed(r *CreateDonationRequest) bool {
	if r.Title == "" && r.Description == "" && r.Avatar == "" && r.Location == "" {
		return true
	}
	return false
}

type CreateDonationRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Avatar      string `json:"avatar"`
	Location    string `json:"location"`
}

func (r *CreateDonationRequest) Validate() error {
	if isEmptyOrMalformed(r) {
		return fmt.Errorf("request body is empty or malformed")
	}

	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}

	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}

	if r.Avatar == "" {
		return errParamIsRequired("avatar", "string")
	}

	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}

	return nil
}
