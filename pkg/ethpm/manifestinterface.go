package ethpm

// ManifestInterface The interfacd for an ethpm manifest
type ManifestInterface interface {
	Read(s string) (err error)
	Write() (s string, err error)
}
