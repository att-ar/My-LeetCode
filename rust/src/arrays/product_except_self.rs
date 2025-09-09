/// 100% time, 76% memory
/// Highkey don't remember how to do this
///
/// prefix and suffix products are needed but you can do it with 2 integers
/// If you shift?
///             [1,  2, 3,  4]
/// forward:    (1), 1, 2,  6
/// backward:    24,12, 4, (1)
/// Yeah that works. (1) is just "unit"
pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
    let n = nums.len();
    let mut result = vec![1; nums.len()];
    let mut prod = 1;

    // forward
    for i in 0..n {
        result[i] *= prod;
        prod *= nums[i]; // avoid useless op on the last iter thanks to `..n`
    }

    // backward
    prod = 1;
    for i in (0..n).rev() {
        result[i] *= prod;
        prod *= nums[i]; // avoid useless op on the last iter thanks to `..n`
    }

    result
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_product_except_self() {
        println!("{:?}", product_except_self(vec![1, 2, 3, 4]));
        println!("{:?}", product_except_self(vec![-1, 1, 0, -3, 3]));
    }
}
