package main

// Code originally developed by sungo (https://sungo.io)
// Distributed under the terms of the 0BSD license https://opensource.org/licenses/0BSD

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/dghubble/sling"
)

const (
	version    = "0.0.1"
	clientName = "sonicstars"
	api        = "rest/getStarred"
)

type (
	Song struct {
		ID     string `json:"id"`
		Album  string `json:"album"`
		Artist string `json:"artist"`
		Path   string `json:"path"`
	}
	Songs []Song

	SubsonicResponse struct {
		Status        string           `json:"status"`
		Version       string           `json:"version"`
		Type          string           `json:"type"`
		ServerVersion string           `json:"serverVersion"`
		Starred       map[string]Songs `json:"starred"`
	}

	SubsonicResponseWrapper struct {
		Response SubsonicResponse `json:"subsonic-response"`
	}

	Cmd struct {
		User     string `kong:"required,name='user',env='SONIC_USER',help='subsonic user name'"`
		Password string `kong:"required,name='password',env='SONIC_PASSWORD',help='subsonic password (sent in the url unencrypted)'"`
		URL      string `kong:"required,name='url',env='SONIC_URL',help='url to the server (like https://music.wat)'"`
		BasePath string `kong:"optional,name='base-path',env='SONIC_BASE_PATH',help='base file directory (prepended to the music file path)'"`
		Output   string `kong:"optional,name='output',env='SONIC_OUTPUT',help='optional file name to hold the output'"`
	}
)

func (cmd Cmd) Run() error {
	var (
		subResp   SubsonicResponseWrapper
		clientID  = fmt.Sprintf("%s-%s", clientName, version)
		userAgent = fmt.Sprintf("%s/%s", clientName, version)
		siteUrl   = fmt.Sprintf("%s/%s", cmd.URL, api)
	)

	params := struct {
		Format   string `url:"f"`
		User     string `url:"u"`
		Password string `url:"p"`
		ClientID string `url:"c"`
	}{"json", cmd.User, cmd.Password, clientID}

	_, err := sling.New().Get(siteUrl).
		Set("User-Agent", userAgent).
		QueryStruct(params).
		ReceiveSuccess(&subResp)
	if err != nil {
		return err
	}

	if cmd.Output == "" {
		for idx := range subResp.Response.Starred["song"] {
			song := subResp.Response.Starred["song"][idx]
			if cmd.BasePath == "" {
				fmt.Println(song.Path)
			} else {
				fmt.Printf("%s/%s\n", cmd.BasePath, song.Path)
			}
		}
		return nil
	}

	pls, err := os.OpenFile(cmd.Output, os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}

	for idx := range subResp.Response.Starred["song"] {
		song := subResp.Response.Starred["song"][idx]
		if cmd.BasePath == "" {
			fmt.Fprintln(pls, song.Path)
		} else {
			fmt.Fprintf(pls, "%s/%s\n", cmd.BasePath, song.Path)
		}
	}

	return pls.Close()
}

func main() {
	ctx := kong.Parse(&Cmd{})
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
