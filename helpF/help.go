package helpF


type Todo struct {
	Stat string `json:"stat"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}
type Error struct {
	Message string
}
var todos = []Todo{
	{Stat: "1", Todo: "Clean room", Done: true},
	{Stat: "2", Todo: "Salat", Done: true},
	{Stat: "3", Todo: "Learning", Done: false},

}
func Reverse(todoe Todo) {
	todoe.Done = !(todoe.Done)
}
func Free(todoe Todo) {
	//todoe = nil
	todoe.Stat = ""
	todoe.Todo = ""
	todoe.Done = false
}
func Helpertodo(stat string) (*Todo, *Error) {
	not := &Error{Message:"not FOUND"}
	
	for i, ele := range todos {
		if ele.Stat == stat {
			return &todos[i],nil 
		}
	}
	return nil, not
}