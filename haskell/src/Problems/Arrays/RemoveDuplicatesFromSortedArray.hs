module Problems.Arrays.RemoveDuplicatesFromSortedArray (doRemoveDuplicatesFromSortedArray) where

import Text.Printf (printf)
import Data.List (foldl')

doRemoveDuplicatesFromSortedArray :: IO ()
doRemoveDuplicatesFromSortedArray = do
    putStrLn "\n26. Remove Duplicates from Sorted Array:"
    let widthCall :: Int
        widthCall = 50
        widthResult :: Int
        widthResult = 15
    printf "%-*s = %*s\n" widthCall "[0,0,1,1,1,2,2,3,3,4]" widthResult (show (removeDuplicatesFromSortedArray [0,0,1,1,1,2,2,3,3,4]))
    printf "%-*s = %*s\n" widthCall "[0,0,1,1,1,2,2,3,3,4,4,4,5,6,8,8,101,2021]" widthResult (show (removeDuplicatesFromSortedArrayReturned [0,0,1,1,1,2,2,3,3,4,4,4,5,6,8,8,101,2021]))

{- 26. Remove Duplicates from Sorted Array
Problem states to mutate inplace and return the number of distinct
I will start with just returning the number of distinct
-}
removeDuplicatesFromSortedArray :: [Int] -> Int
removeDuplicatesFromSortedArray = snd . foldl' step (Nothing, 0)
    -- `snd . foldl' ...` is just `snd (foldl' step ...)` the `.` is the composition operator; (f . g) x is f (g x)
    where
        step :: (Maybe Int, Int) -> Int -> (Maybe Int, Int)
        step (Nothing, curCount) x = (Just x, curCount + 1) -- start of the list init 1 distinct
        step (Just prev, curCount) x
            | x == prev = (Just prev, curCount) -- found a duplicate
            | otherwise = (Just x, curCount + 1) -- found a new element

{- Can't use foldl' because I can't reduce anymore, I need to recursively build the list -}
removeDuplicatesFromSortedArrayReturned :: [Int] -> ([Int], Int)
removeDuplicatesFromSortedArrayReturned = dfs Nothing
    where
        dfs :: Maybe Int -> [Int] -> ([Int], Int)
        dfs _ [] = ([], 0)
        dfs Nothing (x:xs) =
            let (rest, count) = dfs (Just x) xs
            in (x:rest, count + 1)
        dfs (Just y) (x:xs)
            | x == y = dfs (Just x) xs -- skip duplicate
            | otherwise =
                let (rest, count) = dfs (Just x) xs
                in (x:rest, count + 1)
