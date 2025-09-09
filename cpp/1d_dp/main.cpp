#include <iostream>
#include <vector>
#include <algorithm>

/*2873. Maximum Value of an Ordered Triplet II
100% time, 91% memory

Same solution as version 1 but I will space reduce since I don't need arrays
Recurrence Relation is only dependent on the previous state
*/
long long maximumTripletValueSpaceOptimized(std::vector<int> &nums)
{
    long long singlet{0}, doublet{0}, triplet{0};
    for (size_t i = 0; i < nums.size(); i++)
    {
        long long num{nums[i]};
        // order matters because I can't overwrite the dependencies before using them
        triplet = std::max(triplet, doublet * num);
        doublet = std::max(doublet, singlet - num);
        singlet = std::max(singlet, num);
    }

    return triplet;
}

/*2873. Maximum Value of an Ordered Triplet I

1005 time, 5% memory (I can definitely space reduce but i am too lazy rn)

DP solution (technically 2D DP but it's super simple)

The maximum triplet value at index i is computed based on current state and 2 choices
State: number of integers already being used
Choices: Take index i (adding 1 to State), Skip index i (maintaining State)

Parameters:
- s State
- i index of nums
*/
long long maximumTripletValue(std::vector<int> &nums)
{
    const size_t n{nums.size()};
    const int states{3};
    // to initialize a 2D std::vector named dp with rows rows and cols columns, all initialized to zero, the most recommended approach is:
    // here states is rows, n is cols
    std::vector<std::vector<long>> dp(states + 1, std::vector<long>(n, 0));
    // init (1,0) to avoid a ternary in the loop
    dp[1][0] = nums[0];

    // DP loop
    for (size_t i = 1; i < n; i++)
    {
        long num{nums[i]};
        // first number of triplet
        dp[1][i] = std::max(dp[1][i - 1], num);
        // second number of triplet (subtraction term)
        dp[2][i] = std::max(dp[2][i - 1], dp[1][i - 1] - num);
        // third number (multiplication term)
        dp[3][i] = std::max(dp[3][i - 1], dp[2][i - 1] * num);
    }

    // for (std::vector<int> row : dp)
    // {
    //     for (int element : row)
    //     {
    //         std::cout << element << " ";
    //     }
    //     std::cout << '\n';
    // }
    return dp[3][n - 1];
}

/* 31% time, 66% memory
This is the brute force-esque solution
Since al numbers in nums are positive, the moment you have a negative diff, skip */
long long maximumTripletValueBruteForce(std::vector<int> &nums)
{
    long long result{0};
    for (size_t i = 0; i < nums.size(); i++)
    {
        long long diff{0};
        for (size_t j = i + 1; j < nums.size(); j++)
        {
            if (nums[i] - nums[j] < diff)
                continue;
            long long newDiff{nums[i] - nums[j]};
            for (size_t k = j + 1; k < nums.size(); k++)
            {
                result = std::max(result, newDiff * nums[k]);
            }
            diff = std::max(diff, newDiff);
        }
    }
    return result;
}

/* 368. Largest Divisible Subset
74% time, 12% memory

I will sort the input, I can't figure out how without the sort
2 data structures needed:
- Array to save the max length ending at `i`
- Array to save the direct child index leading to the max length at `i`
*/
std::vector<int> largestDivisibleSubset(std::vector<int> &nums)
{
    size_t n{nums.size()};
    std::vector<int> maxL(n, 1);
    // out of bounds idx since I can't use negative numbers:
    std::vector<size_t> child(n, n);
    size_t idx{0};
    for (size_t i = 0; i < n; i++)
    {
        for (size_t j = 0; j < i; j++)
        {
            if ((nums[i] % nums[j] == 0) && (maxL[j] + 1 > maxL[i]))
            {
                child[i] = j;
                maxL[i] = maxL[j] + 1;
            }
        }
        if (maxL[i] > maxL[idx])
        {
            idx = i;
        }
    }

    std::vector<int> result;
    while (idx != n)
    {
        result.emplace_back(nums[idx]);
        idx = child[idx];
    }
    return result;
}

int main()
{
    std::vector<int> result;
    std::vector<int> nums;

    std::cout << "2873. Maximum Value of an Ordered Triplet I\n";
    nums = {12, 6, 1, 2, 7};
    std::cout << maximumTripletValue(nums) << '\n';
    nums = {1000000, 1, 1000000}; // forces me to use long
    std::cout << maximumTripletValue(nums) << '\n';

    std::printf("\n\nLargest Divisible Subset\n");
    nums = {5, 9, 18, 54, 108, 540, 90, 180, 360, 720};
    result = largestDivisibleSubset(nums);
    for (int val : result)
    {
        std::printf("%d ", val);
    }
    std::cout << '\n';

    return 0;
}
