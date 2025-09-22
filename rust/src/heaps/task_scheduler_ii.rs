// 2365. Task Scheduler 2

use std::collections::HashMap;

/// 50% time, 75% memory
/// Seems pretty different from the first version since the tasks are ordered
pub fn task_scheduler_ii(tasks: Vec<i32>, space: i32) -> i64 {
    let mut wait: HashMap<i32, i64> = HashMap::new();
    let mut cur_time = 0;
    let k = (space + 1) as i64;
    for task in tasks {
        // Check if need to wait
        if let Some(&ready_time) = wait.get(&task) {
            cur_time = cur_time.max(ready_time);
        }

        // Insert into wait map
        wait.insert(task, cur_time + k);
        cur_time += 1;
    }

    cur_time
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_least_interval() {
        println!("{}", task_scheduler_ii(vec![1, 2, 1, 2, 3, 1], 3));
        println!("{}", task_scheduler_ii(vec![5, 8, 8, 5], 2));
    }
}
