package main

func fc1() {
    fc2()
}

func fc2() {
}

func fc3(fc func()) {
    fc()
}

func fc4()  {

}

func fc5() {
    fc6()()
}

func fc6() func () {
    return fc7
}

func fc7() {

}

func main() {
    fc1()
    fc3(fc4)
    fc5()
}