package helpF


type Todo struct {
	Stat string `json:"stat"`
	Todo string `json:"todo"`
	Done bool   `json:"done"`
}
type Error struct {
	Msg string
}

func Reverse(todoe *Todo) {
	todoe.Done = !(todoe.Done)
}
func Free(todoe *Todo) {
	//todoe = nil
	todoe.Stat = ""
	todoe.Todo = ""
	todoe.Done = false
}
func Helpertodo(stat string,list []Todo) (*Todo, *Error) {
	not := &Error{Msg:"not FOUND"}

	
	for i, ele := range list {
		if ele.Stat == stat {
			return &list[i],nil 
		}
	}
	return nil, not
}