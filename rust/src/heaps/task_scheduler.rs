// 621. Task Scheduler
use std::collections::BinaryHeap;
use std::collections::VecDeque;

/// 28% time, 18 % memory
/// I couldn't figure out the counting method so I'm just gonna do the heap
///
/// There's a faster way to do it using a wait queue so that you avoid repushing
/// when task is supposed to end before or at current time
///
/// I have that solution written below.
/// One thing I notice: using the vector for counting instead of the hashmap makes quite the difference.
pub fn least_interval(tasks: Vec<char>, n: i32) -> i32 {
    let count: Vec<i32> = {
        let a = 'A' as usize;
        let mut count: Vec<i32> = vec![0; 26];
        tasks.into_iter().for_each(|c| count[c as usize - a] += 1);
        let mut count: Vec<i32> = count.into_iter().filter(|&freq| freq > 0).collect();
        count.sort();
        count.reverse();
        count
    };

    let mut heap: BinaryHeap<(i32, i32)> = count
        .into_iter()
        .enumerate()
        .map(|(idx, freq)| (-(idx as i32) - 1, freq - 1)) // end time (minHeap), freq remaining (maxHeap)
        .collect();

    let mut cur_time = 0;
    while let Some((end, remaining)) = heap.pop() {
        // if the task is supposed to end before or at the current Time
        // repush it with a potentially valid end time
        if end >= cur_time {
            heap.push((cur_time - 1, remaining));
            continue;
        }
        cur_time = end;

        if remaining > 0 {
            heap.push((cur_time - (n + 1), remaining - 1));
        }
    }
    -cur_time
}

/// 100% time, 18% memory - using `while let Some(&(ready_time...))`
/// 72% time, 84% memory - using `if let Some(&(ready_time...))...`
pub fn least_interval_faster(tasks: Vec<char>, n: i32) -> i32 {
    // max heap of task frequencies
    let mut heap: BinaryHeap<i32> = {
        let a = 'A' as usize;
        let mut count: Vec<i32> = vec![0; 26];
        tasks.into_iter().for_each(|c| count[c as usize - a] += 1);
        count.into_iter().filter(|&freq| freq > 0).collect()
        // Don't need to sort and reverse since i am not using the end time as the ordering
    };

    // wait queue initialized with number of tasks occurring more than once
    // (time when ready, remaining freq)
    let mut wait_queue =
        VecDeque::<(i32, i32)>::with_capacity(heap.iter().filter(|&&freq| freq > 0).count());

    // start at time zero, each loop is one timestamp to do an action
    let mut cur_time = 0;
    while !heap.is_empty() || !wait_queue.is_empty() {
        // Move everything from the wait queue, that is ready, back into the heap
        // Note that the ready_time is in non-decreasing order based how i push to the queue
        if let Some(&(ready_time, remain)) = wait_queue.front()
            && ready_time <= cur_time
        {
            heap.push(remain);
            wait_queue.pop_front();
        }

        if let Some(remain) = heap.pop()
            && remain > 1
        {
            wait_queue.push_back((cur_time + n + 1, remain - 1));
        }

        cur_time += 1
    }

    cur_time
}

#[cfg(test)]
pub mod tests {
    use super::*;

    #[test]
    fn test_least_interval() {
        println!("{}", least_interval(vec!['A', 'A', 'A', 'B', 'B', 'B'], 2));
        println!(
            "{}",
            least_interval(vec!['B', 'C', 'D', 'A', 'A', 'A', 'A', 'G', 'H', 'L'], 1)
        );

        println!(
            "{}",
            least_interval_faster(vec!['A', 'A', 'A', 'B', 'B', 'B'], 2)
        );
        println!(
            "{}",
            least_interval_faster(vec!['B', 'C', 'D', 'A', 'A', 'A', 'A', 'G', 'H', 'L'], 1)
        );
    }
}
