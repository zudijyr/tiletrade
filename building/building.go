package building

import ( "github.com/zudijyr/tiletrade/resource" )
//import ( "github.com/zudijyr/tiletrade/unit"     )

//building initializations
var Gold_mine    = Building{Name: "gold mine", Input_type: resource.People, Output_type: resource.Gold}
var Lead_mine    = Building{"lead mine", 2, 2, resource.People, 1, resource.Lead, 2,0,0}
var Grain_mill   = Building{"grain mill", 1, 1, resource.Grain, 1, resource.Food, 2,0,0}
var Lumber_mill  = Building{"lumber mill", 1, 1, resource.Trees, 1, resource.Lumber, 2,0,0}
var Grain_farm   = Building{"grain farm", 1, 1, resource.Null, 1, resource.Grain, 2,0,0}
var Fishing_dock = Building{"fishing dock", 1, 1, resource.Fish, 1, resource.Food, 2,0,0}


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
