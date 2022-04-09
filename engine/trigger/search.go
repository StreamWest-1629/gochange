package trigger

import "io/ioutil"

func EnumInDirectory(path string) (files []string, dirs []string, err error) {

	enums, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	files = make([]string, 0)
	dirs = make([]string, 0)

	for _, enum := range enums {
		if enum.IsDir() {
			dirs = append(dirs, enum.Name())
		} else {
			files = append(files, enum.Name())
		}
	}

	return files, dirs, nil
}
