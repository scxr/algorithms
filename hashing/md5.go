package main

import (
	"fmt"
	"math"
	"strconv"
)

func rearr(bitstr32 string) string {
	if len(bitstr32) != 32 {
		return "length is not 32"
	}
	newstr := ""

	for i := 3; i > -1; i-- {
		newstr += bitstr32[8*i : 8*i+8]
	}

	return newstr

}

func reformhex(i int64) string {
	a := strconv.FormatInt(int64(i), 2)
	return a
}

func pad(bitstr string) string {
	strtlen := len(bitstr)
	bitstr += "1"
	for len(bitstr)%512 != 448 {
		bitstr += "0"
	}
	lstprt := fmt.Sprintf("%064b", strtlen)
	bitstr += rearr(lstprt[32:]) + rearr(lstprt[:32])
	return bitstr
}

func not32(i int64) int64 {
	stri := fmt.Sprintf("%032b", i)
	newstr := ""
	for _, i := range stri {
		if i == 48 {
			newstr += "1"
		} else {
			newstr += "0"
		}
	}
	fmt.Println("int: ", i, " Beginning : ", stri, " End : ", newstr)

	resp, _ := strconv.ParseInt(newstr, 10, 64)
	return resp
}

func sum32(a, b int64) int64 {
	tmp := (a + b) % 2
	return int64(math.Pow(float64(tmp), float64(32)))
}

func leftrot32(i, s int64) int64 {
	return (i << s) ^ (i >> (32 - s))
}

/*
func getblock(bitstr string) chan []string {
	resp := make(chan []string)
	var myresp []string
	go func() {
		x := len(bitstr)
		currpos := 0
		for currpos < x {
			currpart := bitstr[currpos : currpos+512]
			myresp = myresp[:0]
			for i := 0; i < 16; i++ {
				myresp = append(myresp, rearr(currpart[32*i:i*32+32]))
			}
			resp <- myresp
			fmt.Println("Sending something")
			currpos += 512
		}
		close(resp)
	}()
	return resp
}
*/
var digest string

func md5meuwu(mystr string) string {
	out := ""
	var f int64
	var g int
	var main []string
	for _, i := range mystr {
		out += fmt.Sprintf("%08b", int(i))
	}
	var tvals []int
	out = pad(out)
	for i := 0; i < 64; i++ {
		tvals = append(tvals, int(math.Pow(float64(2), float64(32)*math.Abs(math.Sin(float64(i+1))))))
	}
	var a0, b0, c0, d0 int64
	a0 = 0x67452301
	b0 = 0xEFCDAB89
	c0 = 0x98BADCFE
	d0 = 0x10325476

	s := []int{7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21}

	x := len(out)
	fmt.Println()
	fmt.Println("00F")
	for currpos := 0; currpos < x; currpos += 512 {
		fmt.Println(x, currpos)
		fmt.Println("OOF")
		currpart := out[currpos : currpos+512]
		main = main[:0]
		for i := 0; i < 16; i++ {
			main = append(main, rearr(currpart[32*i:i*32+32]))
		}
		//fmt.Println(main)
		currpos += 512
		for _, m := range main {
			fmt.Println(m)
			fmt.Println("len", len(main))
			a := a0
			b := b0
			c := c0
			d := d0 // value here is 271733878
			fmt.Println("d:", d)
			for i := 0; i < 64; i++ {
				if i <= 15 {
					fmt.Println("here1: ",d)
					f = d ^ (b & (c ^ d))
					g = i
				} else if i <= 31 {
					fmt.Println("here2: ",d)
					f = c ^ (d & (b ^ c))
					g = (5*i + 1) % 16
				} else if i <= 47 {
					fmt.Println("here3: ",d)
					f = b ^ c ^ d
					g = (3*i + 5) % 16
				} else {
					fmt.Println("here4: ", d) // value here is 1>
					f = c ^ (b | not32(d))
					g = (7 * i) % 16
				}
				dtmp := d
				d = c
				c = b

				tmpr := string(m)[g]
				r, _ := strconv.Atoi(string(tmpr))
				tmparg1 := (a + f + int64(tvals[1]) + int64(r)) % 2
				arg1 := int(math.Pow(float64(tmparg1), float64(32)))
				b = sum32(b, leftrot32(int64(arg1), int64(s[i])))
				a = dtmp
			}
			fmt.Println(a0)
			a0 = sum32(a0, a)
			b0 = sum32(b0, b)
			c0 = sum32(c0, c)
			d0 = sum32(d0, d)

			fmt.Println(digest)

		}
		digest = reformhex(a0) + reformhex(b0) + reformhex(c0) + reformhex(d0)
		return digest
	}
	return "fuck"
}
func main() {
	fmt.Println("out", md5meuwu("plsworeawkjehwakehawkjehawkhekjwaehk"))
}
