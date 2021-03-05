<?php

use SDK\Service\Greeter\Hello;

include "vendor/autoload.php";

try {
    Hello::ins()->say("eeo");
} catch (Throwable $e) {
    var_dump($e->getCode(), $e->getMessage());
}
