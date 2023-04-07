package mongodb

func (m mongodbImpl) Name() string {
	return m.database.Name()
}

func (m mongodbImpl) Collection(name string) Collection {
	return m.database.Collection(name)
}

func (m mongodbImpl) URI() string {
	return m.uri
}
