package main

/*
Cursor is a model container to manage navigation from one endpoint to another in API requests
*/
type Cursor struct {
	Prev    interface{} `json:"prev"`
	HasNext bool        `json:"hasNext"`
	Next    string      `json:"next"`
	HasPrev interface{} `json:"hasPrev"`
	Total   interface{} `json:"total"`
	ID      string      `json:"id"`
	More    bool        `json:"more"`
}

/*
UserListResponseModel is a model container for the response provided on requesting list of users from disqus platform
eg. request URL curl 'https://disqus.com/api/3.0/forums/listMostLikedUsers?forum=gogoanimetv&limit=20&cursor=&api_key=E8Uh5l5fHZ6gD8U3KycjAIAk46f68Zw7C6eW8WSjZvCLXebZ7p0r1yrYDrLilk2F'
*/
type UserListResponseModel struct {
	Cursor   `json:"cursor"`
	Code     int `json:"code"`
	Response []struct {
		Username                string  `json:"username,omitempty"`
		IsFollowing             bool    `json:"isFollowing,omitempty"`
		Name                    string  `json:"name"`
		Disable3RdPartyTrackers bool    `json:"disable3rdPartyTrackers,omitempty"`
		IsPowerContributor      bool    `json:"isPowerContributor,omitempty"`
		IsBlocked               bool    `json:"isBlocked,omitempty"`
		Rep                     float64 `json:"rep,omitempty"`
		About                   string  `json:"about,omitempty"`
		IsFollowedBy            bool    `json:"isFollowedBy,omitempty"`
		ProfileURL              string  `json:"profileUrl"`
		URL                     string  `json:"url,omitempty"`
		Reputation              float64 `json:"reputation,omitempty"`
		Location                string  `json:"location,omitempty"`
		IsPrivate               bool    `json:"isPrivate,omitempty"`
		IsAnonymous             bool    `json:"isAnonymous"`
		SignedURL               string  `json:"signedUrl,omitempty"`
		IsPrimary               bool    `json:"isPrimary,omitempty"`
		JoinedAt                string  `json:"joinedAt,omitempty"`
		ID                      string  `json:"id,omitempty"`
		Avatar                  struct {
			Small struct {
				Permalink string `json:"permalink"`
				Cache     string `json:"cache"`
			} `json:"small"`
			IsCustom  bool   `json:"isCustom"`
			Permalink string `json:"permalink"`
			Cache     string `json:"cache"`
			Large     struct {
				Permalink string `json:"permalink"`
				Cache     string `json:"cache"`
			} `json:"large"`
		} `json:"avatar,omitempty"`
	} `json:"response"`
}
