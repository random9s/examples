<?php
    $cmd = "ls tmp | wc -l";
    echo(shell_exec($cmd));
?>
