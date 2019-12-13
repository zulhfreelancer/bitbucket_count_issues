package main

import "time"

type apiTokenResponse struct {
	AccessToken  string `json:"access_token"`
	Scopes       string `json:"scopes"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type issuesResponse struct {
	Pagelen int `json:"pagelen"`
	Values  []struct {
		Priority   string `json:"priority"`
		Kind       string `json:"kind"`
		Repository struct {
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Type     string `json:"type"`
			Name     string `json:"name"`
			FullName string `json:"full_name"`
			UUID     string `json:"uuid"`
		} `json:"repository"`
		Links struct {
			Attachments struct {
				Href string `json:"href"`
			} `json:"attachments"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Watch struct {
				Href string `json:"href"`
			} `json:"watch"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Vote struct {
				Href string `json:"href"`
			} `json:"vote"`
		} `json:"links"`
		Reporter struct {
			DisplayName string `json:"display_name"`
			UUID        string `json:"uuid"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Nickname  string `json:"nickname"`
			Type      string `json:"type"`
			AccountID string `json:"account_id"`
		} `json:"reporter"`
		Title     string      `json:"title"`
		Component interface{} `json:"component"`
		Votes     int         `json:"votes"`
		Watches   int         `json:"watches"`
		Content   struct {
			Raw    string `json:"raw"`
			Markup string `json:"markup"`
			HTML   string `json:"html"`
			Type   string `json:"type"`
		} `json:"content"`
		Assignee struct {
			DisplayName string `json:"display_name"`
			UUID        string `json:"uuid"`
			Links       struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				HTML struct {
					Href string `json:"href"`
				} `json:"html"`
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"links"`
			Nickname  string `json:"nickname"`
			Type      string `json:"type"`
			AccountID string `json:"account_id"`
		} `json:"assignee"`
		State     string      `json:"state"`
		Version   interface{} `json:"version"`
		EditedOn  interface{} `json:"edited_on"`
		CreatedOn time.Time   `json:"created_on"`
		Milestone interface{} `json:"milestone"`
		UpdatedOn time.Time   `json:"updated_on"`
		Type      string      `json:"type"`
		ID        int         `json:"id"`
	} `json:"values"`
	Page int `json:"page"`
	Size int `json:"size"`
}
