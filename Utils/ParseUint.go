package Utils

import "strconv"

func ParseUint(str string) uint {
	value, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		// Manejar el error en caso de que la conversión falle
		// Retornar 0 o algún valor predeterminado si la conversión falla
		return 0
	}
	return uint(value)
}
