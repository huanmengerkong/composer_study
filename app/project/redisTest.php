<?php
$redis = new Redis();
$result = $redis->connect('redis', 6379);
if ($result) {
	echo '<center><h2>成功通过 PHP 连接到 Redis </h2></center>' . PHP_EOL;
}
// $redis->auth('123456');
$redis->set('key1', 'val1');
echo '<center><h2>Set Redis: key1 = ' . $redis->get('key1') . '</h2></center>' . PHP_EOL;
