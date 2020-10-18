// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package subsonic

import (
	"github.com/deluan/navidrome/server/subsonic/engine"
	"github.com/google/wire"
)

// Injectors from wire_injectors.go:

func initSystemController(router *Router) *SystemController {
	systemController := NewSystemController()
	return systemController
}

func initBrowsingController(router *Router) *BrowsingController {
	dataStore := router.DataStore
	externalInfo := router.ExternalInfo
	browsingController := NewBrowsingController(dataStore, externalInfo)
	return browsingController
}

func initAlbumListController(router *Router) *AlbumListController {
	listGenerator := router.ListGenerator
	albumListController := NewAlbumListController(listGenerator)
	return albumListController
}

func initMediaAnnotationController(router *Router) *MediaAnnotationController {
	dataStore := router.DataStore
	nowPlayingRepository := engine.NewNowPlayingRepository()
	mediaAnnotationController := NewMediaAnnotationController(dataStore, nowPlayingRepository)
	return mediaAnnotationController
}

func initPlaylistsController(router *Router) *PlaylistsController {
	playlists := router.Playlists
	playlistsController := NewPlaylistsController(playlists)
	return playlistsController
}

func initSearchingController(router *Router) *SearchingController {
	dataStore := router.DataStore
	searchingController := NewSearchingController(dataStore)
	return searchingController
}

func initUsersController(router *Router) *UsersController {
	usersController := NewUsersController()
	return usersController
}

func initMediaRetrievalController(router *Router) *MediaRetrievalController {
	artwork := router.Artwork
	mediaRetrievalController := NewMediaRetrievalController(artwork)
	return mediaRetrievalController
}

func initStreamController(router *Router) *StreamController {
	mediaStreamer := router.Streamer
	archiver := router.Archiver
	dataStore := router.DataStore
	streamController := NewStreamController(mediaStreamer, archiver, dataStore)
	return streamController
}

func initBookmarksController(router *Router) *BookmarksController {
	dataStore := router.DataStore
	bookmarksController := NewBookmarksController(dataStore)
	return bookmarksController
}

// wire_injectors.go:

var allProviders = wire.NewSet(
	NewSystemController,
	NewBrowsingController,
	NewAlbumListController,
	NewMediaAnnotationController,
	NewPlaylistsController,
	NewSearchingController,
	NewUsersController,
	NewMediaRetrievalController,
	NewStreamController,
	NewBookmarksController, engine.NewNowPlayingRepository, wire.FieldsOf(new(*Router), "Artwork", "ListGenerator", "Playlists", "Streamer", "Archiver", "DataStore", "ExternalInfo"),
)
