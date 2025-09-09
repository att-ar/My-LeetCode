module MyLib (atoiSafe, atoi, itoa, time) where

import qualified Data.Text as T
import qualified Data.Text.Read as TR
import System.CPUTime ( getCPUTime )
import Text.Printf (printf)
-- import Data.List (nub)
import Control.DeepSeq (force, NFData)
import Control.Exception (evaluate)


-- Safe conversion using decimal (returns Either with error message or result with remaining text)
atoiSafe :: T.Text -> Either String (Int, T.Text)
atoiSafe = TR.decimal

-- A more practical version that ignores remaining text
-- atoi :: T.Text -> Either String Int
-- atoi t = fst <$> TR.decimal t

atoi :: T.Text -> Int
atoi t = case TR.decimal t of
    Right (num, _) -> num
    Left err -> error $ "Fatal error: Cannot parse '" ++ T.unpack t ++ "' as Int: " ++ err

-- Convert Int to Text
itoa :: Int -> T.Text
itoa = T.pack . show


-- General benchmark function
time :: NFData a => String -> a -> IO ()
time label x = do
    start <- getCPUTime
    _ <- evaluate (force x)
    end <- getCPUTime
    let diff = fromIntegral (end - start) / (10^12)
    printf "%s: %0.6f sec\n" label (diff :: Double)
