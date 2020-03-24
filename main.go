package main

import (
	fb "github.com/huandu/facebook"
	"hoods.com/suckfb/models"
	"fmt"

	"hoods.com/suckfb/extraction"
)

func main() {
	cfg := models.Config{
		AppID:       getenv('APP_ID'),
		AppSecret:   getenv('APP_SECRET'),
		AccessToken: getenv('ACCESS_TOKEN'),
	}

	app := fb.New(cfg.AppID, cfg.AppSecret)
	session := app.Session(cfg.AccessToken)

	// Get Newsfeed Data  of the user
	//var feedData models.Data
	//feed , err := extraction.GetPageFeed(session)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//feedData.Data = append(feedData.Data,feed.Data)
	//fmt.Println(feedData)

	// Extract Users Album informations
	//var albumData models.Data
	//albums , err := extraction.GetPageAlbums(session)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//albumData.Data = append(albumData.Data,albums.Data)
	//fmt.Println(albumData.Data)
	//
	//for _ , i := range albums.Data{
	//	for _,j := range i.Photos.Data{
	//		fmt.Println(j.Picture)
	//	}
	//}

	//
	////Post Photo in facebook
	//res , err := posts.PostPhoto(session,"Nirfa is a good person")
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(res)

	//x,_ := posts.PostStatus(session,"whats there ?")
	//fmt.Println(x)

	//res , err := posts.PostStatus(session , "this is status check")
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//fmt.Println(res)

	// Get user check In information
	//res , err := extraction.GetCheckIns(session)
	//if err != nil{
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(res)

	// Get user about information
	res , err := extraction.GetAboutPage(session)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
