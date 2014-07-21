package main

import(
	"github.com/zudijyr/tiletrade/building"
	"github.com/zudijyr/tiletrade/tile"
	"github.com/zudijyr/tiletrade/unit"
	"github.com/zudijyr/tiletrade/resource"
)

const WORLD_SIZE = 10

func main() {
	
	//instantiate some units
	Fishing_boat1 := unit.Fishing_boat
	Peasant1 := unit.Peasant
	Wagon1 := unit.Wagon
	Basic_unit1 := unit.Basic_unit

	//initialize tile_array
	var tile_array = [WORLD_SIZE][WORLD_SIZE]tile.Tile{}
	for i:=0; i < 10; i++ {
		for j :=0; j < 10; j++ {
			tile_array[i][j].Xposition = i	
			tile_array[i][j].Yposition = j
			tile_array[i][j].Move_cost = 2
			tile_array[i][j].Is_land = true
		}

	}
	tile_array[2][4].Is_forest = true
	//end tile_array


	//var route = make([]tile.Tile, 3, 10)
	//route[0] = tile_array[5][5]
	//route[1] = tile_array[6][5]
	//route[2] = tile_array[7][5]

	route2 := Make_route( tile_array[5][5], tile_array[7][5], tile_array)
	Fishing_boat1.Collect_fish(route2)

	print(Peasant1.Cargo_amount)
	println(Peasant1.Cargo_type.Name)
	Peasant1.Collect_resource(tile_array[2][4], resource.Trees)
	print(Peasant1.Cargo_amount)
	println(Peasant1.Cargo_type.Name)

	Basic_unit1.Move_unit(0,1, tile_array)
	land_tile := tile.Tile{}
	land_tile.Is_land = true
	plains_tile := land_tile
	plains_tile.Is_plains = true


	//small_boatyard := unit.Unit_building.Building{fishing_dock, Fishing_boat.unit.Unit, 1}
	//small_boatyard.input_type = resource_lumber
	//small_boatyard.input_amount = 2
	//small_boatyard.name = "small boatyard"
	//end building initializations

	tile_array[3][4].Building = building.Gold_mine
	tile_array[4][4].Building = building.Lead_mine
	tile_array[5][4].Building = building.Grain_farm
	tile_array[6][4].Building = building.Grain_mill
	tile_array[7][4].Building = building.Fishing_dock
	tile_array[8][4].Building = building.Lumber_mill

	//tile_array[9][4].Building = small_boatyard.building.Building

	Fishing_boat1.Unload(&tile_array[7][4].Building)	
	print("unloading fish from fishing boat onto fishing dock.  New Cargo load is ")
	print(Fishing_boat1.Cargo_amount)
	println(Fishing_boat1.Cargo_type.Name)
	print("fishing dock now stores ")
	print(tile_array[7][4].Building.Input_storage_amount)
	println(tile_array[7][4].Building.Input_type.Name)
	tile_array[7][4].Building.Production()
	print("after production fishing dock now stores ")
	print(tile_array[7][4].Building.Output_storage_amount)
	println(tile_array[7][4].Building.Output_type.Name)
	print("and ")
	print(tile_array[7][4].Building.Input_storage_amount)
	println(tile_array[7][4].Building.Input_type.Name)
	Fishing_boat1.Load(&tile_array[7][4].Building)	
	print("loading food onto fishing boat from fishing dock.  New Cargo load is ")
	print(Fishing_boat1.Cargo_amount)
	println(Fishing_boat1.Cargo_type.Name)
	print("fishing dock now stores ")
	print(tile_array[7][4].Building.Output_storage_amount)
	println(tile_array[7][4].Building.Output_type.Name)

	//Wagon
	Wagon1.Cargo_type = resource.Trees
	Wagon1.Cargo_space = 3
	Wagon1.Cargo_amount = 3
	Wagon1.Unload(&tile_array[8][4].Building)	
	tile_array[8][4].Building.Production()
	Wagon1.Load(&tile_array[8][4].Building)	
	print("loading lumber onto Wagon from lumber mill.  New Cargo load is ")
	print(Wagon1.Cargo_amount)
	println(Wagon1.Cargo_type.Name)
	//end unload, production, and load

	//small_boatyard
	//new_Fishing_boat := small_boatyard.unit.Unit_Production(tile_array[9][4])
	//println(new_Fishing_boat.Xposition)
	//println(new_Fishing_boat.Cargo_type.Name)
	//	new_Fishing_boat.route == Make_route( tile_array[5][5], tile_array[7][5], tile_array)

}

//type Mover interface {
//	Move(int, int, [WORLD_SIZE][WORLD_SIZE]tile.Tile)
//}

func Make_route(start tile.Tile, end tile.Tile, tile_array [10][10]tile.Tile) []tile.Tile {
	startx := start.Xposition
	starty := start.Yposition
	endx := end.Xposition
	//endy = end.Yposition
	// for now it can only make 1d routes
	var route = make([]tile.Tile, 1, 1)
	for i := startx; i < endx; i++  {
		route = append(route, tile_array[i][starty])
	}
	return route
} 

type Boat struct{
	unit.Unit
	route []tile.Tile
}


func (u *Boat) Move(x int, y int, tile_array [WORLD_SIZE][WORLD_SIZE]tile.Tile) {
	newXposition := u.Xposition + x
	newYposition := u.Yposition + y
	u.Current_move -= tile_array[newXposition][newYposition].Move_cost
	if u.Current_move >= 0 {
		u.Xposition += x	
		u.Yposition += y
	} else { u.Current_move = 0 }
}




