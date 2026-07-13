from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        s_map: dict[int, int] = {}
        
        for i, num in enumerate(nums):
            diff = target - num
            if diff in s_map:
                return [s_map[diff], i]
            s_map[num] = i
        
        return []