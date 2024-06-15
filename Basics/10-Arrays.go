package main

import "fmt"

func main8() {
	fmt.Println("Peddireddy Hari Vardhan Reddy")
	var ages [5]int = [5]int{1,2,3,4,5}
	fmt.Println(ages)
	
	var arrs = [6]int{20,30,40,50,60}
	fmt.Println(arrs)

	names := [5]string{"alien","arpit","sumit","george","grocery"}
	fmt.Println(names)

	nums := [6]int{}
	fmt.Println(nums)

	// Twod Arrays

	var twoD [3][3]int = [3][3]int{
								{1,2,3},
								{4,5,6},
							}
	fmt.Println(twoD)

	var twoDR [2][3]int
	var twarr [2][3]int
	for i:= 0; i<2; i++ {
		for j:=0;j<3;j++{
			twarr[i][j] = i+j
		}
	}
	fmt.Println(twarr)

	twoDR = [2][3]int{
		{6,7,8},
		{9,10,11},
	}
	fmt.Println(twoDR)

	b := [...]int{100,4:500}
	fmt.Println(b)

	a := [...]int{1,9:10}
	fmt.Println(a)

	// a[0] = 100
	// a[1] = 200
	// a[2] = 300
	// a[3] = 400
	// a[4] = 500
	fmt.Println(a)

	for i:=0;i<10;i++{
		fmt.Print(i," ")
	}

	fmt.Println()

	// slices use arrays underthe hood
	var scores = []int{100,200,300}
	fmt.Println(scores)

	scores[1] = 400
	scores = append(scores, 200)
	fmt.Println(scores)

	var marks []int = []int{1,9:100}
	fmt.Println(marks)
	for i := 0 ;i<len(marks);i++{
		marks[i] = i+i
	}
	marks = append(marks,20)
	fmt.Println(marks)
	//Slice will be created using array

	rangeOne := marks[1:4]
	fmt.Println(rangeOne)

	mySlice := []int{}
	fmt.Println(mySlice)
	println(cap(mySlice))
	fmt.Println(marks)

	// Slices:
	var s []string
	fmt.Println(s)
	s = []string{"Al","be","ce","dl"}
	fmt.Println(s)
	// Ranges Of an Arrays

	twoDr := make([][]int,3)
	fmt.Println(twoDr)
	for i:=0;i<3;i++{
		innerLen := i+1
		twoDr[i] = make([]int, innerLen)
		for j:= 0;j<innerLen;j++{
			twoDr[i][j] = i+j
		}
	}
	fmt.Println("2d",twoDr)

	// Slices  We can do slices with make function
	sli := make([]int,5,10)
	fmt.Println(sli,cap(sli),len(sli))

}