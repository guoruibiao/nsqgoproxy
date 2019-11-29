<?php

require "./DemoService.php";
$rawdata = file_get_contents('php://input');
$rows = explode("=", $rawdata);
foreach($rows as $key=>$data) {
    if($key == "data") {
        $data = json_decode(urldecode($rows[1]), true);
        var_dump($data);
        $classname = $data["Classname"];
        $object = new $classname();
        $ret = call_user_func_array(array($object, $data["MethodName"]), $data["Parameters"][0]);
        file_put_contents("./nsq-worker.log", json_encode($ret));
    }
}
