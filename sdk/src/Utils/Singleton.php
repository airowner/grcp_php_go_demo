<?php

namespace SDK\Utils;

trait Singleton
{
    public static function ins()
    {
        return new static;
    }

    private function __construct()
    {
    }

    private function __clone()
    {
    }
}
