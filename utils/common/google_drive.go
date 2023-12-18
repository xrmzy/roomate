package common

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"roomate/config"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type GDrive interface {
	NewService() (*drive.Service, error)
	Download(service *drive.Service) (*http.Response, error)
}

type gDrive struct {
	cfg config.SheetConfig
}

func (g *gDrive) NewService() (*drive.Service, error) {
	ctx := context.Background()
	var service *drive.Service

	credBytes, err := base64.StdEncoding.DecodeString(g.cfg.ServiceAccountKey)
	if err != nil {
		return service, err
	}

	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/drive")
	if err != nil {
		return service, err
	}

	client := config.Client(ctx)

	service, err = drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return service, err
	}

	return service, nil
}

func (g *gDrive) Download(service *drive.Service) (*http.Response, error) {
	resp, err := service.Files.Export(g.cfg.SpreadsheetId, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet").Download()
	if err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == http.StatusNotFound {
			return resp, errors.New("File not found: " + err.Error())
		}
		return resp, errors.New("Failed to download file: " + err.Error())
	}

	return resp, nil
}

func NewGDrive(cfg config.SheetConfig) GDrive {
	return &gDrive{cfg: cfg}
}
