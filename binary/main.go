package main

// 异或运算的三个性质
// 任何数与0异或，结果任然是本身 a^0 = a
// 任何数与自身异或，结果是零 a^a = 0
// 异或运算满足交换律和结合律 a^b^a=b^(a^a)=b^0=b

// 136 只出现一次的数字
func singleNumber(nums []int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}

// 137 只出现一次的数字II
// 取一个数的第i个二进制位 num>>i & 1
// 用二进制位组成一个数（类似于十进制的加法）第i个二进制位左移i位与ans按位或
func singleNumberii(nums []int) int {
	var result int32
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}

		result |= total % 3 << i

	}
	return int(result)
}

// 191 位1的个数
// func hammingWeight(n int) int {
// 	var count int
// 	for i := 0; i < 32; i++ {
// 		if int32(n)>>i&1 == 1 {
// 			count++
// 		}
// 	}
// 	return count
// }
// n & (n-1) 把n的二进制位的最低位的1变为0
// 利用 Brian Kernighan 算法，可以在一定程度上进一步提升计算速度。
// Brian Kernighan 算法的原理是：对于任意整数 x，令 x=x & (x−1)，
// 该运算将 x 的二进制表示的最后一个 1 变成 0。因此，对 x 重复该操作，
// 直到 x 变成 0，则操作次数即为 x 的「一比特数」。
func hammingWeight(n int) int {
	var count int
	num := uint32(n)
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}

// 338 比特位计数

// func countBits(n int) []int {
// 	ans := make([]int, n+1)
// 	for i:=0; i< n+1; i++ {
// 		ans[i]= hammingWeight(i)
// 	}
// 	return ans
// }
// func countBits(n int) []int {
// 	ans := make([]int, n+1)
// 	for i := 0; i < n+1; i++ {
// 		ans[i] = bits.OnesCount32(uint32(i))
// 	}
// 	return ans
// }
// 动态规划
// 最高有效位
// func countBits(n int) []int {
// 	bits := make([]int, n+1)
// 	bits[0] = 0
// 	var hightBit int
// 	for i := 1; i < n+1; i++ {
// 		if i&(i-1) == 0 {
// 			hightBit = i
// 		}
// 		bits[i] = bits[i-hightBit] + 1
// 	}
// 	return bits
// }
// 最低有效位
// 第i位等于第i/2位（i>>1) + 奇数加1，偶数不加
func countBits(n int) []int {
	bits := make([]int, n+1)
	bits[0] = 0
	for i := 1; i < n+1; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

// 190 颠倒二进制位
// func reverseBits(num uint32) uint32 {
// 	var res uint32
// 	for i := 0; i < 32; i++ {
// 		res = res>>i | num&1
// 		num = num >> 1
// 	}
// 	return res
// }

// 分治

func reverseBits(num uint32) uint32 {
	left := num >> 16
	right := num << 16
}
