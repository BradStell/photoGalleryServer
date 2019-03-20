package images

type Image struct {
	ID   string
	Name string
}

func (i *Image) create() *Image {
	return Create(i)
}

func (i *Image) update() *Image {
	return nil
}

func (i *Image) delete() *Image {
	return nil
}

func (i *Image) softDelete() *Image {
	return nil
}

func Get(id string) *Image {
	return nil
}

func GetAll() []*Image {
	return nil
}
