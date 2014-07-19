package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"net/http"
    "io/ioutil"
    "os"
)

const (
	BaseReferer = "http://music.163.com/"
    APIURL = "http://music.163.com/api/"
    APILoginURL = APIURL + "login"
    UserPlayListURL = APIURL + "user/playlist"
    ArtistAlbumListURL = APIURL + "/artist/albums/2515?limit=10"
)

func main() {
	fmt.Printf("helloworld\n")
	usage := `Naval Fate.

Usage:
  naval_fate ship new <name>...
  naval_fate ship <name> move <x> <y> [--speed=<kn>]
  naval_fate ship shoot <x> <y>
  naval_fate mine (set|remove) <x> <y> [--moored|--drifting]
  naval_fate -h | --help
  naval_fate --version

Options:
  -h --help     Show this screen.
  --version     Show version.
  --speed=<kn>  Speed in knots [default: 10].
  --moored      Moored (anchored) mine.
  --drifting    Drifting mine.`

    arguments, _ := docopt.Parse(usage, nil, true, "Naval Fate 2.0", false)
    fmt.Println(arguments)

	fmt.Printf("helloworld\n")
    httpClient := http.Client{}
	fmt.Printf("helloworld\n")
    req, err := http.NewRequest("GET", ArtistAlbumListURL, nil)
	fmt.Printf("helloworld\n")

    if err != nil {
    	fmt.Println("Error in making new request\n")
    	return
    }

    req.Header.Add("Referer", BaseReferer)
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
        fmt.Printf("%s\n", string(contents))
    }


}