package trigger

import "io/ioutil"

func EnumInDirectory(path string, recursive bool) (files []string, dirs []string, err error) {

	enums, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	files = make([]string, 0)
	dirs = make([]string, 0)

	for _, enum := range enums {
		if enum.IsDir() {
			dirs = append(dirs, enum.Name())

			if recursive {
				childFiles, childDirs, err := EnumInDirectory(enum.Name(), recursive)
				if err != nil {
					return nil, nil, err
				}
				files = append(files, childFiles...)
				dirs = append(dirs, childDirs...)
			}

		} else {
			files = append(files, enum.Name())
		}
	}
	return files, dirs, nil
}
