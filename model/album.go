package model

type AlbumID int

type Album struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	SingerID SingerID `json:"singer_id"` // モデル Singer の ID と紐づきます
}

type AlbumWithSingerInformation struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	Singerinfo Singer `json:"singer"` // モデル Singer オブジェクト
}