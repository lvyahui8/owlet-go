package utils

import (
    "io/fs"
    "io/ioutil"
    "path"
    "path/filepath"
    "runtime"
)

func CopyFile(src, dst string) error {
    content, err := ioutil.ReadFile(src)
    if err != nil {
        return err
    }
    return ioutil.WriteFile(dst, content, fs.ModePerm)
}

func RootDir() string {
    _, b, _, _ := runtime.Caller(0)
    d := path.Join(path.Dir(b))
    return filepath.Dir(d)
}
