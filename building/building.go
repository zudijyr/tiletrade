package building

import ( "github.com/zudijyr/tiletrade/resource" )
//import ( "github.com/zudijyr/tiletrade/unit"     )

type Building struct {
	Name string
	Size, Staff int
	Input_type resource.Resource
	Input_amount int	
	Output_type resource.Resource
	Output_amount int	
	Input_storage_amount, Output_storage_amount int
}

func (building *Building) Production() {
		production := 0
		if (building.Input_storage_amount > building.Staff) {
			//i'll change this later, once i implement real time, it should just dump the cargo
			production = building.Staff*building.Output_amount
			building.Input_storage_amount -= building.Staff
		}
		if (building.Input_storage_amount < building.Staff) {
			production = building.Output_amount*building.Input_storage_amount
			building.Input_storage_amount = 0
		}
		building.Output_storage_amount += production
}

//type Unit_Building struct {
//	Building
//	Output_unit_type unit.Unit
//	Output_unit_amount int
//}

//func (building *Unit_Building) Unit_Production(tile tile.Tile) *(unit.Unit){
//	unit := building.Output_unit_type 
//unit.Xposition = tile.Xposition
//	unit.Yposition = tile.Yposition
//	return &unit
//}
