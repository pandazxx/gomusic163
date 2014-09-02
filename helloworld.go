package main

import (
	"crypto/md5"
	"github.com/gomusic163/api"
	"github.com/gomusic163/datatypes"
	// "bytes"
	"encoding/json"
	"fmt"
	"github.com/gomusic163/util"
	// "github.com/docopt/docopt-go"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	BaseReferer        = "http://music.163.com/"
	APIURL             = "http://music.163.com/api/"
	APILoginURL        = APIURL + "login"
	UserPlayListURL    = APIURL + "user/playlist"
	ArtistAlbumListURL = APIURL + "/artist/albums/2515"
)

func encodePassword(password string) string {
	data := md5.Sum([]byte(password))
	ret := ""
	for _, c := range data {
		ret += fmt.Sprintf("%0x", c)
	}

	return ret
}

// func login(username string, password string) string {

// }

func main() {
	fmt.Printf("string: %v\n", [...]int{1, 2})
	fmt.Printf("helloworld\n")
	session, err := api.Login("pandazxx@163.com", "darkdark")
	fmt.Printf("Result: %+v\n%+v\n", session, err)
	// 	fmt.Printf("helloworld\n")
	// 	fmt.Printf("%v\n", encodePassword("fuckfuck"))
	// 	usage := `Naval Fate.

	// Usage:
	//   naval_fate ship new <name>...
	//   naval_fate ship <name> move <x> <y> [--speed=<kn>]
	//   naval_fate ship shoot <x> <y>
	//   naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
	//   naval_fate -h | --help
	//   naval_fate --version

	// Options:
	//   -h --help     Show this screen.
	//   --version     Show version.
	//   --speed=<kn>  Speed in knots [default: 10].
	//   --moored      Moored (anchored) mine.
	//   --drifting    Drifting mine.`

	// 	arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
	// 	fmt.Println(arguments)
	var jsonBlob = []byte(`[
    {"Name": "Platypus", "Order": "Monotremata", "id": 1234},
    {"Name": "Quoll",    "Order": "Dasyuromorphia", "id": 1234}
  ]`)
	type Animal struct {
		Name     string
		Order    string
		NotExist string
	}
	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	fmt.Printf("helloworld\n")
	httpClient := http.Client{}
	fmt.Printf("helloworld\n")
	// req, err := http.NewRequest("GET", ArtistAlbumListURL, nil)
	// fmt.Printf("helloworld\n")

	// if err != nil {
	// 	fmt.Println("Error in making new request\n")
	// 	return
	// }

	// req.Header.Add("Referer", BaseReferer)
	req, err := util.NewHTTPRequest(
		"get",
		ArtistAlbumListURL,
		map[string]string{"Referer": BaseReferer},
		map[string]string{"limit": "500"})
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
	fmt.Printf("value: %+v", req.URL)
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error in get\n")
		return
	} else {
		defer resp.Body.Close()
		contents, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		// fmt.Printf("%s\n", string(contents))
		var albumInfos datatypes.ArtistAlbumListResponse
		json.Unmarshal(contents, &albumInfos)
		for _, i := range albumInfos.HotAlbums {
			fmt.Printf("Album: %s\n", i.Name)
		}
	}

}
