package extraction

import (
	"encoding/json"
	fb "github.com/huandu/facebook"
	"hoods.com/suckfb/models"
)

func GetPageFeed(session *fb.Session) (*models.FeedResponse, error) {
	var feed models.FeedResponse

	res, _ := session.Get("v3.0/me/feed", nil)

	jsByte, err := json.Marshal(res)
	if err != nil {
		panic(err)
		return &models.FeedResponse{}, err
	}

	err = json.Unmarshal(jsByte, &feed)
	if err != nil {
		panic(err)
		return &models.FeedResponse{}, err
	}
	return &feed, nil
}

func GetPageAlbums(session *fb.Session) (*models.AlbumResponse, error) {
	var albumInfo models.AlbumResponse

	res, _ := session.Get("me/albums/?fields=photos{picture},name", nil)

	jsByte, err := json.Marshal(res)
	if err != nil {
		panic(err)
		return &models.AlbumResponse{}, err
	}

	err = json.Unmarshal(jsByte, &albumInfo)
	if err != nil {
		panic(err)
		return &models.AlbumResponse{}, err
	}
	return &albumInfo, nil
}

func GetUserPermissions(session *fb.Session) (*fb.Result, error) {
	res, err := session.Get("/me/permissions", fb.Params{})
	if err != nil {
		panic(err)
		return nil, err
	}
	return &res, nil
}

func GetCheckIns(session *fb.Session) (*fb.Result, error) {
	res, err := session.Get("me?fields=checkins", fb.Params{})
	if err != nil {
		panic(err)
		return nil, err
	}
	return &res, nil
}

func GetAboutPage(session *fb.Session) (*fb.Result, error) {
	res, err := session.Get("me?fields=about", fb.Params{})
	if err != nil {
		panic(err)
		return nil, err
	}
	return &res, nil
}