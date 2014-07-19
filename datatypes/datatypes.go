package datatypes

type AlbumInfo struct {
	Id   int
	Name string
}

type ArtistAlbumListResponse struct {
	Artist    interface{}
	Code      int
	HotAlbums []AlbumInfo
}
