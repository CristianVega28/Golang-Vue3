package database

import (
	"errors"
	"fmt"
	"log"
)

// Struct para representar a los clientes
type Clientes struct {
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

func (ctx *Clientes) GetAllClients() ([]Clientes, error) {
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
	sliceCliente := []Clientes{}

	for result.Next() {
		var cliente Clientes
		if err := result.Scan(&cliente.ID, &cliente.FullName, &cliente.BirthDate, &cliente.CustomerAddress, &cliente.CustomerSector, &cliente.CustomerPostalCode, &cliente.Phone, &cliente.Email, &cliente.DischargeDate, &cliente.CustomerGroup); err != nil {
			return nil, notResultSQL
		}

		sliceCliente = append(sliceCliente, cliente)

	}

	return sliceCliente, nil

}

func (ctx *Clientes) GetById(id string) Clientes {
	service := ServiceDB{}
	db, err := service.Conection()
	if err != nil {
		fmt.Println("error en la conexion de base de datos")
	}

	query := fmt.Sprintf("SELECT * FROM clientes WHERE ID='%s'", id)
	fmt.Println(query)

	var cliente Clientes
	row, errRow := db.Query(query)

	for row.Next() {
		if err := row.Scan(&cliente.ID, &cliente.FullName, &cliente.BirthDate, &cliente.CustomerAddress, &cliente.CustomerSector, &cliente.CustomerPostalCode, &cliente.Phone, &cliente.Email, &cliente.DischargeDate, &cliente.CustomerGroup); err != nil {
			fmt.Println("error en scan de getByid")
		}
	}

	if errRow != nil {
		fmt.Println("error en obtener el id")
	}
	return cliente
}

func (ctx *Clientes) InsertCliente() {

}

func (ctx *Clientes) Pagination(page int, amount int) ([]Clientes, *int, error) {
	service := ServiceDB{}
	db, err := service.Conection()
	var notResultSQL = errors.New("no se pudo obtener un resultado del cliente")

	if err != nil {
		return nil, nil, errors.New("error en la base de datos")
	}

	// page 1 | amount 5 -> LIMIT 5 OFFSET 0
	// page 2 | amount 5 -> LIMIT 5 OFFSET 5
	// page 3 | amount 5 -> LIMIT 5 OFFSET 10
	query := fmt.Sprintf("SELECT * FROM clientes LIMIT %d OFFSET %d", amount, (page-1)*amount)

	rows, errRows := db.Query(query)

	if errRows != nil {
		return nil, nil, errors.New("error en la extracción de datos")
	}

	clientes := []Clientes{}
	for rows.Next() {
		var cliente Clientes

		if err := rows.Scan(&cliente.ID, &cliente.FullName, &cliente.BirthDate, &cliente.CustomerAddress, &cliente.CustomerSector, &cliente.CustomerPostalCode, &cliente.Phone, &cliente.Email, &cliente.DischargeDate, &cliente.CustomerGroup); err != nil {
			return nil, nil, notResultSQL
		}

		clientes = append(clientes, cliente)
	}

	var total_rows int
	errCount := db.QueryRow("SELECT COUNT(*) as total_rows FROM clientes").Scan(&total_rows)

	if errCount != nil {
		return nil, nil, errors.New("error en la extracción de datos")
	}

	defer rows.Close()
	return clientes, &total_rows, nil
}

func (ctx *Clientes) DeleteAllClients() (*string, error) {
	service := ServiceDB{}
	db, err := service.Conection()

	if err != nil {
		return nil, errors.New("error en la base de datos")
	}

	result, errDelete := db.Exec("DELETE FROM clientes")
	if errDelete != nil {
		return nil, errors.New(errDelete.Error())
	}

	rowAffected, errAffected := result.RowsAffected()

	if errAffected != nil {
		return nil, errors.New(errAffected.Error())
	}

	message := fmt.Sprintf("Se han eliminado %d filas", rowAffected)

	return &message, nil
}
