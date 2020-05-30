package packer

type FileWriter interface {
	WriteFile(name string, data []byte) bool
}
