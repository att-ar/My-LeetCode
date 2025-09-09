use std::collections::BTreeMap;

/// 100% time
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut bmap = BTreeMap::<i32, usize>::new();
    for (i, n) in nums.into_iter().enumerate() {
        match bmap.get(&n) {
            Some(&x) => {
                return vec![i as i32, x as i32];
            }
            None => {
                bmap.insert(target - n, i);
            }
        }
    }
    vec![] // Unreachable due to Leetcode invariant
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        println!("{:?}", two_sum(vec![2, 7, 11, 15], 9));
        println!(
            "{:?}",
            two_sum(
                vec![2, 120, 12312, 3, 11, 3, 4, 3, 1, 1, 2, 3, 7, 11, 15],
                9
            )
        );
        println!("{:?}", two_sum(vec![3, 2, 4], 6));
        println!("{:?}", two_sum(vec![3, 3], 6));
    }
}
