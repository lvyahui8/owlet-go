package async_call

func fc1() {
    defer fc3()
    go fc2()
}

func fc2() {

}

func fc3() {

}