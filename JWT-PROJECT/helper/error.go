package helper

func Errorpanic(err error) {
	if err != nil {

		panic(err)
	}
}
