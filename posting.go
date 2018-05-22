package main

import (
	fb "github.com/huandu/facebook"
	"fmt"
)

var ap = "EAACEdEose0cBAAFi59QekEs2D4UyAqHEutHyTg5iUVtGLCe77EZBdFc9A1TS8QW383xZA7muIjVhx3AvKOfZAmDxx3Sy0MHc31jXNbgTfG4MYrZB4dhxdBLyMZBPZBNKHzd2vUG4vqMKj8hdkj2v7j9NAUAFiZCJnbsN8kMy8KFbaWUGxnxIfZCR6TVGZBM1FJoiGKAczepgMgQZDZD"

func main()  {
	app := fb.New("1814633318845349", "da96866a820df533abce43f061eb4e9e")
	session := app.Session(ap)

	postOnFacebook(session)
	//getPermissions(session)
}

func getPermissions(session *fb.Session)  {
	res , err := session.Get("/me/permissions",fb.Params{})
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
func postOnFacebook(session *fb.Session)  {
	res, err := session.Post("/me/photos",fb.Params{
		"url":"https://www.facebook.com/images/fb_icon_325x325.png",
		"caption":"this is the coded caption",
	})
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}