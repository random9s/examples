## Shelled Commands in Golang
* Explore the high level os/exec package 
    - Look at creating / killing a process
    - Talk about blocking and non-blocking while running shelled commands
    - Look at Output() function
* Show / Explain a few simple commands  
    - Running unix commands
    - Running scripts
    - Running other programs (ffmpeg, avconv)
* Using context to kill a process
    - os.Process.Kill

## A Deeper Look 
* Note: This part can be ignored for anyone who is not interested in how the language implements this
* Language Comparison / Design implementations
    - PHP Implementation
    - Haskell Implementation
* Using StartProcess
    - When / Why use StartProcess as opposed to os/exec 
* Using syscall
    - No reason to do this, but it's interesting to understand how things work
* Challenge (maybe?) 
    - Create your own high level wrapper for StartProcess

## Caveats
* Joining multiple commands using os.Pipe
    - Example using anonymous pipe | in command
    - Example using os.Pipe
    - Example using bash -c <string>
* Error handling
    - Does Go provide easy access to error messages from binaries?
    - What do the exit statuses mean?

## Conclusion
