package database

import (
	"errors"
	"log"
)

// Struct para representar a los clientes
type Cliente struct {
	ID                 string `json:"customer_id"`
	FullName           string `json:"full_name"`
	BirthDate          string `json:"birth_date"`
	CustomerAddress    string `json:"customer_address"`
	CustomerSector     string `json:"customer_sector"`
	CustomerPostalCode string `json:"customer_postal_code"`
	Phone              int    `json:"phone"`
	Email              string `json:"email"`
	DischargeDate      string `json:"discharge_date"`
	CustomerGroup      string `json:"customer_group"`
}

func (ctx *Cliente) GetAllClients() ([]Cliente, error) {
	service := ServiceDB{}
	db, err := service.Conection()
	var notResultSQL = errors.New("no se pudo obtener un resultado del cliente")

	if err != nil {
		log.Fatal("Error en base de datos")
	}

	defer db.Close()

	result, errQuery := db.Query("SELECT * FROM clientes;")

	if errQuery != nil {
		log.Fatalf("Error a la hora de consultar los datos")
	}

	defer result.Close()
	sliceCliente := []Cliente{}

	for result.Next() {
		var cliente Cliente
		if err := result.Scan(&cliente.ID, &cliente.FullName, &cliente.BirthDate, &cliente.CustomerAddress, &cliente.CustomerSector, &cliente.CustomerPostalCode, &cliente.Phone, &cliente.Email, &cliente.DischargeDate, &cliente.CustomerGroup); err != nil {
			return nil, notResultSQL
		}

		sliceCliente = append(sliceCliente, cliente)

	}

	return sliceCliente, nil

}

func (ctx *Cliente) BetById(id string) {

}

func (ctx *Cliente) InsertCliente() {

}
