package frm_test

import (
	"github.com/speedyhoon/frm"
	"testing"
)

func TestGetForms_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected an unpopulated GetGields() to panic when executed")
		}
	}()

	frm.GetForms(0)
}

func TestGetForms_OK(t *testing.T) {
	const frmNew, frmSave = 0, 1
	frm.GetFields = func(id uint8) []frm.Field {
		switch id {
		case frmNew:
			return []frm.Field{{Name: "new"}}
		case frmSave:
			return []frm.Field{{Name: "id"}, {Name: "save"}}
		}
		return nil
	}

	f := frm.GetForms(frmNew)
	if len(f) != 1 {
		t.Errorf("expected GetForms() to return 1 form, have %d", len(f))
	}

	f = frm.GetForms(frmSave, frmNew)
	if len(f) != 2 {
		t.Errorf("expected GetForms() to return 2 forms, have %d", len(f))
	}

	if len(f[frmNew].Fields) != 1 {
		t.Errorf("expected frmNew to contain 1 field, have %d", len(f))
	}

	if f[frmNew].Fields[0].Name != "new" {
		t.Errorf("expected frmNew field to be `new`, have %s", f[frmNew].Fields[0].Name)
	}

	if len(f[frmSave].Fields) != 2 {
		t.Errorf("expected frmNew to contain 2 fields, have %d", len(f))
	}

	if f[frmSave].Fields[1].Name != "save" {
		t.Errorf("expected frmSave field to be `save`, have %s", f[frmSave].Fields[1].Name)
	}
}
