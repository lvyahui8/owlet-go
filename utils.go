package owlet

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

func GetCommonPrefix(paths []string) string {
    res := ""
    for i := 0; true; i++ {
        var preCh uint8
        for _, p := range paths {
            if i >= len(p) {
                return res
            }
            ch := p[i]
            if preCh == 0 {
                preCh = ch
            }
            if ch != preCh {
                return res
            }
        }
        res = res + string(preCh)
    }
    return res
}