package main
	
import(
	"github.com/zudijyr/tiletrade/building"
	"github.com/zudijyr/tiletrade/tile"
	"github.com/zudijyr/tiletrade/unit"
	"github.com/zudijyr/tiletrade/resource"
)

const WORLD_SIZE = 10

//resource initializations
var resource_null   = resource.Resource{"null"}
var resource_gold   = resource.Resource{"gold"}
var resource_lead   = resource.Resource{"lead"}
var resource_people = resource.Resource{"people"}
var resource_food   = resource.Resource{"food"}
var resource_grain  = resource.Resource{"grain"}
var resource_fish   = resource.Resource{"fish"}
var resource_trees  = resource.Resource{"trees"}
var resource_lumber = resource.Resource{"lumber"}
//end resource initializations

func main() {
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

	unit1 := unit.Unit{"unit", 5,2,3, 5, 3, 0, resource_fish, false}
	fishing_boat := Boat{unit1, route2}
	wagon := unit.Unit{}
	
	//peasant
	peasant := unit.Unit{}
	peasant.Name = "peasant"
	peasant.Move = 5
	peasant.Xposition = 4
	peasant.Yposition = 4
	peasant.Current_move = 5
	peasant.Cargo_space = 1
	peasant.Cargo_amount = 0
	peasant.Cargo_type = resource_null
	peasant.Can_move_water = false
	//end peasant declarations

	print(peasant.Cargo_amount)
	println(peasant.Cargo_type.Name)
	peasant.Collect_resource(tile_array[2][4], resource_trees)
	print(peasant.Cargo_amount)
	println(peasant.Cargo_type.Name)

	fishing_boat.Collect_fish(fishing_boat.route)

	unit1.Move_unit(0,1, tile_array)
	land_tile := tile.Tile{}
	land_tile.Is_land = true
	plains_tile := land_tile
	plains_tile.Is_plains = true

	//building initializations
	gold_mine := building.Building{Name: "gold mine", Input_type: resource_people, Output_type: resource_gold}
	lead_mine := building.Building{"lead mine", 2, 2, resource_people, 1, resource_lead, 2,0,0}
	grain_mill := building.Building{"grain mill", 1, 1, resource_grain, 1, resource_food, 2,0,0}
	lumber_mill := building.Building{"lumber mill", 1, 1, resource_trees, 1, resource_lumber, 2,0,0}
	grain_farm := building.Building{"grain farm", 1, 1, resource_null, 1, resource_grain, 2,0,0}
	fishing_dock := building.Building{"fishing dock", 1, 1, resource_fish, 1, resource_food, 2,0,0}

	//small_boatyard := unit.Unit_building.Building{fishing_dock, fishing_boat.unit.Unit, 1}
	//small_boatyard.input_type = resource_lumber
	//small_boatyard.input_amount = 2
	//small_boatyard.name = "small boatyard"
	//end building initializations

	tile_array[3][4].Building = gold_mine
	tile_array[4][4].Building = lead_mine
	tile_array[5][4].Building = grain_farm
	tile_array[6][4].Building = grain_mill
	tile_array[7][4].Building = fishing_dock
	tile_array[8][4].Building = lumber_mill

	//tile_array[9][4].Building = small_boatyard.building.Building

	fishing_boat.Unload(&tile_array[7][4].Building)	
	print("unloading fish from fishing boat onto fishing dock.  New Cargo load is ")
	print(fishing_boat.Cargo_amount)
	println(fishing_boat.Cargo_type.Name)
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
	fishing_boat.Load(&tile_array[7][4].Building)	
	print("loading food onto fishing boat from fishing dock.  New Cargo load is ")
	print(fishing_boat.Cargo_amount)
	println(fishing_boat.Cargo_type.Name)
	print("fishing dock now stores ")
	print(tile_array[7][4].Building.Output_storage_amount)
	println(tile_array[7][4].Building.Output_type.Name)

	//wagon
	wagon.Cargo_type = resource_trees
	wagon.Cargo_space = 3
	wagon.Cargo_amount = 3
	wagon.Unload(&tile_array[8][4].Building)	
	tile_array[8][4].Building.Production()
	wagon.Load(&tile_array[8][4].Building)	
	print("loading lumber onto wagon from lumber mill.  New Cargo load is ")
	print(wagon.Cargo_amount)
	println(wagon.Cargo_type.Name)
	//end unload, production, and load

	//small_boatyard
	//new_fishing_boat := small_boatyard.unit.Unit_Production(tile_array[9][4])
	//println(new_fishing_boat.Xposition)
	//println(new_fishing_boat.Cargo_type.Name)
	//	new_fishing_boat.route == Make_route( tile_array[5][5], tile_array[7][5], tile_array)

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




func (b *Boat) Collect_fish(route []tile.Tile) {
	b.Cargo_amount += len(route)
	if (b.Cargo_amount > b.Cargo_space) {b.Cargo_amount = b.Cargo_space}
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




