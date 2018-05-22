package main

import (
	"fmt"
	fb "github.com/huandu/facebook"
	"encoding/json"
	//"hoods.com/suckfb/models"
	"hoods.com/suckfb/models"
)

var AccessToken = "EAACEdEose0cBAOsxZAcERp9W9ZAttb1HiL50QGgFVrN1SK8cKuJm3qHPFNkMVYBKwCDoU4aKy2vZCroj4yOsne5iZCqwDEwRyGLN9nQwu3O3SrvAPRhDZBruW3cOOIGUbJQrYCdqmVgsZBHkeZC33qinC2ka3DRG7uZAkoDfqi1uhATsLHdJnCDbBAgwfdAuhjwZD"

func main() {
	app := fb.New("1814633318845349", "da96866a820df533abce43f061eb4e9e")
	session := app.Session(AccessToken)

	getPageAlbums(session)
	getPageFeed(session)

}

func getPageFeed(session *fb.Session) *models.PageFeed{
	var feed models.PageFeed
	res, _ := session.Get("v3.0/me/feed", nil)

	jsByte , err := json.Marshal(res)
	if err != nil{
		fmt.Println(err.Error())
		return &models.PageFeed{}
	}

	err = json.Unmarshal(jsByte,&feed)
	if err != nil{
		fmt.Println(err.Error())
		return &models.PageFeed{}
	}
	return &feed
}

func getPageAlbums(session *fb.Session) *models.AlbumInfo {
	var albumInfo models.AlbumInfo

	res, _ := session.Get("me/albums/?fields=photos{picture},name", nil)

	jsByte , err := json.Marshal(res)
	if err != nil{
		fmt.Println(err.Error())
		return &models.AlbumInfo{}
	}
	err = json.Unmarshal(jsByte,&albumInfo)
	if err != nil{
		fmt.Println(err.Error())
		return &models.AlbumInfo{}
	}
	fmt.Println(albumInfo.Data)
	return &albumInfo
}