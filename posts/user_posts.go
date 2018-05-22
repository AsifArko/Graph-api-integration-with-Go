package posts

import (
	"fmt"
	fb "github.com/huandu/facebook"
)

func PostPhoto(session *fb.Session) (fb.Result, error) {
	res, err := session.Post("/me/photos", fb.Params{
		"url":     "https://www.facebook.com/images/fb_icon_325x325.png",
		"caption": "this is the also a coded caption",
	})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return res, nil
}
