package tile

import ( "github.com/zudijyr/tiletrade/building" )

type Tile struct {
	Move_cost int
	Xposition int
	Yposition int
	Is_land bool
	Building building.Building
	Is_grass bool
	Is_plains bool
	Is_hills bool
	Is_forest bool
}

func (t *Tile) print_is_land() {
	println(t.Is_land)
}
