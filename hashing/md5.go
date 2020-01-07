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

func reformhex(i int) string {
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

func not32(i int) int {
	stri := fmt.Sprintf("%032b", i)
	newstr := ""
	for _, i := range stri {
		if i == 48 {
			newstr += "1"
		} else {
			newstr += "0"
		}
	}
	fmt.Println("Beginning : ", stri, " End : ", newstr)
	resp, _ := strconv.Atoi(newstr)
	return resp
}

func sum32(a, b int) int {
	tmp := (a + b) % 2
	return int(math.Pow(float64(tmp), float64(32)))
}

func leftrot32(i, s int) int {
	return (i << s) ^ (i >> (32 - s))
}

func md5meuwu(mystr string) string {
	out := ""
	var f, g int
	for _, i := range mystr {
		out += fmt.Sprintf("%08b", int(i))
	}
	var tvals []int
	out = pad(out)
	for i := 0; i < 64; i++ {
		tvals = append(tvals, int(math.Pow(float64(2), float64(32)*math.Abs(math.Sin(float64(i+1))))))
	}
	var (
		a0 = 0x67452301
		b0 = 0xEFCDAB89
		c0 = 0x98BADCFE
		d0 = 0x10325476
	)

	s := []int{7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21}

	for m := range getblock(out) {
		fmt.Println(m)
		a := a0
		b := b0
		c := c0
		d := d0
		for i := 0; i < 64; i++ {
			if i <= 15 {
				f = d ^ (b & (c ^ d))
				g = i
			} else if i <= 31 {
				f = c ^ (d & (b ^ c))
				g = (5*i + 1) % 16
			} else if i <= 47 {
				f = b ^ c ^ d
				g = (3*i + 5) % 16
			} else {
				f = c ^ (b | not32(d))
				g = (7 * i) % 16
			}
			dtmp := d
			d = c
			c = b
			r, _ := strconv.Atoi(m[g])
			tmparg1 := (a + f + int(tvals[1]) + r) % 2
			arg1 := int(math.Pow(float64(tmparg1), float64(32)))
			b = sum32(b, leftrot32(arg1, s[i]))
			a = dtmp
		}
		a0 = sum32(a0, a)
		b0 = sum32(b0, b)
		c0 = sum32(c0, c)
		d0 = sum32(d0, d)
		digest := reformhex(a0) + reformhex(b0) + reformhex(c0) + reformhex(d0)
		fmt.Println(digest)
		return digest
	}
	return "fuck"
}
func main() {
	fmt.Println("out", md5meuwu("plsworeawkjehwakehawkjehawkhekjwaehk"))
}
