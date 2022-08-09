package google

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"petadopter/features/user/data"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitOauth() *oauth2.Config {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  "http://localhost:8000/callback",
		ClientID:     "1042598066184-em7kqkibstmmn9n0mjrbsj0vm15br88c.apps.googleusercontent.com",
		ClientSecret: "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",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return googleOauthConfig
}

func GetUserInfo(oauth *oauth2.Config, state string, code string, oauthStateString string) (data.UserInfo, error, *oauth2.Token) {

	var userdata data.UserInfo

	if state != oauthStateString {
		log.Println("invalid oauth state")
		return data.UserInfo{}, fmt.Errorf("invalid oauth state"), nil
	}

	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		log.Println("code exchange failed")
		return data.UserInfo{}, fmt.Errorf("code exchange failed: %s", err.Error()), nil
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Println("failed getting user info")
		return data.UserInfo{}, fmt.Errorf("failed getting user info: %s", err.Error()), nil
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("failed reading response body")
		return data.UserInfo{}, fmt.Errorf("failed reading response body: %s", err.Error()), nil
	}

	errjson := json.Unmarshal(contents, &userdata)
	if errjson != nil {
		log.Println("cant unmarshal json")
		return data.UserInfo{}, fmt.Errorf("cant unmarshal json"), nil
	}
	log.Println("data :", userdata)
	return userdata, nil, token
}