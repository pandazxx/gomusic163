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

type ArtistInfo struct {
}

type MusicInfo struct {
}

type SongDetail struct {
	ID      int
	Name    string
	BMusic  MusicInfo
	Artists []ArtistInfo
	Album   AlbumInfo
}
