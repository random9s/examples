import System.Process

main = do
    let cmd = "ls tmp | wc -l; rm -rf target" 
    callCommand cmd
