package raise

func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}
