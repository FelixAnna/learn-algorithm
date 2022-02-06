public class Solution {
    public int Search(int[] nums, int target) {
        var len = nums.Length;
        if(len == 1) {
            return nums[0] == target ? 0 : -1;
        }

        if(nums[0] > nums[len-1]) {
            //rorated
            //1. search rotated
            var rotatedIndex =0;
            for(int i=1; i<len; i++){
                if(nums[i]<nums[i-1]){
                    rotatedIndex = i-1;
                }
            }

            //2. bin search in left and right
            var leftIndex = BinSearch(nums, target, 0, rotatedIndex);
            var rightIndex = BinSearch(nums, target, rotatedIndex+1, len-1);

            //return max index
            return leftIndex>rightIndex?leftIndex:rightIndex;

        }else{
            return BinSearch(nums, target, 0, len-1);
        }
    }

    private int BinSearch(int[] nums, int target, int start, int end) {
        if(start<=end){
            var middle = (end-start)/2 + start;
            var currentElement = nums[middle];
            if(target == currentElement) {
                return middle;
            }else if (target < currentElement) {
                return BinSearch(nums, target, start, middle+1);
            }else{
                return BinSearch(nums, target, middle + 1, end);
            }
        }

        return -1;
    }
}