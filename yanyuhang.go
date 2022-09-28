package main

import (
	"fmt"
	"math"
	"strconv"
)

type Complex struct {
	realPart      float64
	imaginaryPart float64
}

func (self *Complex) Add(addNum *Complex) (result *Complex) { //此时result没定义!!!
	result = new(Complex)
	result.realPart = self.realPart + addNum.realPart
	result.imaginaryPart = self.imaginaryPart + addNum.imaginaryPart
	return
}

func (self *Complex) Minus(minusNum *Complex) (result *Complex) {
	result = new(Complex)
	result.realPart = self.realPart - minusNum.realPart
	result.imaginaryPart = self.imaginaryPart - minusNum.imaginaryPart
	return
}

func (self *Complex) Multi(multiNum *Complex) (result *Complex) {
	result = new(Complex)
	result.realPart = (self.realPart * multiNum.realPart) - (self.imaginaryPart * multiNum.imaginaryPart)
	result.imaginaryPart = (self.imaginaryPart * multiNum.realPart) + (self.realPart * multiNum.imaginaryPart)
	return
}

func (self *Complex) Divi(diviNum *Complex) (result *Complex) {
	result = new(Complex)
	lengthPower2 := (diviNum.realPart * diviNum.realPart) + (diviNum.imaginaryPart * diviNum.imaginaryPart)
	result.realPart = (self.realPart*diviNum.realPart + self.imaginaryPart*diviNum.imaginaryPart) / lengthPower2
	result.imaginaryPart = (self.imaginaryPart*diviNum.realPart - self.realPart*diviNum.imaginaryPart) / lengthPower2
	return
}

func (self *Complex) Length() float64 {
	return math.Sqrt((self.realPart * self.realPart) + (self.imaginaryPart * self.imaginaryPart))
}

func (self *Complex) String() string {
	if self.imaginaryPart < 0 {
		return "(" + strconv.FormatFloat(self.realPart, 'f', -1, 64) + strconv.FormatFloat(self.imaginaryPart, 'f', -1, 64) + "i" + ")"
	} else {
		return "(" + strconv.FormatFloat(self.realPart, 'f', -1, 64) + "+" + strconv.FormatFloat(self.imaginaryPart, 'f', -1, 64) + "i" + ")"

	}
}

func (self *Complex) SetValue(real, imag float64) {
	self.realPart = real
	self.imaginaryPart = imag
}
func main() {
	var nums1 = new(Complex)
	var nums2 = new(Complex)
	fmt.Print("请输入第一个复数")
	nums1.SetValue(checkInput())
	fmt.Print("请输入第二个复数")
	nums2.SetValue(checkInput())
	fmt.Println(nums1.Add(nums2))
	fmt.Println(nums1.Minus(nums2))
	fmt.Println(nums1.Multi(nums2))
	fmt.Println(nums1.Divi(nums2))
	fmt.Println(nums1.Length(), ",", nums2.Length())
}

func checkInput() (real float64, imag float64) {
	var err error
	for {
		_, err = fmt.Scan(&real, &imag)
		if err != nil {
			if err.Error() == "strconv.ParseFloat: parsing \"\": invalid syntax" {
				continue
			}
			fmt.Println(fmt.Printf("输入错误请重新输入\n"))
			//fmt.Println(err)
			//fmt.Printf("输入错误请重新输入\n")
			continue
		} else {
			break
		}
	}
	return
}
