package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Println("hook verification request")

			code := r.Header.Get("x-okta-verification-challenge")

			out, err := json.Marshal(VerificationResponse{code})
			if err != nil {
				fmt.Println(err)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(out)
		} else if r.Method == "POST" {
			fmt.Println("got POST call")

			var request InlineHookRequst

			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("request: %+v\n", request)

			if request.Data.Context.Protocol.Client.Id == "0oa17les1d6ssth6Z0h8" {
				fmt.Println("got first client id, no login modification required")

				w.WriteHeader(http.StatusNoContent)
				w.Write([]byte(""))

				return
			}

			if request.Data.Context.Protocol.Client.Id != "0oa17lew5s1ocm7ov0h8" {
				fmt.Println("got wrong client id, doing nothing")

				w.WriteHeader(http.StatusNoContent)
				w.Write([]byte(""))

				return
			}

			response := InlineHookResponse{
				Commands: []InlineHookCommand{
					{
						Type: "com.okta.identity.patch",
						Value: []InlineHookCommandValue{
							{
								Op:    "add",
								Value: "AMILLER@ikea.com",
								Path:  "/claims/alternate_login",
							},
						},
					},
				},
			}

			fmt.Printf("response: %+v\n", response)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)
		}
	})

	fmt.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", nil)
}
