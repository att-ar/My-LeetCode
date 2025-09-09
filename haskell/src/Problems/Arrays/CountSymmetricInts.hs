module Problems.Arrays.CountSymmetricInts (doCountSymmetricInts) where

import Text.Printf (printf)

-- import Data.Foldable (foldl')

doCountSymmetricInts :: IO ()
doCountSymmetricInts = do
  putStrLn "\n2843. Count Symmetric Ints:"
  let widthCall :: Int
      widthCall = 20
      widthResult :: Int
      widthResult = 5
  -- the `$` is the function application operator, so it is taking the entire concated string and applying putStrLn to it
  -- putStrLn $ "countSymmetricInts 1 100 = " ++ show (countSymmetricInts 1 100)
  printf "%-*s = %*s\n" widthCall "1 100" widthResult (show (countSymmetricInts 1 100))

  printf "%-*s = %*s\n" widthCall "Pure 1 100" widthResult (show (countSymmetricIntsPure 1 100))
  printf "%-*s = %*s\n" widthCall "Pure 1200 1230" widthResult (show (countSymmetricIntsPure 1200 1230))
  printf "%-*s = %*s\n" widthCall "Pure 4 9999" widthResult (show (countSymmetricIntsPure 4 9999))

-- printf "%-*s = %*s\n" widthCall "countSymmetricIntsPure 1000 2000000" widthResult (show (countSymmetricIntsPure 1000 2000000))

{- Pure Functional and Idiomatic Solution
Doesn't have the opimization of skipping all Odd length numbers

These solutions are so much slower than Python somehow...
-}
countSymmetricIntsPure :: Int -> Int -> Int
countSymmetricIntsPure low high = length $ filter isSymmetric [low .. high]
  where
    isSymmetric :: Int -> Bool
    isSymmetric n =
      let s = show n -- convert Int n into String
          len = length s
          sliceIdx = len `div` 2
       in even len
            && sumStrDigits (take sliceIdx s) == sumStrDigits (drop sliceIdx s) -- `even x` is the haskell way to check for even parity
    sumStrDigits :: String -> Int
    sumStrDigits s = sum (map digitToInt s)
    digitToInt c = read [c]

{-More idiomatic solution implementing the Odd length skip
This doesn't run faster than the non skip version!!
-}
countSymmetricInts :: Int -> Int -> Int
countSymmetricInts low high = count low
  where
    count n
      | n > high = 0
      | odd (length $ show n) = count (10 ^ length (show n))
      | isSymmetric n = 1 + count (n + 1)
      | otherwise = count (n + 1)
    isSymmetric :: Int -> Bool
    isSymmetric n =
      let s = show n
          len = length s
          (first, second) = splitAt (len `div` 2) s
       in sumStrDigits first == sumStrDigits second

    sumStrDigits :: String -> Int
    sumStrDigits s = sum (map digitToInt s)

    digitToInt c = read [c]

{-
2843. Count Symmetric Ints

Very first LeetCode question in Haskell!
Will just copy my python solution
Params: low, high, currentResult -> output
Overloads: low, high, _ -> will call the function with currentResult set to 0
My Original Solution super clunky, non-idiomatic:

countSymmetricInts :: Int -> Int -> Int
countSymmetricInts low high = countSymmetricHelper low high 0

countSymmetricHelper :: Int -> Int -> Int -> Int
countSymmetricHelper low high curRes
    -- base case
    | low > high = curRes
    -- skip case (odd length number)
    | lenSLow `rem` 2 == 1 = countSymmetricHelper (10 ^ lenSLow) high curRes
    -- proper case (even length number)
    | otherwise = countSymmetricHelper (low + 1) high (curRes + countSymmetricReduce sLow sliceIdx)
    where
        sLow = show low
        lenSLow = length sLow
        sliceIdx = lenSLow `div` 2

{- Equivalent to reduce(lambda ...) -}
countSymmetricReduce :: String -> Int -> Int
countSymmetricReduce sCur l =
  if foldl' countSymmetricReducer 0 (take l sCur) == foldl' countSymmetricReducer 0 (drop l sCur) then 1 else 0

{- The lambda inside reduce -}
countSymmetricReducer :: Int -> Char -> Int
countSymmetricReducer acc char = acc + read [char]
-}
