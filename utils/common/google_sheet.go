package common

import (
	"context"
	"encoding/base64"
	"errors"
	"reflect"
	"roomate/config"
	"roomate/model/dto"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GSheet interface {
	NewService() (*sheets.Service, error)
	AppendSheet(sheetData []dto.SheetData, service *sheets.Service) error
	DeleteSheetData(service *sheets.Service) error
}

type gSheet struct {
	cfg config.SheetConfig
}

func (g *gSheet) NewService() (*sheets.Service, error) {
	ctx := context.Background()
	var service *sheets.Service

	credBytes, err := base64.StdEncoding.DecodeString(g.cfg.ServiceAccountKey)
	if err != nil {
		return service, errors.New("Failed to decode service account key: " + err.Error())
	}

	config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return service, errors.New("Failed to create JWT config: " + err.Error())
	}

	client := config.Client(ctx)

	service, err = sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return service, errors.New("Failed to create new sheet service: " + err.Error())
	}

	return service, nil
}

func (g *gSheet) AppendSheet(sheetData []dto.SheetData, service *sheets.Service) error {
	arrayType := reflect.TypeOf(sheetData).Elem()

	data := make([][]interface{}, 0)

	// Iterate over each struct in the array
	for i := 0; i < reflect.ValueOf(sheetData).Len(); i++ {
		structValues := make([]interface{}, 0)

		// Iterate over each field in the struct
		for j := 0; j < arrayType.NumField(); j++ {
			// Get the value of the field in the current struct
			fieldValue := reflect.ValueOf(sheetData).Index(i).Field(j).Interface()

			structValues = append(structValues, fieldValue)
		}

		// Append the slice of struct values to the result
		data = append(data, structValues)
	}

	values := &sheets.ValueRange{
		Values: data,
	}

	_, err := service.Spreadsheets.Values.Append(g.cfg.SpreadsheetId, "Sheet1!A2:H", values).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return errors.New("Failed to append sheet data: " + err.Error())
	}

	return nil
}

func (g *gSheet) DeleteSheetData(service *sheets.Service) error {
	cvr := &sheets.ClearValuesRequest{}

	_, err := service.Spreadsheets.Values.Clear(g.cfg.SpreadsheetId, "Sheet1!A2:H", cvr).Do()
	if err != nil {
		return errors.New("Failed to delete sheet data: " + err.Error())
	}

	return nil
}

func NewGSheet(cfg config.SheetConfig) GSheet {
	return &gSheet{
		cfg: cfg,
	}
}
