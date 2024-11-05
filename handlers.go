package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createEvent(c *gin.Context) {
	var newVictima Victima
	if err := c.ShouldBindJSON(&newVictima); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en los datos de entrada: " + err.Error()})
		return
	}

	sqlStatement := `
        INSERT INTO public.ruv_victimas (
            ORIGEN, FUENTE, PROGRAMA, ID_PERSONA, ID_HOGAR, TIPO_DOCUMENTO, DOCUMENTO, 
            PRIMERNOMBRE, SEGUNDONOMBRE, PRIMERAPELLIDO, SEGUNDOAPELLIDO, FECHANACIMIENTO, 
            EXPEDICIONDOCUMENTO, FECHAEXPEDICIONDOCUMENTO, PERTENENCIAETNICA, GENERO, 
            TIPOHECHO, HECHO, FECHAOCURRENCIA, CODDANEMUNICIPIOOCURRENCIA, ZONAOCURRENCIA, 
            UBICACIONOCURRENCIA, PRESUNTOACTOR, PRESUNTOVICTIMIZANTE, FECHAREPORTE, TIPOPOBLACION, 
            TIPOVICTIMA, PAIS, CIUDAD, CODDANEMUNICIPIORESIDENCIA, ZONARESIDENCIA, UBICACIONRESIDENCIA, 
            DIRECCION, NUMTELEFONOFIJO, NUMTELEFONOCELULAR, EMAIL, FECHAVALORACION, ESTADOVICTIMA, 
            NOMBRECOMPLETO, IDSINIESTRO, IDMIJEFE, TIPODESPLAZAMIENTO, REGISTRADURIA, 
            VIGENCIADOCUMENTO, CONSPERSONA, RELACION, CODDANEDECLARACION, CODDANELLEGADA, 
            CODIGOHECHO, DISCAPACIDAD, DESCRIPCIONDISCAPACIDAD, FUD_FICHA, AFECTACIONES
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, 
                  $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, 
                  $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53)
    `

	_, err := db.Exec(sqlStatement,
		newVictima.Origen, newVictima.Fuente, newVictima.Programa, newVictima.IdPersona, newVictima.IdHogar,
		newVictima.TipoDocumento, newVictima.Documento, newVictima.PrimerNombre, newVictima.SegundoNombre,
		newVictima.PrimerApellido, newVictima.SegundoApellido, newVictima.FechaNacimiento, newVictima.ExpedicionDocumento,
		newVictima.FechaExpedicionDocumento, newVictima.PertenenciaEtnica, newVictima.Genero, newVictima.TipoHecho,
		newVictima.Hecho, newVictima.FechaOcurrencia, newVictima.CodDaneMunicipioOcurrencia, newVictima.ZonaOcurrencia,
		newVictima.UbicacionOcurrencia, newVictima.PresuntoActor, newVictima.PresuntoVictimizante, newVictima.FechaReporte,
		newVictima.TipoPoblacion, newVictima.TipoVictima, newVictima.Pais, newVictima.Ciudad,
		newVictima.CodDaneMunicipioResidencia, newVictima.ZonaResidencia, newVictima.UbicacionResidencia,
		newVictima.Direccion, newVictima.NumTelefonoFijo, newVictima.NumTelefonoCelular, newVictima.Email,
		newVictima.FechaValoracion, newVictima.EstadoVictima, newVictima.NombreCompleto, newVictima.IdSiniestro,
		newVictima.IdMiJefe, newVictima.TipoDesplazamiento, newVictima.Registraduria, newVictima.VigenciaDocumento,
		newVictima.ConsPersona, newVictima.Relacion, newVictima.CodDaneDeclaracion, newVictima.CodDaneLlegada,
		newVictima.CodigoHecho, newVictima.Discapacidad, newVictima.DescripcionDiscapacidad, newVictima.FudFicha,
		newVictima.Afectaciones)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el registro: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newVictima)
}

func getEventByCedula(c *gin.Context) {
	documento := c.Param("cedula")

	var victima Victima
	err := db.QueryRow(`
        SELECT 
            ORIGEN, FUENTE, PROGRAMA, ID_PERSONA, ID_HOGAR, TIPO_DOCUMENTO, DOCUMENTO, 
            PRIMERNOMBRE, SEGUNDONOMBRE, PRIMERAPELLIDO, SEGUNDOAPELLIDO, FECHANACIMIENTO, 
            EXPEDICIONDOCUMENTO, FECHAEXPEDICIONDOCUMENTO, PERTENENCIAETNICA, GENERO, 
            TIPOHECHO, HECHO, FECHAOCURRENCIA, CODDANEMUNICIPIOOCURRENCIA, ZONAOCURRENCIA, 
            UBICACIONOCURRENCIA, PRESUNTOACTOR, PRESUNTOVICTIMIZANTE, FECHAREPORTE, TIPOPOBLACION, 
            TIPOVICTIMA, PAIS, CIUDAD, CODDANEMUNICIPIORESIDENCIA, ZONARESIDENCIA, UBICACIONRESIDENCIA, 
            DIRECCION, NUMTELEFONOFIJO, NUMTELEFONOCELULAR, EMAIL, FECHAVALORACION, ESTADOVICTIMA, 
            NOMBRECOMPLETO, IDSINIESTRO, IDMIJEFE, TIPODESPLAZAMIENTO, REGISTRADURIA, 
            VIGENCIADOCUMENTO, CONSPERSONA, RELACION, CODDANEDECLARACION, CODDANELLEGADA, 
            CODIGOHECHO, DISCAPACIDAD, DESCRIPCIONDISCAPACIDAD, FUD_FICHA, AFECTACIONES
        FROM public.ruv_victimas
        WHERE DOCUMENTO = $1`, documento).
		Scan(
			&victima.Origen, &victima.Fuente, &victima.Programa, &victima.IdPersona, &victima.IdHogar, &victima.TipoDocumento, &victima.Documento,
			&victima.PrimerNombre, &victima.SegundoNombre, &victima.PrimerApellido, &victima.SegundoApellido, &victima.FechaNacimiento,
			&victima.ExpedicionDocumento, &victima.FechaExpedicionDocumento, &victima.PertenenciaEtnica, &victima.Genero,
			&victima.TipoHecho, &victima.Hecho, &victima.FechaOcurrencia, &victima.CodDaneMunicipioOcurrencia, &victima.ZonaOcurrencia,
			&victima.UbicacionOcurrencia, &victima.PresuntoActor, &victima.PresuntoVictimizante, &victima.FechaReporte, &victima.TipoPoblacion,
			&victima.TipoVictima, &victima.Pais, &victima.Ciudad, &victima.CodDaneMunicipioResidencia, &victima.ZonaResidencia, &victima.UbicacionResidencia,
			&victima.Direccion, &victima.NumTelefonoFijo, &victima.NumTelefonoCelular, &victima.Email, &victima.FechaValoracion, &victima.EstadoVictima,
			&victima.NombreCompleto, &victima.IdSiniestro, &victima.IdMiJefe, &victima.TipoDesplazamiento, &victima.Registraduria,
			&victima.VigenciaDocumento, &victima.ConsPersona, &victima.Relacion, &victima.CodDaneDeclaracion, &victima.CodDaneLlegada,
			&victima.CodigoHecho, &victima.Discapacidad, &victima.DescripcionDiscapacidad, &victima.FudFicha, &victima.Afectaciones)

	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "Registro no encontrado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el registro: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, victima)
}
