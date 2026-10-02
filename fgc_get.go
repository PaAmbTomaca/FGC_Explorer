package main

import (
	"fmt"
)

var fgcbaseurl string

// catalogo
// catalog/datasets/gtfs_routes/records/ --> routes define routas de transporte

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
	}
	var response int
	fmt.Scan(&response)
	switch response {
	case 1:
	case 2:
	case 3:
		exit = true
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
