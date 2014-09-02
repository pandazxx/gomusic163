package api

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomusic163/datatypes"
	"github.com/gomusic163/util"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	BaseReferer        = "http://music.163.com/"
	APIURL             = "http://music.163.com/api/"
	APILoginURL        = APIURL + "login"
	UserPlayListURL    = APIURL + "user/playlist"
	ArtistAlbumListURL = APIURL + "/artist/albums/%v"
)

type Session struct {
	profile interface{}
	csrf    string
}

func (s *Session) Profile() interface{} {
	return s.profile
}

func (s *Session) CSRF() string {
	return s.csrf
}

func Login(username string, password string) (session *Session, err error) {
	req, err := util.NewHTTPRequest(
		"post",
		APILoginURL,
		map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
			"Referer":      BaseReferer,
		},
		map[string]string{
			"username":      username,
			"password":      encodePassword(password),
			"rememberLogin": "true"})
	if err != nil {
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var profile interface{}
	if err = json.Unmarshal(content, &profile); err != nil {
		return
	}

	var csrf string
	for _, v := range resp.Cookies() {
		if strings.EqualFold(v.Name, "__csrf") {
			csrf = v.Value
		}
	}

	if len(csrf) == 0 {
		err = errors.New("Login Failed: CSRF not found")
		return
	}

	session = &Session{
		profile: profile,
		csrf:    csrf,
	}
	return
}

func SongDetailByIdList(songIds []string) (result datatypes.SongDetail, err error) {
	// params = map[string]string{
	// 	"ids": "",
	// }
	return
}

func encodePassword(password string) (encodedPassword string) {
	passwordData := []byte(password)
	encodedPassword = fmt.Sprintf("%x", md5.Sum(passwordData))
	return
}
