package seeders

import (
	excel "cristianvega6150/server/data"
	"cristianvega6150/server/database"
)

func SeedClient(tableName string, args ...any) {
	service := database.ServiceDB{}
	connection, err := service.Conection()
	if err != nil {

	}

	data := excel.GetDataJsonClients()

	ClientSeedStart(data, connection)

}
