package rom

import (
	"encoding/json"

	"github.com/gobuffalo/packr"
)

func init() {
	data := packr.NewBox("../data")
	if err := initLoadActors(data); err != nil {
		panic(err)
	}
}

func initLoadActors(data packr.Box) error {
	actors, err := data.Find("actors.json")
	if err != nil {
		return err
	}

	list := make([]ActorDescription, 0, 512)
	if err := json.Unmarshal(actors, &list); err != nil {
		return err
	}

	for _, v := range list {
		ActorDescriptions[v.ID] = v
	}

	return nil
}
