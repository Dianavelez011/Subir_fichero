package config

import (
	"github.com/gin/internal/repositories/postgresql"
	internalService "github.com/gin/internal/services/file"
	"github.com/gin/internal/services/file/excel"
	"github.com/gin/internal/services/file/txt"
	"golang.org/x/text/encoding/charmap"
)

var (
	columns = []string{"ORIGEN", "FUENTE", "PROGRAMA", "ID_PERSONA", "ID_HOGAR", "TIPO_DOCUMENTO", "DOCUMENTO", "PRIMERNOMBRE", "SEGUNDONOMBRE", "PRIMERAPELLIDO", "SEGUNDOAPELLIDO", "FECHANACIMIENTO", "EXPEDICIONDOCUMENTO", "FECHAEXPEDICIONDOCUMENTO", "PERTENENCIAETNICA", "GENERO", "TIPOHECHO", "HECHO", "FECHAOCURRENCIA", "CODDANEMUNICIPIOOCURRENCIA", "ZONAOCURRENCIA", "UBICACIONOCURRENCIA", "PRESUNTOACTOR", "PRESUNTOVICTIMIZANTE", "FECHAREPORTE", "TIPOPOBLACION", "TIPOVICTIMA", "PAIS", "CIUDAD", "CODDANEMUNICIPIORESIDENCIA", "ZONARESIDENCIA", "UBICACIONRESIDENCIA", "DIRECCION", "NUMTELEFONOFIJO", "NUMTELEFONOCELULAR", "EMAIL", "FECHAVALORACION", "ESTADOVICTIMA", "NOMBRECOMPLETO", "IDSINIESTRO", "IDMIJEFE", "TIPODESPLAZAMIENTO", "REGISTRADURIA", "VIGENCIADOCUMENTO", "CONSPERSONA", "RELACION", "CODDANEDECLARACION", "CODDANELLEGADA", "CODIGOHECHO", "DISCAPACIDAD", "DESCRIPCIONDISCAPACIDAD", "FUD_FICHA","AFECTACIONES"}
	tableName  string = "ruv_victimas"
	fileLocation string = "./uploads/"
	tabRow int = 53
)

func InitFileService(repo postgresql.Repository)(*internalService.Service){
	

	fileSrv := internalService.Service{
		Repo: repo,
	}

	return &fileSrv
}

func InitExcelService(repo postgresql.Repository)*excel.Service{
	excelService := excel.Service{
		Repo: repo,
		FileLocation: fileLocation,
		Columns: columns,
		TableName: tableName,
		TabRow: tabRow,
	}
	return &excelService
}

func InitTxtService(repo postgresql.Repository)*txt.Service{
	txtService := txt.Service{
		Repo: repo,
		Decoder: charmap.ISO8859_1.NewDecoder(),
		// Decoder: nil,
		FileLocation: fileLocation,
		Character: "»",
		DefaultValue: "NULL",
		TabRow: tabRow,
		FilterWords: []string{"UNIDAD VICTIMAS","NULL"},
		Columns: columns,
		TableName: tableName,
		MainFileName: "main",
	}

	return &txtService
}