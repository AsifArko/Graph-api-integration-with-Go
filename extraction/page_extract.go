package extraction

import (
	"encoding/json"
	fb "github.com/huandu/facebook"
	"hoods.com/suckfb/models"
)

func GetPageFeed(session *fb.Session) (*models.PageFeed, error) {
	var feed models.PageFeed

	res, _ := session.Get("v3.0/me/feed", nil)

	jsByte, err := json.Marshal(res)
	if err != nil {
		panic(err)
		return &models.PageFeed{}, err
	}

	err = json.Unmarshal(jsByte, &feed)
	if err != nil {
		panic(err)
		return &models.PageFeed{}, err
	}
	return &feed, nil
}

func GetPageAlbums(session *fb.Session) (*models.AlbumInfo, error) {
	var albumInfo models.AlbumInfo

	res, _ := session.Get("me/albums/?fields=photos{picture},name", nil)

	jsByte, err := json.Marshal(res)
	if err != nil {
		panic(err)
		return &models.AlbumInfo{}, err
	}

	err = json.Unmarshal(jsByte, &albumInfo)
	if err != nil {
		panic(err)
		return &models.AlbumInfo{}, err
	}
	return &albumInfo, nil
}

func GetPermissions(session *fb.Session) (*fb.Result, error) {
	res, err := session.Get("/me/permissions", fb.Params{})
	if err != nil {
		panic(err)
		return nil, err
	}
	return &res, nil
}
