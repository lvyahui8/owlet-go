package return_call


func fc5() {
    fc6()()
}

func fc6() func () {
    return fc7
}

func fc7() {

}