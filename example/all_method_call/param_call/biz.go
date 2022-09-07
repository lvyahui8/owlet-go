package param_call


func fc3(fc func()) {
    fc()
}

func fc4()  {

}

func fc8(fc func()) func() {
    return func() {
        fc()
    }
}

func fc9() {

}

func home() {
    fc3(fc4)
    fc8(fc9)()
}

