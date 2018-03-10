import System.Process
main = do
    callCommand "ls tmp | wc -l"
