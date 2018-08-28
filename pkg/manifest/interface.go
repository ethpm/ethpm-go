package manifest

// Interface The interfacd for an ethpm manifest
type Interface interface {
	Read(s string) (err error)
	Write() (s string, err error)
}
