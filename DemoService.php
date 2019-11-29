<?php

class DemoService {
    public function say($message) {
        $retstring = "hello {$message}, this is from nsqworker.";
        echo $retstring;
        return $retstring;
    }
}
