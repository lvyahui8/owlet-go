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

func main() {
    fc1()
    fc3(fc4)
}
