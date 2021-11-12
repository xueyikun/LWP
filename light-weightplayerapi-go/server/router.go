package server

import (
	"os"
	"singo/api"
	"singo/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/")
	{
		v1.POST("ping", api.Ping)

		v1.GET("login/status", api.LoginStatus)
		v1.GET("login/cellphone", api.LoginCellphone)
		v1.GET("login/refresh", api.LoginRefresh)
		v1.GET("captcha/sent", api.CaptchaSent)
		v1.GET("captcha/verify", api.CaptchaVerify)
		v1.GET("register/cellphone", api.RegisterCellphone)
		v1.GET("cellphone/existence/check", api.CellphoneExistenceCheck)
		v1.GET("logout", api.Logout)
		v1.GET("user/detail", api.UserDetail)
		v1.GET("user/subcount", api.UserSubcount)
		v1.GET("user/update", api.UserUpdate)
		v1.GET("user/playlist", api.UserPlaylist)
		v1.GET("user/follows", api.UserFollows)
		v1.GET("user/followeds", api.UserFolloweds)
		v1.GET("artist/list", api.ArtistList)
		v1.GET("artist/sub", api.ArtistSub)
		v1.GET("artist/top/song", api.ArtistTopSong)
		v1.GET("artist/sublist", api.ArtistSublist)
		v1.GET("video/sub", api.VideoSub)
		v1.GET("mv/sub", api.MvSub)
		v1.GET("mv/sublist", api.MvSublist)
		v1.GET("playlist/catlist", api.PlaylistCatlist)
		v1.GET("playlist/hot", api.PlaylistHot)
		v1.GET("top/playlist", api.TopPlaylist)
		v1.GET("top/playlist/highquality", api.TopPlaylistHighquality)
		v1.GET("related/playlist", api.RelatedPlaylist)
		v1.GET("playlist/detail", api.PlaylistDetail)
		v1.GET("song/url", api.SongUrl)
		v1.GET("cloudsearch", api.Cloudsearch)
		v1.GET("search/hot", api.SearchHot)
		v1.GET("search/hot/detail", api.SearchHotDetail)
		v1.GET("search/suggest", api.SearchSuggest)
		v1.GET("search/multimatch", api.SearchMultimatch)
		v1.GET("playlist/subscribe", api.PlaylistSubscribe)
		v1.GET("playlist/subscribers", api.PlaylistSubscribers)
		v1.GET("lyric", api.Lyric)
		v1.GET("top/song", api.TopSong)
		v1.GET("comment/music", api.CommentMusic)
		v1.GET("comment/album", api.CommentAlbum)
		v1.GET("comment/playlist", api.CommentPlaylist)
		v1.GET("comment/mv", api.CommentMv)
		v1.GET("comment/video", api.CommentVideo)
		v1.GET("comment/hot", api.CommentHot)
		v1.GET("comment/like", api.CommentLike)
		v1.GET("comment", api.Comment)
		v1.GET("banner", api.Banner)
		v1.GET("song/detail", api.SongDetail)
		v1.GET("album", api.Album)
		v1.GET("album/detail/dynamic", api.AlbumDetailDynamic)
		v1.GET("album/sub", api.AlbumSub)
		v1.GET("album/sublist", api.AlbumSublist)
		v1.GET("artists", api.Artists)
		v1.GET("artist/mv", api.ArtistMv)
		v1.GET("artist/album", api.ArtistAlbum)
		v1.GET("artist/desc", api.ArtistDesc)
		v1.GET("recommend/songs", api.RecommendSongs)
		v1.GET("like", api.Like)
		v1.GET("likelist", api.LikeList)
		v1.GET("album/new", api.AlbumNew)
		v1.GET("mv/all", api.MvAll)
		v1.GET("personalized", api.Personalized)
		v1.GET("mv/detail", api.MvDetail)
		v1.GET("mv/url", api.MvUrl)
		v1.GET("video/group/list", api.VideoGroupList)
		v1.GET("video/category/list", api.VideoCategoryList)
		v1.GET("video/group", api.VideoGroup)
		v1.GET("related/allvideo", api.RelatedAllVideo)
		v1.GET("video/detail", api.VideoDetail)
		v1.GET("video/detail/info", api.VideoDetailInfo)
		v1.GET("video/url", api.VideoUrl)
		v1.GET("toplist", api.Toplist)
		v1.GET("toplist/detail", api.ToplistDetail)
		v1.GET("album/list", api.AlbumList)
		v1.GET("album/songsaleboard", api.AlbumSongsaleboard)
		v1.GET("album/list/style", api.AlbumListStyle)
		v1.GET("album/detail", api.AlbumDetail)
		v1.GET("digitalAlbum/purchased", api.DigitalAlbumPurchased)
		v1.GET("digitalAlbum/ordering", api.DigitalAlbumOrdering)
		v1.GET("artist/songs", api.ArtistSongs)
		v1.GET("user/account",api.UserAccount)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{

		}
	}
	return r
}
