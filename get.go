package frm

// GetFields defines a function signature to retrieve a form's fields, instead of passing function definitions, pointers or interfaces between packages.
// [BuildIt.ninja] can generate this code when using the Forms feature.
// GetFields is used by [session].
//
// [BuildIt.ninja]: https://github.com/speedyhoon/BuildIt.ninja
// [session]: https://github.com/speedyhoon/session
var GetFields func(uint8) []Field

// GetForms retrieves forms by their ids.
// The duplicate ids parameter is to ensure at least one uint8 is passed to the function without adding additional checks (same as Go's built-in exec.Command()).
func GetForms(id uint8, ids ...uint8) (f map[uint8]Form) {
	f = map[uint8]Form{
		id: {Action: id, Fields: GetFields(id)},
	}
	for _, id = range ids { // Reuse the existing `id` variable.
		f[id] = Form{Action: id, Fields: GetFields(id)}
	}

	return
}
