package main

func sort(nums []int, l, r int) {
	if r-l <= 0 {
		return
	}

	pivot := nums[r]
	insert := l
	for i := l; i < r; i++ {
		if nums[i] <= pivot {
			nums[insert], nums[i] = nums[i], nums[insert]
			insert++
		}
	}
	nums[insert], nums[r] = nums[r], nums[insert]

	sort(nums, l, insert-1)
	sort(nums, insert+1, r)
}

func search(nums []int, l, r, target int) int {
	if r-l < 0 {
		return -1
	}

	m := (l + r) / 2

	if nums[m] == target {
		return m
	} else if nums[m] < target {
		return search(nums, m+1, r, target)
	} else {
		return search(nums, l, m-1, target)
	}
}

// [89585 47074 51199 30881 61256 41926 82925 51892 88913]
// [47074 51199 89585 =  insert 30881 = i 61256 41926 82925 51892 88913 = pivot, r]
