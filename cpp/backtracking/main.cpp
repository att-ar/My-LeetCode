#include <vector>
#include <numeric>
#include <iostream>
#include <iomanip>

/* 416. Partition Equal Subset Sum
96% time, 76% memory

I think this is a backtracking problem. Don't really remember tbh but I'll find out.

Problem is asking if you can for a subset such that sum(subset) == sum(array) / 2
- Requires an even number sum for the whole array.
- Will use a 2D array (index, current sum) as a cache
    - true : unvisited, false : failed
    - the entire call stack will pop back up on the first success so don't need to cache it
*/
bool canPdfs(std::vector<int> &nums, std::vector<std::vector<bool>> &cache, size_t i, size_t cursum, size_t target)
{
    // base cases:
    if (cursum == target)
    {
        return true;
    }
    else if (i == nums.size() || cursum > target)
    {
        return false;
    }
    else if (!cache[i][cursum])
    {
        return cache[i][cursum];
    }

    // 2 options:
    // take
    bool take{canPdfs(nums, cache, i + 1, cursum + static_cast<size_t>(nums[i]), target)};
    if (take)
        return true;
    // skip
    bool skip{canPdfs(nums, cache, i + 1, cursum, target)};
    if (skip)
        return true;
    // both failed
    cache[i][cursum] = false;
    return false;
}
bool canPartition(std::vector<int> &nums)
{
    // check for sum parity, need to static cast because size_t is unsigned so its narrowing
    const size_t sum{static_cast<size_t>(std::reduce(nums.begin(), nums.end()))};
    if (sum % 2)
    {
        return false; // odd number sum so impossible
    }
    const size_t target{sum / 2};
    std::vector<std::vector<bool>> cache(nums.size(), std::vector<bool>(target + 1, true));
    return canPdfs(nums, cache, 0, 0, target);
}

int main()
{
    std::vector<int> nums;
    std::cout << std::boolalpha;

    printf("\n\nPartition Equal Subset Sum\n");
    nums = {1, 5, 11, 5};
    std::cout << canPartition(nums) << '\n';
    nums = {1, 2, 3, 5};
    std::cout << canPartition(nums) << '\n';
    return 0;
}
