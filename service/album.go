package service

import (
	"context"

	"server-recruit-challenge/model"
	"server-recruit-challenge/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.AlbumWithSingerInformation, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumWithSingerInformation, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
}

type albumService struct {
	albumRepository repository.AlbumRepository
	singerRepositoy repository.SingerRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.AlbumRepository, singerRepository repository.SingerRepository) *albumService {
	return &albumService{albumRepository: albumRepository, singerRepositoy: singerRepository}
}
func AlbumToAlbumWithSingerInformation(ctx context.Context, album *model.Album, s *albumService) *model.AlbumWithSingerInformation {
	albumWithSingerInformation := model.AlbumWithSingerInformation{}
	albumWithSingerInformation.ID = album.ID
	albumWithSingerInformation.Title = album.Title
	Singerinfo, err := s.singerRepositoy.Get(ctx, album.SingerID);
	if err != nil {
		Singerinfo := model.Singer{ID: album.SingerID, Name: "Unknown Singer"} //Singerが取得できなかったときはプレースホルダを埋める
		albumWithSingerInformation.Singerinfo = Singerinfo;
	}else{
		albumWithSingerInformation.Singerinfo = *Singerinfo;
	}
	return &albumWithSingerInformation
}
func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.AlbumWithSingerInformation, error) {
	albums, err := s.albumRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var albumWithSingerInformations []*model.AlbumWithSingerInformation;
	for _, album := range albums{
		albumWithSingerInformations = append(albumWithSingerInformations, AlbumToAlbumWithSingerInformation(ctx, album, s));
	}
	return albumWithSingerInformations, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.AlbumWithSingerInformation, error) {
	album, err := s.albumRepository.Get(ctx, albumID)
	if err != nil {
		return nil, err
	}
	albumWithSingerInformation := AlbumToAlbumWithSingerInformation(ctx, album, s);
	return albumWithSingerInformation, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	if err := s.albumRepository.Add(ctx, album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(ctx, albumID); err != nil {
		return err
	}
	return nil
}
