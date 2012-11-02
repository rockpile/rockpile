import System.IO
import System.Environment
import Data.Fixed

main :: IO ()
main = do
  args <- getArgs
  let item = fromIntegral (read $ head args :: Int)
      xs = case (mod' item 4) of 3 -> item
                                 x -> item - x + 3
      pi = 4 * (addItem xs) 
  putStrLn $ show pi

addItem :: (RealFloat a) => a -> a
addItem a
  | a > 4     = 1 / (a - 2) - (1 / a) + (addItem (a - 4))
  | a == 3    = 1 - 1 / a
  | otherwise = 0

addItem_loop :: (RealFloat a) => a -> a
addItem_loop a
  | a > 4     = 1 / (a - 2) - (1 / a)
  | a == 3    = 1 - 1 / a
  | otherwise = 0

addItem_help :: (RealFloat a) => a -> a -> a
addItem_help a b
  | b < 4     = a
  | otherwise = addItem_help (a + addItem_loop b) (b - 4)
                   
  