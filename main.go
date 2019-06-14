package findthing

import "log"

const (
	VERSION_MAJOR = 0
	VERSION_MINOR = 1
	VERSION_PATCH = 1
)

func main() {
	log.Printf("Find thing v%v.%v.%v", VERSION_MAJOR, VERSION_MINOR, VERSION_PATCH)


}