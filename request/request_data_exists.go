package request

func (data RequestData) EntryExists(entry string) bool {
	_, exists := data.Content[entry]

	return exists
}
