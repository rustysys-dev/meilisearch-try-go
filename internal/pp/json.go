package pp

import (
	"github.com/TylerBrock/colorjson"
)

func MarshalJSON(i interface{}) ([]byte, error) {
	json := colorjson.NewFormatter()
	json.Indent = 2
	return json.Marshal(i)
}
