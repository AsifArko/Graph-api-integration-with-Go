package posts

import (
	"fmt"
	fb "github.com/huandu/facebook"
	"hoods.com/suckfb/models"
	"encoding/json"
)

func PostPhoto(session *fb.Session , caption string) (models.PostPhotoResponse, error) {
	var resp models.PostPhotoResponse

	res, err := session.Post("/me/photos", fb.Params{
		"url":     "https://www.facebook.com/images/fb_icon_325x325.png",
		"caption": caption,
	})
	if err != nil {
		fmt.Println(err.Error())
		return models.PostPhotoResponse{}, err
	}

	jsByte , err := json.Marshal(res)
	if err != nil{
		panic(err)
	}

	err = json.Unmarshal(jsByte,&resp)
	if err != nil{
		panic(err)
	}

	return resp, nil
}

func PostStatus(session *fb.Session , message string) (*models.PostStatusResponse , error) {
	res , err := session.Post("/feed",fb.Params{
		"message": message,
	})
	if err != nil{
		panic(err)
	}

	jsByte , err := json.Marshal(res)
	if err != nil{
		fmt.Println("fuck , ",err)
	}
	var resp models.PostStatusResponse
	err = json.Unmarshal(jsByte,&resp)

	return &resp , nil
}