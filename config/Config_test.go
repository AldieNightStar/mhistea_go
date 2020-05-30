package config

import "testing"

// ==========================
// Test Function: ParseConfig
// ==========================
func TestParseConfig(t *testing.T) {
	text := "name = Marta\n\t\tage = 32\n\t count= 13"
	cfg := ParseConfig(text)

	if len(cfg) < 3 {
		t.Fail()
		return
	}
	if cfg["name"] != "Marta" {
		t.Fail()
		return
	}
	if cfg["age"] != "32" {
		t.Fail()
		return
	}
	if cfg["count"] != "13" {
		t.Fail()
		return
	}
}

// ==========================
// Test Function: ReadConfig
// ==========================
type FileRD struct{}

func (f FileRD) ReadFile(_ string) []byte {
	return []byte("cafe babe")
}

type SectionRD struct{}

func (s SectionRD) GetSectionList(text string) []string {
	if text == "cafe babe" {
		return []string{"sec1", "sec2"}
	}
	return []string{}
}

func (s SectionRD) GetSectionByName(_, name string) string {
	if name == "sec1" {
		return "name = Ihor\n\t\tage =21"
	} else if name == "sec2" {
		return "name = Andre\n\t\tage = 30\ntype = Human"
	}
	return ""
}

func TestReadConfig(t *testing.T) {
	freader := FileRD{}
	sreader := SectionRD{}
	cfg := ReadConfig(freader, sreader, "SomeFile.txt")

	if cfg == nil {
		t.Fail()
		return
	}
	if cfg.Get("sec1", "name") != "Ihor" {
		t.Fail()
		return
	}
	if cfg.Get("sec1", "age") != "21" {
		t.Fail()
		return
	}
	if cfg.Get("sec2", "name") != "Andre" {
		t.Fail()
		return
	}
	if cfg.Get("sec2", "age") != "30" {
		t.Fail()
		return
	}
	if cfg.Get("sec2", "type") != "Human" {
		t.Fail()
		return
	}
	if cfg.Get("xxx", "abc") != "" {
		t.Fail()
		return
	}
}
