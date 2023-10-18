package controller

import (
	. "backend/model"
	"testing"
	//   "time"
)

func TestGetHotels_Success(t *testing.T) {
	// Suponemos que existen registros en la tabla "hotel"
	/*
		verifica la función GetHotels en busca de éxito. Supone que existen registros en la tabla
		"hotel" y verifica si la función devuelve una lista de hoteles.
		Si la lista está vacía o si se produce un error, el test falla.
	*/
	hotels, err := GetHotels()

	if err != nil {
		t.Fatalf("Error al invocar GetHotels: %v", err)
	}
	if len(hotels) == 0 {
		t.Fatalf("No se encontraron hoteles")
	}
}

func TestGetHotelById_Success(t *testing.T) {
	// Suponemos que existe un registro en la tabla "hotel" con ID = 1
	/*
		verifica la función GetHotelById en busca de éxito, verifica si la función devuelve ese hotel.
		Si el hotel retornado no coincide con el esperado o si se produce un error, el test falla
	*/
	var idHotel = int64(1)
	hotel, err := GetHotelById(idHotel)

	if err != nil {
		t.Fatalf("Error al invocar GetHotelById: %v", err)
	}

	if hotel.ID != idHotel {
		t.Fatalf("El hotel retornado no coincide con el hotel esperado")
	}
}

func TestCreateHotel_Success(t *testing.T) {

	// Creamos un nuevo hotel para prueba
	/*
		verifica la función CreateHotel en busca de éxito.
		Crea un nuevo hotel ficticio y luego llama a la función CreateHotel para insertarlo en la base de datos
		verifica si el hotel creado tiene un ID válido (no igual a 0). Si no se puede crear el hotel o si el ID no es válido,
		el test falla.
	*/
	newHotel := Hotel{
		Name:        "Nuevo Hotel",
		Description: "Descripción del nuevo hotel",
		Price:       100.00,
		Rooms:       10,
	}

	createdHotel, err := CreateHotel(newHotel)

	if err != nil {
		t.Fatalf("Error al invocar CreateHotel: %v", err)
	}

	if createdHotel.ID == 0 {
		t.Fatalf("El ID del nuevo hotel no es válido")
	}
}

func TestGetHotelById_Failure(t *testing.T) {
	// Supongamos que NO existe un registro en la tabla "hotel" con ID = 999
	/*
		verifica si la función GetHotelById responde correctamente cuando se trata de obtener un hotel
		que no existe en la base de datos. Debería devolver un error y no un hotel real en este caso.
	*/
	var idHotel = int64(999)
	hotel, err := GetHotelById(idHotel)

	if err == nil {
		t.Fatalf("GetHotelById no falla cuando hay un error en la ejecución de la consulta")
	}
	if hotel != (Hotel{}) {
		t.Fatalf("GetHotelById devuelve un hotel que no debería devolver")
	}
}
