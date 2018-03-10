## Shelled Commands in Golang
* Explore the high level os/exec package 
    - Look at creating / killing a process
    - Talk about blocking and non-blocking while running shelled commands
    - Look at Output() function
* Show / Explain a few simple commands  
    - Running unix commands
    - Running scripts
    - Running other programs (ffmpeg, avconv)
* Joining multiple commands using os.Pipe
    - Example using anonymous pipe | in command
    - Example using os.Pipe
    - Example using bash -c <string>
    - Remember that Golang is a web server programming langauge first.
* Error handling
    - Creating and reading from stderr pipe
    - Killing process (using context)
    - Does Go provide easy access to error messages from binaries?
    - All I get is this exit status, what's that?

## A Deeper Look 
* Note: This part can be ignored for anyone who is not interested in how the language implements this
* Language Comparison / Design implementations
    - PHP Implementation
    - Haskell Implementation
    - Wherever there is a possibility for user input, you have to be aware of injections.  
* Explore the lower level os package
    - Calling StartProcess 
    - When / Why use StartProcess as opposed to os/exec 
* Using syscall
    - No reason to do this, but it's interesting to understand how things work
* Challenge (maybe?) 
    - Create your own high level wrapper for StartProcess

## Caveats
* Performance
    - Generally slower
* Security
    - Injections 
    - Only call hard coded commands, don't allow commands to be run with user input
* Portability
    - Do you have a backup func if the binary you're trying to execute doesn't exist on a specific OS?

## Conclusion
