package rom

// ActorDescription is scapped actor data from the CloudModding wiki
type ActorDescription struct {
	ID             uint16
	FileName       string
	Object         uint16
	Translation    string
	Identification string
}

// ActorDescriptions maps Actor IDs to their debug information and human-readable names.
var ActorDescriptions = map[uint16]ActorDescription{}
