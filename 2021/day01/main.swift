
import Foundation

let filename = "in.txt"

let arr = try String(contentsOfFile: "in.txt")
print(arr)
let nums = arr.split(separator:"\n").map{Int($0)!}
print(nums)

var last = nums[0]
var sum = 0
for n in nums {
	if n>last {
		sum += 1
	}
	last = n
}
print(sum) // 1696
print()

last = nums[2]
sum = 0
for (i,_) in nums.enumerated() {
	if i<3 { continue }
	if (nums[i]+nums[i-1]+nums[i-2])>last {
		sum += 1
	}
	last = nums[i]+nums[i-1]+nums[i-2]
}
print(sum) // 1737

