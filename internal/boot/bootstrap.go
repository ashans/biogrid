package boot

import "os"

func Init() {
	err := os.Setenv("FYNE_SCALE", "0.8")
	if err != nil {
		panic(err)
	}
}
