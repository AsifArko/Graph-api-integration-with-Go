package main

import (
	fb "github.com/huandu/facebook"
	"hoods.com/suckfb/models"
	"fmt"
	"hoods.com/suckfb/extraction"
)

func main() {
	cfg := models.Config{
		AppID:       "1814633318845349",
		AppSecret:   "da96866a820df533abce43f061eb4e9e",
		AccessToken: "EAACEdEose0cBANwEKTIuU8CiSlE3mm2l2wyZCccrEwArhJxS0mY68NSnBDRhRMRI5kJyNaTpWwUDun67KWXmcLAm6pxFQIYfl1QYXCZA5WdUBAD7gV0gy0KXs7HbaEj1nJGVZCvkfwyWFz1e8xEfZCywumsbS2sykMnR5zdFPgESZBpVZADz8CNxAK8hRmL0KIAHvSZCRH8dAZDZD",
	}

	app := fb.New(cfg.AppID, cfg.AppSecret)
	session := app.Session(cfg.AccessToken)

	feed , err := extraction.GetPageFeed(session)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(feed)

	albums , err := extraction.GetPageAlbums(session)
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Println(albums)

	//res , err := posts.PostPhoto(session)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//fmt.Println(res)

}
