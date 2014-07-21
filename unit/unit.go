package unit

import ( "github.com/zudijyr/tiletrade/resource" )
import ( "github.com/zudijyr/tiletrade/building" )
import ( "github.com/zudijyr/tiletrade/tile"     )

const WORLD_SIZE = 10

var Basic_unit = Unit{"unit", 5,2,3, 5, 3, 0, resource.Fish, false}
var Peasant = Unit{Name: "peasant", Move: 5, Xposition: 4, Yposition: 4, Current_move: 5, 
	Cargo_space: 1, Cargo_amount: 0, Cargo_type: resource.Null, Can_move_water: false}
var Wagon = Unit{}
var Fishing_boat = Unit{Name: "fishing boat", Move: 5, Xposition: 4, Yposition: 4, Current_move: 5, 
	Cargo_space: 3, Cargo_amount: 0, Cargo_type: resource.Fish, Can_move_water: true}

type Unit struct {
	Name string
	Move, Xposition, Yposition int
	Current_move int
	Cargo_space int
	Cargo_amount int
	Cargo_type resource.Resource
	Can_move_water bool
}

func (unit *Unit) Load(building *building.Building) {
	freespace := unit.Cargo_space - unit.Cargo_amount
	unit.Cargo_type = building.Output_type //eventually i'll need a way to store multiple types of Cargo
	if building.Output_storage_amount >= freespace {
		unit.Cargo_amount += freespace
		building.Output_storage_amount -= freespace
	}
	if building.Output_storage_amount < freespace {
		unit.Cargo_amount += building.Output_storage_amount
		building.Output_storage_amount = 0
	}
}

func (unit *Unit) Unload(building *building.Building) {
	if (unit.Cargo_type == building.Input_type) {
		building.Input_storage_amount += unit.Cargo_amount
		unit.Cargo_amount = 0
	}
}

func (unit *Unit) Collect_resource(tile tile.Tile, resource resource.Resource) {
	if (tile.Is_forest == true && unit.Cargo_amount == 0) {
		unit.Cargo_type = resource
		unit.Cargo_amount = 1
	}
}

func (b *Unit) Collect_fish(route []tile.Tile) {
	if (b.Can_move_water == false) {return}
	b.Cargo_amount += len(route)
	if (b.Cargo_amount > b.Cargo_space) {b.Cargo_amount = b.Cargo_space}
}

func (u *Unit) Move_unit(x int, y int, tile_array [10][10]tile.Tile) {
	newxposition := u.Xposition + x
	newyposition := u.Yposition + y
	u.Current_move -= tile_array[newxposition][newyposition].Move_cost
	if u.Current_move >= 0 {
		u.Xposition += x	
		u.Yposition += y
	} else { u.Current_move = 0 }
}
