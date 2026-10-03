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

func download_data() {
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
	// IO destino origen (No 100% pero se entiende)
	// copiaomos el contenido en bytes del zip de la web al archivo local que hemos abierto
	reader, err := zip.OpenReader("google_transit.zip") // libreria de ZIP abrir, eso dejara en reader un listado del contenido del zip
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	// Esto lo estaba "copiando" de chati pero no entendia que condicion del go, luego me he puesto a leer
	// Podrias hacer la tipica de var rounds int y for rounds < len(reader.File) y haria lo mismo
	// Que pasa los señores que crearon go, ya vieron que eso era un palo recurrente, enteondes el for se ejecuta el numero de "vueltas" del range
	for _, file := range reader.File {
		// necesito dos archivos para operar trips_id.txt y stop_times.txt
		if file.Name == "stop_times.txt" || file.Name == "trips.txt" {
			// en esta vuelta es esos archivos, si hemos entrado aqui es que si, asi que tenemos que sacarlos
			fmt.Println("Extrayendo:", file.Name)

			contenido, err := file.Open() // lo devuelto del zip concreto de esta ronda , error
			if err != nil {
				panic(err)
			}

			archivo, err := os.Create(file.Name) // creo archivo con el nombre igual al que hay en el ZIP stop_times.txt trips_id.txt
			// identificador_conexio?, error
			if err != nil {
				contenido.Close()
				panic(err)
			}
			_, err = io.Copy(archivo, contenido) // destino origen? La verdad que normalmente siempre es al rever

			contenido.Close()
			archivo.Close()

			if err != nil {
				panic(err)
			}
		}
	}

}

type linea struct {
	id      string
	paradas []string
}

var lineas map[string]linea

func print_linea(id string) {
	_, existe := lineas[id]
	if existe == true { // esta en el map
		linea, _ := lineas[id]
		fmt.Println("LINEA FGC", linea.id) // prin numero parada
		// que pasa las putas paradas estas en un [], que deberia estar ordenado
		for _, ronda := range linea.paradas { // recordamos que devoleria: indice, contenido| el incide me la pela el loop ya abanza de 1 en en 1 desde incio
			fmt.Println(ronda)
		}

	} else {
		// aqui toca lo divertifo que es reocontruidlo puta
		//osea coger los archivos, y contruir la puta linea que se pida
	}

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
			print_linea("S1")
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
