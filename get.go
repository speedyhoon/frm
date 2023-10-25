package frm

var GetFields func(uint8) []Field

// GetForms retrieves forms by their ids.
//
// The duplicate ids parameter is to ensure at least one uint8 is passed to the function without adding additional checks (same as the built-in exec.Command()).
func GetForms(id uint8, ids ...uint8) (f map[uint8]Form) {
	f = map[uint8]Form{
		id: {Action: id, Fields: GetFields(id)},
	}
	for _, id = range ids { // Reuse the existing `id` variable.
		f[id] = Form{Action: id, Fields: GetFields(id)}
	}

	return
}
