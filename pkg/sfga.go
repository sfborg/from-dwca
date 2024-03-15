package fdwca

func (fd *fdwca) checkSFGA() bool {
	return fd.stor.Exists()
}
