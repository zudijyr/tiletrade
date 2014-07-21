package resource


//resource initializations
var Null   = Resource{Name: "null"}
var Gold   = Resource{Name: "gold"}
var Lead   = Resource{Name: "lead"}
var People = Resource{Name: "people"}
var Food   = Resource{Name: "food"}
var Grain  = Resource{Name: "grain"}
var Fish   = Resource{Name: "fish"}
var Trees  = Resource{Name: "trees"}
var Lumber = Resource{Name: "lumber"}
//end resource initializations

type Resource struct {
	Name string
}
