// Package frm provides several structs to represent data received from HTML <form> elements via Get or Post HTTP requests.
//
// [GetFields] must be assigned for this package to be useful.
package frm

// GetFields defines a function signature to retrieve a form's fields,
// instead of passing function definitions, pointers, or interfaces between packages.
// GetFields is relied upon by external packages including [session] and [vl]
// to provide additional session handling and validation functionality.
//
// [GetForms] requires GetFields to be populated.
//
// An example implementation:
//
//	func init() {
//		const frmNew, frmSave = 0, 1
//		frm.GetFields = func(id uint8) []frm.Field {
//			switch id {
//			case frmNew:
//				return []frm.Field{{Name: "new", Vl: vl.StrReq}}
//			case frmSave:
//				return []frm.Field{{Name: "id", Vl: vl.Regex}, {Name: "save", Vl: vl.StrReq}}
//			}
//			return nil
//		}
//	}
//
// A more advanced implementation is available for reference in [EventBucket/go/forms.go].
//
// [BuildIt.ninja] can generate the GetFields function when using its Forms code generation feature.
//
// [BuildIt.ninja]: https://github.com/speedyhoon/BuildIt.ninja
// [session]: https://github.com/speedyhoon/session
// [vl]: https://github.com/speedyhoon/vl
// [EventBucket/go/forms.go]: https://github.com/speedyhoon/EventBucket/blob/main/go/forms.go#L66
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
