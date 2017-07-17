package lockfile

import (
	"encoding/json"
	"io/ioutil"
)

// Unmarshal the json values into a lockfile go struct
func UnmarshalLock(path string) (lockfile Lock, err error) {
	if data, err := ioutil.ReadFile(path); err != nil {
		panic(err)
	} else {
		if err = json.Unmarshal(data, &lockfile); err != nil {
			panic(err)
		}
		return
	}
}

// Validate all the lockfile fields
func Validate(lockfile Lock) (Lock, error) {

}
