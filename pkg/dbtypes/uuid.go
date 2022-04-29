package dbtypes

type UUID string

func (b *UUID) UnmarshalJSON(data []byte) error {

	(*b) = UUID(string(data))

	return nil
}
