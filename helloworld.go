package main 
import ( 
	"fmt"
	"time"
)

type Book struct { 
	title string
}

func main(){ 
	// learn_var()
	// learn_struct()
	// learn_map()
	learn_channel()
}

func learn_channel() { 
	var c = make(chan string,1)
	go say("hello",c)
	go say("world",c)
	go say("theThre3",c)
	fmt.Printf("%s\n%s\n%s",<-c,<-c,<-c)
}

func say(s string,c chan string){ 
	for i:=0;i<100;i++ { 
		time.Sleep(10*time.Millisecond)
		//fmt.Println(s,i)
	}
	c <- s+" finish"
}

func learn_map(){ 
	nums := []int{1,2,3,5,7}
	sum := 0
	for _,num := range nums{ 
		sum += num 
	}
	fmt.Println("sum:",sum)

	var kvs map[int]string
	kvs = make(map[int]string)
	kvs[1] = "apple"
	kvs[2] = "banana"
	kvs[3] = "grape"
 	for k,v := range kvs { 
		fmt.Printf("%d -> %s\n",k,v)
	}
	fmt.Println("delete kvs key=1")
	delete(kvs,1)
	for k,v := range kvs { 
		fmt.Printf("%d -> %s\n",k,v)
	}


	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "hello" {
        fmt.Println(i, c)
    }
}

func learn_struct(){
	
	var book Book;
	book.title = "h1"
	fmt.Println(book)

	setBook1(book)
	fmt.Println(book)
	setBook2(&book)
	fmt.Println(book) 
}

func setBook1(book Book){ 
	book.title = "h2"
}

func setBook2(book *Book){ 
	book.title = "h3"
}

func learn_var(){ 
	var a = "RUNOOB"
	fmt.Println(a)

	var b int 
	fmt.Println(b)

	var c bool 
	fmt.Println(c)

	d := calc()
	fmt.Println(d)
}

func calc() int { 
	len := 5
	wid := 6

	area := len * wid 
	fmt.Println("area:",area)

    return area
}

