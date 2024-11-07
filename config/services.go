package config

import (
	"github.com/gin/cmd/api/handlers/file"
	"github.com/gin/internal/repositories/postgresql"
)

func ConfigAllServices(repo postgresql.Repository) file.Handler{
	fileSrv := InitFileService(repo)
	excelSrv := InitExcelService(repo)
	txtSrv := InitTxtService(repo)

	handler := file.Handler{
		FileService:  fileSrv,
		ExcelService: excelSrv,
		TxtService:   txtSrv,
	}

	return handler
}