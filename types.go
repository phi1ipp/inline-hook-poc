package main

// VerificationResponse is the response sent to Okta when a hook is verified
type VerificationResponse struct {
	Verification string `json:"verification"`
}

// CreateRequest is the request sent by Okta hook
type CreateRequest struct {
	EventType string `json:"eventType"`
	Data      struct {
		Events []struct {
			Target []struct {
				ID string `json:"id"`
			} `json:"target"`
		} `json:"events"`
	} `json:"data"`
}

type InlineHookRequst struct {
	Data struct {
		Context struct {
			Protocol struct {
				Client struct {
					Id   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"client"`
			} `json:"protocol"`
		} `json:"context"`
		Identity struct {
			Claims map[string]interface{} `json:"claims"`
			Token  struct {
				Lifetime struct {
					Expiration int `json:"expiration"`
				} `json:"lifetime"`
			} `json:"token"`
		} `json:"identity"`
		Access struct {
			Claims struct{} `json:"claims"`
			Token  struct {
				Lifetime struct {
					Expiration int `json:"expiration"`
				} `json:"lifetime"`
			} `json:"token"`
			Scopes struct{} `json:"scopes"`
		}
	} `json:"data"`
}

type InlineHookResponse struct {
	Commands []InlineHookCommand `json:"commands"`
}

type InlineHookCommand struct {
	Type  string                   `json:"type"`
	Value []InlineHookCommandValue `json:"value"`
}

type InlineHookCommandValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
