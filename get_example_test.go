package frm_test

import (
	"github.com/speedyhoon/frm"
)

func ExampleGetFields() {
	// Implementation:
	const frmClub, frmQty = 0, 1
	frm.GetFields = func(formId uint8) []frm.Field {
		switch formId {
		case frmClub:
			return []frm.Field{{Name: "n"}}
		case frmQty:
			return []frm.Field{{Value: 4}}
		}
		return nil
	}

	// Usage:
	frm.GetFields(frmClub)
	// Output: []frm.Field{{Name: "n", V8: vl.StrReq}}
}
