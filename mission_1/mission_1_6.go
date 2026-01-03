package main
import(
	"fmt"
)

func TwoSum(nums [] int, target int) [] int{
	for i:=0; i<len(nums); i++{
		for j:=i+1; j<len(nums); j++{
			if nums[i]+nums[j]==target{
				return []int{i,j}
			}
		}
	}
	return nil
}

func main(){
	nums := []int{11,15,2,7}
	target := 13
	result := TwoSum (nums, target)
	if result != nil {
		fmt.Println("找到两个数的索引:", result)
	} else {
		fmt.Println("未找到符合条件的两个数")
	}
}