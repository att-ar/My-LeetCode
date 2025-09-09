#include <iostream>
#include <vector>
#include <numeric>

/* 1863. Sum of All Subset XOR Totals
39% time, 24% memory
(My solution is much faster in reality but the input size is so small that the non-cached DFS with less overhead runs faster)

Keep running score of all subsets' XOR totals
Can do this instead of storing the subsets and computing XOR at the end
Saves memory

```python
subset = [[]]
for num in nums:
    for i in range(len(subset)):
        # dub = subset[i].copy()
        # dub.append(num)
        subset.append(subset[i].copy() + [num])
```
*/
int subsetXORSum(std::vector<int> &nums)
{
    std::vector<int> result;
    for (int num : nums)
    {
        size_t resultSize{result.size()};
        for (size_t j = 0; j < resultSize; j++)
        {
            result.emplace_back(result[j] ^ num);
        }
        result.emplace_back(num);
    }
    // I don't need to pass std::plus<int>() because it will use std::plus<> automatically
    // but I imagine telling it the type explicitly saves some work at compile time
    // and its to tell myself that I am allowed to change the Reducer()
    return std::reduce(result.begin(), result.end(), 0, std::plus<int>{});
}

/*
3396. Minimum Number of Operations to Make Elements in Array Distinct
100% time, 94% memory

Asking you to find the highest index such that the number at that index is a duplicate
*/
int minimumOperations(std::vector<int> &nums)
{
    std::bitset<101> seen; // use bitset as a seen map
    size_t idx{0};
    for (size_t i = nums.size(); i > 0; i--)
    {
        // first duplicate found in reverse order is the return condition
        if (seen[nums[i - 1]])
        {
            idx = i;
            break;
        }
        seen.set(nums[i - 1]); // mark as seen
    }
    // std::cout << idx << '\n';
    return (idx + 2) / 3; // + 2 to replicate ceiling function
}

/*
3375. Minimum Operations to Make Array Values Equal to K

99% time, 92% memory
Just need to count the number of distinct integers > k
Failure if any number is less than k
*/
int minOperations(std::vector<int> &nums, int k)
{
    std::bitset<101> seen;
    int count{0};
    for (int num : nums)
    {
        if (num < k)
            return -1;
        count += (!seen[num] && num > k) ? 1 : 0;
        seen.set(num);
    }
    return count;
}

/* 1534. Count Good Triplets
97% tie, 43% memory
Declaring `ni` and `nj` makes a big difference because of they allow the k-loop to be cache optimized
*/
int countGoodTriplets(std::vector<int> &arr, int a, int b, int c)
{
    size_t n{arr.size()};
    int result{};
    for (size_t i = 0; i < n; i++)
    {
        int ni{arr[i]};
        for (size_t j = i + 1; j < n; j++)
        {
            int nj{arr[j]};
            if (std::abs(ni - nj) > a)
                continue;
            for (size_t k = j + 1; k < n; k++)
            {
                int nk{arr[k]};
                if ((std::abs(nj - nk) <= b) && (std::abs(ni - nk) <= c))
                {
                    result++;
                }
            }
        }
    }
    return result;
}

int main()
{
    std::vector<int> result;
    std::vector<int> nums;

    std::printf("\n\nSum of All Subset XOR Totals\n");
    nums = {1, 3};
    std::printf("%d\n", subsetXORSum(nums));

    return 0;
}
