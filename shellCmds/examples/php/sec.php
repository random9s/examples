<?php
    $bin = "ls ";
    $userIn = "tmp | wc -l; rm -rf target/";
    $cmd = $bin.$userIn;
    //This command will echo what we want, then delete the directory without any notice 
    echo(shell_exec($cmd));
?>
