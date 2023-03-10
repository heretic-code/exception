package exception

// Don't throw errors in Go, you idiot. Use the appropriate way of doing things
// instead of trying to force your preferred way onto a language which has not
// been designed to do it.
func Throw(err error) {
	panic(err)
}

// Catch is something that shouldn't be possible because you don't throw.
func Catch(f func(), handling func(any)) {
	defer func() {
		handling(recover())
	}()

	f()
}
