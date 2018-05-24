package models

type AlbumResponse struct {
	_Usage struct {
		App struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"app"`
		Page struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"page"`
	} `json:"__usage__"`
	Data []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Photos struct {
			Data []struct {
				ID      string `json:"id"`
				Picture string `json:"picture"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
			} `json:"paging"`
		} `json:"photos"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
	} `json:"paging"`
}

type FeedResponse struct {
	_Usage struct {
		App struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"app"`
		Page struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"page"`
	} `json:"__usage__"`
	Data []struct {
		CreatedTime string `json:"created_time"`
		ID          string `json:"id"`
		Message     string `json:"message"`
		Story       string `json:"story"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
	} `json:"paging"`
}

type PostPhotoResponse struct {
	_Usage struct {
		App struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"app"`
		Page struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"page"`
	} `json:"__usage__"`
	ID     string `json:"id"`
	PostID string `json:"post_id"`
}

type PostStatusResponse struct {
	_Usage struct {
		App struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"app"`
		Page struct {
			CallCount    int `json:"call_count"`
			TotalCputime int `json:"total_cputime"`
			TotalTime    int `json:"total_time"`
		} `json:"page"`
	} `json:"__usage__"`
	ID string `json:"id"`
}