<?php

namespace SDK\Service\Greeter;

use Exception;
use Grpc\ChannelCredentials;
use SDK\Api\Common\RequestHeader;
use SDK\Api\Greeter\HelloClient;
use SDK\Api\Greeter\HelloRequest;
//use SDK\Api\Greeter\HelloReply;
use SDK\Utils\Singleton;

class Hello
{
    use Singleton;

    private function client()
    {
        static $client;

        $hostname = '127.0.0.1:8080';
        $opts = [
            'credentials' => ChannelCredentials::createInsecure(),
        ];
        $channel = null;

        if (is_null($client)) {
            $client = new HelloClient($hostname, $opts, $channel);
        }
        return $client;
    }

    public function say(string $name, $timeout = null): string
    {
        // curl_init();
        // curl_exec($ch, $url);
        $request = new HelloRequest();
        $request->setName($name);
        $header = new RequestHeader();
        $request->setHeader($header);

        // HelloReply
        [$reply, $status] = $this->client()->say($request)->wait();

        /**
         * status object
         * object(stdClass)#15 (3) {
         *  ["metadata"]=>
         *  array(0) {
         *  }
         *  ["code"]=>
         *  int(14)
         *  ["details"]=>
         *  string(35) "DNS resolution failed for service: "
         * }
         */
        if ($status) {
            //TODO log
            throw new Exception($status->details, $status->code);
        }

        if ($reply->getHeader()->getCode()) {
            //TODO log
            var_dump($reply->getHeader());
            exit;
            throw new Exception($reply->getHeader()->getMessage(), $reply->getHeader()->getCode());
        }

        return $reply->getMessage();
    }
}
