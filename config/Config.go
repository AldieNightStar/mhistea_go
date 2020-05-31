package config

//	Configuration
//		cfg := ReadConfig(fileReader, sectionReader, "config.txt")
//		cfg.Get("section_a", "name") // etc
type cfg struct {
	m map[string]map[string]string
}

func (c cfg) Get(section, key string) (value string) {
	return c.m[section][key]
}

func (c cfg) Sections() []string {
	var list []string
	for k, _ := range c.m {
		list = append(list, k)
	}
	return list
}

func (c cfg) Keys(sectionName string) []string {
	section, ok := c.m[sectionName]
	if !ok {
		return nil
	}
	var list []string
	for k, _ := range section {
		list = append(list, k)
	}
	return list
}
