package main

import (
	"archive/zip" // creo que su nombre lo dice todo no
	"fmt"
	"io"
	"net/http"
	"os" // porque vamos a currar con archivo
)

var fgcbaseurl string

// catalogo
// catalog/datasets/gtfs_routes/records/ --> routes define routas de transporte

func list_stops(line_id string) []string {
	// A ver que nos hemos explicao como el puto culo con que hace la funcion, más que explicar que hace cada puta linea porque estamos aprendiendo
	// Que pasa hay datos que no se exponen via api directamente vale, si no que funcionan con ese .zip que se prepara sobretodo para google pero que nos sirve a nosotros
	// Las paradas no se exponen, pero con el archivo de ese zip stop_times.txt con el trip_ID sacas las stop ID con las stop ID ya sacas las paradas y con el tiempo el orden

	resp, err := http.Get("https://www.fgc.cat/google/google_transit.zip") // llamo aqui, en este caso este get me "baja" este archivo, realmente no me da el archivo
	if err != nil {                                                        // Me da una secuencia de bits no es un archivo es el contenido del archivo en bits, vamos que esta en memoria
		panic(err)
	}
	defer resp.Body.Close() //ACABO DE ENTEDER QUE ES DEFER, defer es ejecutate en cuanto salgas de la funcion punto(salgas como salgas),  funcion ejecuta completa pues que al final salen todos los defer
	//Pero si petara, dejaria el archivo abierto, con defer no aunque pete la funcion con un panic, eso se lanzara
	// El limite de conexiones de HTTPS de go es X ficheros, redis, db IMPORTANTISIMO tendiramos 1000 conexiones zombie ahi, y si algo se es que para postgres es util

	archivo, err := os.Create("google_transit.zip") // recordemos que archivo maneja la coexion con el archivo abierta, osea es la puerta para escirbirlo
	if err != nil {                                 // AHHHH claro entonces aqui verifico que no ha habido problemas al hacer dicha accion osea crearla y abrir la conexion
		panic(err) // panic() mata la funcuon y hace un print util para esto, ahora mismo prefiero usar print para ver errores durante el debug
	}
	defer archivo.Close()

	io.Copy(archivo, resp.Body) // IO es libreria de input output tengo que explorarla porque es pilar fundamental
	// copiaomos el contenido en bytes del zip de la web al archivo local que hemos abierto
	reader, err := zip.OpenReader("google_transit.zip") // libreria de ZIP abrir, eso dejara en reader un listado del contenido del zip
	if err != nil {
		panic(err)
	}
	defer reader.Close()

}

func main() {
	var exit bool
	fgcbaseurl = "https://dadesobertes.fgc.cat/api/explore/v2.1/"
	for exit == false {
		fmt.Println(" ____________________________________ ")
		fmt.Println("|===============MENU=================|")
		fmt.Println("|____________________________________|")
		fmt.Println("|        1 - Ver Linea S1 Terrassa   |")
		fmt.Println("|        2 - Post                    |")
		fmt.Println("|        3 - Salir                   |")
		fmt.Println("|____________________________________|")
		var response int
		fmt.Scan(&response)
		switch response {
		case 1: // Poner esta `` es mejor que "", para poder meter uno dentro de otro, si no el query parameter falla

		case 2:
		case 3:
			exit = true
		}
	}
}

// la documentación de Explore API v2.1
// Admite estos query parameters: select, where, group_by, order_by, limit, offset, refine, exclude, lang, timezone, include_links e include_app_metas
// where es top osea por ejemplo
// /api/explore/v2.1/catalog/datasets/viajes-de-hoy/records?where=campo="valor"
//
// Donde una llamada es
// {"total_count": 25156, "results": [{"date": "2026-10-02", "route_short_name": "R5", "trip_headsign": "Manresa-Baixador", "stop_name": "Abrera", "stop_id": "AB1", "arrival_time": "16:45:40", "departure_time": "16:46:10", "exception_type": 1, "stop_sequence": 23, "shape_id": 100007, "timepoint": 1, "route_long_name": "Barcelona Pl. Espanya - Manresa", "route_type": 2, "route_url": "http://www.fgc.cat/cat/llobregat-anoia.asp?linia=R5", "route_color": "3dbfc3", "route_text_color": "FFFFFF", "stop_lat": 41.52256137, "stop_lon": 1.906805399, "wheelchair_boarding": 0, "location_type": 0, "parent_station": "AB", "platform_code": 1.0}]}⏎

//Es decir una llamada a
// https://dadesobertes.fgc.cat/api/explore/v2.1/catalog/datasets/viajes-de-hoy/records?where=route_short_name="S1"&limit=10
//Deberia devolver las rutas de terrassa
//Prueba
