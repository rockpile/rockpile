import System.IO
import Control.Concurrent
import Control.Concurrent.Chan

main :: IO ()
main =
  do a <- newChan
     pid <- forkIO $ write_data a "new"
     putStrLn $ show pid
     str <- readChan a
     putStrLn str

write_data :: Chan a -> a -> IO ()
write_data ch d = do writeChan ch d
        