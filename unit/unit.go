package unit

import ( "github.com/zudijyr/tiletrade/resource" )
import ( "github.com/zudijyr/tiletrade/building" )
import ( "github.com/zudijyr/tiletrade/tile"     )

const WORLD_SIZE = 10

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

func (u *Unit) Move_unit(x int, y int, tile_array [10][10]tile.Tile) {
	newxposition := u.Xposition + x
	newyposition := u.Yposition + y
	u.Current_move -= tile_array[newxposition][newyposition].Move_cost
	if u.Current_move >= 0 {
		u.Xposition += x	
		u.Yposition += y
	} else { u.Current_move = 0 }
}
