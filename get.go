package frm

var GetFields func(uint8) []Field

// GetForms retrieves forms by their ids.
func GetForms(formIDs ...uint8) (f map[uint8]Form) {
	if len(formIDs) == 0 {
		return nil
	}

	f = make(map[uint8]Form, len(formIDs))
	for _, formId := range formIDs {
		f[formId] = Form{Action: formId, Fields: GetFields(formId)}
	}

	return
}
