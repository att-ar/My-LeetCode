module Main (main) where

import Problems.Arrays.Arrays (arrays)

-- import Problems.BinarySearch
-- import Problems.LinkedList (ListNode(..), reverseList) -- If you have custom data types

main :: IO ()
main = do
    arrays

-- Helper function to show a ListNode (assuming you have this defined)
-- showList :: Maybe ListNode -> String
-- showList Nothing = "Nil"
-- showList (Just (ListNode val next)) = show val ++ " -> " ++ showList next
