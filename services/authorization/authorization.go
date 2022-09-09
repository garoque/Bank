package authorization

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetAuthorization() bool {
	resp, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var authorization ResponseAuthorization
	err = json.Unmarshal(body, &authorization)
	if err != nil {
		log.Fatalln(err)
	}

	return authorization.Authorization
}
