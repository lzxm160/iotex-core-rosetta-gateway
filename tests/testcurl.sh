#!/usr/bin/env bash
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},
    "account_identifier": {
        "address": "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
    }}' http://127.0.0.1:8080/account/balance
#response:
#{"block_identifier":{"index":4051388,"hash":"b2a68ad3dcc3c61ad6a0cec3431400a40a195def87251c9c60920d7fcf973f2f"},"balances":[{"value":"11989999999999999999","currency":{"symbol":"IOTX","decimals":18}}],"metadata":{"nonce":1}}


curl -X POST --data '{"metadata": {}}' http://127.0.0.1:8080/network/list
#response:
#{"network_identifiers":[{"blockchain":"IoTeX","network":"testnet"}]}

curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},"metadata": {}}' http://127.0.0.1:8080/network/options
#response:
#{"version":{"rosetta_version":"1.3.5","node_version":"v1.0.0"},"allow":{"operation_statuses":[{"status":"success","successful":true},{"status":"fail","successful":false}],"operation_types":["fee","transfer","execution","depositToRewardingFund","claimFromRewardingFund","stakeCreate","stakeWithdraw","stakeAddDeposit","candidateRegister"],"errors":[{"code":1,"message":"unable to get chain ID","retriable":true},{"code":2,"message":"invalid blockchain specified in network identifier","retriable":false},{"code":3,"message":"invalid sub-network identifier","retriable":false},{"code":4,"message":"invalid network specified in network identifier","retriable":false},{"code":5,"message":"network identifier is missing","retriable":false},{"code":6,"message":"unable to get latest block","retriable":true},{"code":7,"message":"unable to get genesis block","retriable":true},{"code":8,"message":"unable to get account","retriable":true},{"code":9,"message":"blocks must be queried by index and not hash","retriable":false},{"code":10,"message":"invalid account address","retriable":false},{"code":11,"message":"a valid subaccount must be specified ('general' or 'escrow')","retriable":false},{"code":12,"message":"unable to get block","retriable":true},{"code":13,"message":"operation not implemented","retriable":false},{"code":14,"message":"unable to get transactions","retriable":true},{"code":15,"message":"unable to submit transaction","retriable":false},{"code":16,"message":"unable to get next nonce","retriable":true},{"code":17,"message":"malformed value","retriable":false},{"code":18,"message":"unable to get node status","retriable":true}]}}

curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},"metadata": {}}' http://127.0.0.1:8080/network/status
#response:
#{"current_block_identifier":{"index":4051411,"hash":"17ef91a70d8d25a0396e4cb6da669e1ffb217dc00436effbba3f6ccea1e53a39"},"current_block_timestamp":1592884250000,"genesis_block_identifier":{"index":1,"hash":"663fc0a40a4943f1b56f501aee3ad626b5396e850aa53c5bd8759d0d47694dfc"},"peers":null}

curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"
    },
    "options": {"id":"io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"}
}' http://127.0.0.1:8080/construction/metadata
#response:
#{"metadata":{"nonce":{"Nonce":1,"Balance":"11989999999999999999"}}}

curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},
    "signed_transaction": "0a470801100118c0843d220d31303030303030303030303030522e0a01311229696f316c397661716d616e776a3437746c7270763665746633707771307330736e73713476786b6532124104ea8046cf8dc5bc9cda5f2e83e5d2d61932ad7e0e402b4f4cb65b58e9618891f54cba5cfcda873351ad9da1f5a819f54bba9e8343f2edd1ad34dcf7f35de552f31a41d53b8aa4b0165326dcf2eddf4da1fcba8e864f805426c73ee7e73748713c48774bf117f7e78f18459645386ecbb644ca3cca89069920b20ff405768d3d1d6bb301"
}' http://127.0.0.1:8080/construction/submit
#response:
#{"transaction_identifier":{"hash":"292cda920534be56c78d6f13686dc7dbb94b77714b93abefb9f1e18679e2ae27"}}

# transfer action
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},
    "block_identifier": {"index": 390873}}' http://127.0.0.1:8080/block
#response:
#{"block":{"block_identifier":{"index":390873,"hash":"5c084459315fcf0839ed9f2d8b89ca8fb039695a56007a071e5ce9d3c8908d95"},"parent_block_identifier":{"index":390872,"hash":"3ae76de97535f4908d7dd6b2d5f232543b1e5a9fe80a0e9d8f91fdd27d9363eb"},"timestamp":1573620900000,"transactions":[{"transaction_identifier":{"hash":"b37d5db44bd3dc182617b56744e12cab94486808eae1dc401599b611ed388164"},"operations":[{"operation_identifier":{"index":0},"type":"fee","status":"success","account":{"address":"io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"},"amount":{"value":"-10000000000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":1},"type":"transfer","status":"success","account":{"address":"io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"},"amount":{"value":"-10000000000000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":2},"type":"transfer","status":"success","account":{"address":"io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"},"amount":{"value":"10000000000000000000","currency":{"symbol":"IOTX","decimals":18}}}]}]}}

# Execution multisend,from mainnet
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "mainnet"},
    "block_identifier": {"index": 5542366}}' http://127.0.0.1:8080/block
#response
#{"block":{"block_identifier":{"index":5542366,"hash":"0379f5ea19f5e69df911c1bd2fbbe89429e2358fd573ebf36b432c98661dcfb7"},"parent_block_identifier":{"index":5542365,"hash":"2fe0e7f8adb3fafbaa91bbe86f93ceaddb61605534bb887ca54cd624ac2ec5aa"},"timestamp":1592942380000,"transactions":[{"transaction_identifier":{"hash":"10b5e5ee554c02980277873f3cf80a0190c611e77e19d58b4161110beefb643b"},"operations":[{"operation_identifier":{"index":0},"type":"fee","status":"success","account":{"address":"io1yftf6xuar26lsaher0qa73d0zc3csm884hjwvh"},"amount":{"value":"-797150000000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":1},"type":"execution","status":"success","account":{"address":"io1yftf6xuar26lsaher0qa73d0zc3csm884hjwvh"},"amount":{"value":"-93385016309186326947000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":2},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-2574859143759090000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":3},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-125766229055977000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":4},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-2943587464794670000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":5},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-518429045809640000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":6},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-656778422420071000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":7},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-709082192569296000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":8},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-46149204079090300000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":9},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-1036150252248720000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":10},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-14005032093352100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":11},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-4774353096089010000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":12},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-23059631035726500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":13},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-250970570046752000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":14},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-100599772697421000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":15},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-246491560521753000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":16},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-2374904506079770000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":17},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-21693027340822100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":18},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-92146377765141200000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":19},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-5370824911487300000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":20},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-2686948200006130000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":21},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-10038822801870100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":22},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-7469076435667700000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":23},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-15233287250867600000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":24},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-74163610412132300000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":25},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-10387513134484700000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":26},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-52190993935752000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":27},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-115182972206426000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":28},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-971614308125297000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":29},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-2530889760003160000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":30},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-20077645603740200000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":31},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-19620309947250500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":32},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-19620309947250500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":33},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-8308033225301630000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":34},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-23074602039545100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":35},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-462588954710173000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":36},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-419124856180216000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":37},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-28400220190272100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":38},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-69109783323855900000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":39},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-714289516214088000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":40},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-35359286853399500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":41},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-2688584395541960000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":42},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-154504506430765000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":43},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-1754101284134070000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":44},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-57391033173728000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":45},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-876356512450687000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":46},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-277805536152447000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":47},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-4676502542873220000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":48},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"-18987396723145600000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":49},"type":"execution","status":"success","account":{"address":"io1u0pttscsz2hkyuugdzew6zg0kmg6xm7kd7kzhw"},"amount":{"value":"57391033173728000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":50},"type":"execution","status":"success","account":{"address":"io1shgp88ucfsrqg76maqvt7sky0t3jg0ylf0kydf"},"amount":{"value":"709082192569296000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":51},"type":"execution","status":"success","account":{"address":"io1s6mfntw5882yeus2m88lqkmykythjnecr7dd9z"},"amount":{"value":"125766229055977000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":52},"type":"execution","status":"success","account":{"address":"io1ntaumzzqekkswla6hcpljd77suc8dg7r650keq"},"amount":{"value":"246491560521753000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":53},"type":"execution","status":"success","account":{"address":"io1nrusmxggfqwn63c6g05kl6f392hg2yrdjq8rcz"},"amount":{"value":"23059631035726500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":54},"type":"execution","status":"success","account":{"address":"io1n9jy6xnehrlktxccxzg0e44adxmn73zuv8pl0w"},"amount":{"value":"250970570046752000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":55},"type":"execution","status":"success","account":{"address":"io1n8av5f9vaj0ysnf294vgyys0dapdxvnc97s6ff"},"amount":{"value":"100599772697421000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":56},"type":"execution","status":"success","account":{"address":"io1mz7hpkherhdn06l2q52g2tlaqxg2gzd8yastlt"},"amount":{"value":"714289516214088000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":57},"type":"execution","status":"success","account":{"address":"io1mxlk9cp6yu4gcpl5kg7ulucga835ljcfqp6gkv"},"amount":{"value":"35359286853399500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":58},"type":"execution","status":"success","account":{"address":"io1mwykqzm45wpjfkxtvzwuvzpfczh9pm33x3ht2x"},"amount":{"value":"154504506430765000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":59},"type":"execution","status":"success","account":{"address":"io1mt7dma8gve52pw7r9y0vp87z0a6240xsjcqke4"},"amount":{"value":"2688584395541960000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":60},"type":"execution","status":"success","account":{"address":"io1mscvxapc0n78gmrm3nsrlv07klj5e7365c64ec"},"amount":{"value":"1754101284134070000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":61},"type":"execution","status":"success","account":{"address":"io1lvemm43lz6np0hzcqlpk0kpxxww623z5hs4mwu"},"amount":{"value":"93385016309186326947000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":62},"type":"execution","status":"success","account":{"address":"io1lq7egn36nmcympgpwzeffragys0wd9m6dhjx84"},"amount":{"value":"18987396723145600000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":63},"type":"execution","status":"success","account":{"address":"io1kwwqpu9d8pvp9kg9xvm4g8wka26cu2d0rnfmwk"},"amount":{"value":"52190993935752000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":64},"type":"execution","status":"success","account":{"address":"io1kfhua8r7m2mv99rx65v073av796ne5j8r9azmw"},"amount":{"value":"10387513134484700000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":65},"type":"execution","status":"success","account":{"address":"io1k0t4lxp9psykwku7wy27y7kpzxn6hsesqvgjty"},"amount":{"value":"115182972206426000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":66},"type":"execution","status":"success","account":{"address":"io1jem747ntxrg6gumfcjrf5vqp3lftcghh7p0c9v"},"amount":{"value":"4774353096089010000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":67},"type":"execution","status":"success","account":{"address":"io1j8mt8sl9g7tg84nvwwp5jpve6zcrjuzmqreslc"},"amount":{"value":"14005032093352100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":68},"type":"execution","status":"success","account":{"address":"io1hvs50m994k0h2xr7awkcec3lqg7pgcvaqu0n3s"},"amount":{"value":"2530889760003160000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":69},"type":"execution","status":"success","account":{"address":"io1hn6vaewrdggnqggvcehhnhye3rhw62p3zdczz2"},"amount":{"value":"20077645603740200000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":70},"type":"execution","status":"success","account":{"address":"io1hegc8lql0ag5uhy25w4tma9ht8783jkdh5udfn"},"amount":{"value":"971614308125297000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":71},"type":"execution","status":"success","account":{"address":"io1have0dwme4zjsfuv2vhs796j3qs4sxghpqthc2"},"amount":{"value":"19620309947250500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":72},"type":"execution","status":"success","account":{"address":"io1h5um8en82m5emwk5gl9zyzkpchwtk2juzm5az7"},"amount":{"value":"19620309947250500000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":73},"type":"execution","status":"success","account":{"address":"io1e0lcengs3ev6wnna4kstqvvjfnt3lpad5skj6t"},"amount":{"value":"462588954710173000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":74},"type":"execution","status":"success","account":{"address":"io1cfhmscmuffnfptk9v8md5j96wdx04xsujtavxr"},"amount":{"value":"8308033225301630000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":75},"type":"execution","status":"success","account":{"address":"io1c66s432ru5eh5qz82a4cvzqs5pv8fhx4lpamt9"},"amount":{"value":"23074602039545100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":76},"type":"execution","status":"success","account":{"address":"io1amjlf53j6r05g48j7xq85wdczhdz230zfspzuv"},"amount":{"value":"876356512450687000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":77},"type":"execution","status":"success","account":{"address":"io17h84l26352s78hmjttyvpng6wam92ghra5kjpt"},"amount":{"value":"4676502542873220000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":78},"type":"execution","status":"success","account":{"address":"io17d2wfh0rzt376tx5ex0mmyavq7n4djnuwkw7mx"},"amount":{"value":"277805536152447000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":79},"type":"execution","status":"success","account":{"address":"io16unehh4mu357d8few0eryvyj48upxu47suhtnw"},"amount":{"value":"28400220190272100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":80},"type":"execution","status":"success","account":{"address":"io16lvsxw7rymw4l39yew509w6hwpsgk9u8mamkq9"},"amount":{"value":"69109783323855900000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":81},"type":"execution","status":"success","account":{"address":"io168wqnnrgwu34lyjaf8pprszvwv7frddzrte2hd"},"amount":{"value":"419124856180216000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":82},"type":"execution","status":"success","account":{"address":"io15r7xjxtnmn6unvjz56ht9g2e8p9u00akz66ghe"},"amount":{"value":"21693027340822100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":83},"type":"execution","status":"success","account":{"address":"io15pv9a7d8559garn36knss9uwxpwtpd25x0v6f7"},"amount":{"value":"2374904506079770000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":84},"type":"execution","status":"success","account":{"address":"io15hkae5vjtklxud46mhxmahpv7v500fnj5ms0um"},"amount":{"value":"5370824911487300000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":85},"type":"execution","status":"success","account":{"address":"io15exa76f7h0rewfwfvy6rm8fkc4e5hrgy5q2hqm"},"amount":{"value":"2686948200006130000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":86},"type":"execution","status":"success","account":{"address":"io150zr6w4088z0xrd27u6pahj22dyq69jvym6kpr"},"amount":{"value":"92146377765141200000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":87},"type":"execution","status":"success","account":{"address":"io14wf0udxeunqjl0wp5cnkufmumjnv8q9e90zehe"},"amount":{"value":"74163610412132300000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":88},"type":"execution","status":"success","account":{"address":"io14pdasjny9kaagmqu05k8rjg4cngl07kyv70pna"},"amount":{"value":"10038822801870100000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":89},"type":"execution","status":"success","account":{"address":"io14fqw22gq2584cxyauuekp3mxxweddm0z0rm996"},"amount":{"value":"7469076435667700000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":90},"type":"execution","status":"success","account":{"address":"io14fczqdntsx5cuc03phmtkkhlspxku74smp6qqx"},"amount":{"value":"15233287250867600000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":91},"type":"execution","status":"success","account":{"address":"io13z29ktuaymuyv2fx2xp8kf9gq452qmvtj4yta7"},"amount":{"value":"2943587464794670000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":92},"type":"execution","status":"success","account":{"address":"io13rucfl7rwvrc9ldjl04ga7pmw2w72el38ykcls"},"amount":{"value":"518429045809640000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":93},"type":"execution","status":"success","account":{"address":"io13mtqzza9wlm326g2gwvg264mfagwdqhmqu6r7c"},"amount":{"value":"46149204079090300000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":94},"type":"execution","status":"success","account":{"address":"io13mms09vs68zv5y59nfn97khnvyhmxmm80q2623"},"amount":{"value":"1036150252248720000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":95},"type":"execution","status":"success","account":{"address":"io13mf9mjzutqk6qlj7q7ze35y0t8a4l209ydff9d"},"amount":{"value":"2574859143759090000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":96},"type":"execution","status":"success","account":{"address":"io132ezac8cy36sgld3x79p3fjxf5hwj5vqljuhvv"},"amount":{"value":"656778422420071000000","currency":{"symbol":"IOTX","decimals":18}}}]}]}}

# stakeCreate action
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},
    "block_identifier": {"index": 4034780}}' http://127.0.0.1:8080/block
#response:
#{"block":{"block_identifier":{"index":4034780,"hash":"bc1ad74d423f84e553602798e86019254b70d4499f1738a11c285ab9e31ea3b2"},"parent_block_identifier":{"index":4034779,"hash":"ac25b97cb7c9743b496cf45586d442ae4777753c77ca08d539bb96f30bca08c6"},"timestamp":1592801095000,"transactions":[{"transaction_identifier":{"hash":"9f261c47ad6611388c8e4569d2db378d2a7d98607c4259f5f9819ae6703742e6"},"operations":[{"operation_identifier":{"index":0},"type":"fee","status":"success","account":{"address":"io1mflp9m6hcgm2qcghchsdqj3z3eccrnekx9p0ms"},"amount":{"value":"-10000000000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":1},"type":"stakeCreate","status":"success","account":{"address":"io1mflp9m6hcgm2qcghchsdqj3z3eccrnekx9p0ms"},"amount":{"value":"-100000000000000000000","currency":{"symbol":"IOTX","decimals":18}}}]}]}}

# CandidateRegister,from mainnet
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "mainnet"},
    "block_identifier": {"index": 5160923}}' http://127.0.0.1:8080/block
#respone
#{"block":{"block_identifier":{"index":5160923,"hash":"fe653d92713bbd0f438fb8426edc3daab10a2eef28aa95d81fdaab460dd52f8e"},"parent_block_identifier":{"index":5160922,"hash":"40a7de13ef6f183c6c5dbb4bcb18aad3cc78e56c0c2ced1c1e1602bc4cc5c5e9"},"timestamp":1591033570000,"transactions":[{"transaction_identifier":{"hash":"a7e60af13e75646293c412b96e959d5edb8c254d3a1bcd9b6d7536a1af37e7c6"},"operations":[{"operation_identifier":{"index":0},"type":"fee","status":"success","account":{"address":"io180w8cmhrxpjjulry2g4zpyhm8mx933yhqkr4g6"},"amount":{"value":"-10000000000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":1},"type":"candidateRegister","status":"success","account":{"address":"io180w8cmhrxpjjulry2g4zpyhm8mx933yhqkr4g6"},"amount":{"value":"-2000000000000000000000000","currency":{"symbol":"IOTX","decimals":18}}}]}]}}

# StakeAddDeposit
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "testnet"},
    "block_identifier": {"index": 4066726}}' http://127.0.0.1:8080/block
#response
#{"block":{"block_identifier":{"index":4066726,"hash":"de51a09829c13480da1e5bb11d3ee907cd79094d3c5acfa0e66daf88fcfac31a"},"parent_block_identifier":{"index":4066725,"hash":"c5c3ee4bfadfcea70dccca9056ccd48c1bfc122279059f32a08a78ad454a33cc"},"timestamp":1592960825000,"transactions":[{"transaction_identifier":{"hash":"c0477d9d735ce2b5cd7d9f8e48b1be113988dfa0147f4ffc4055dc4e38ff751c"},"operations":[{"operation_identifier":{"index":0},"type":"fee","status":"success","account":{"address":"io1mflp9m6hcgm2qcghchsdqj3z3eccrnekx9p0ms"},"amount":{"value":"-10000000000000000","currency":{"symbol":"IOTX","decimals":18}}},{"operation_identifier":{"index":1},"type":"stakeAddDeposit","status":"success","account":{"address":"io1mflp9m6hcgm2qcghchsdqj3z3eccrnekx9p0ms"},"amount":{"value":"-1000000000000000000","currency":{"symbol":"IOTX","decimals":18}}}]}]}}

# claimFromRewardingFund,from mainnet
curl -X POST --data '{
    "network_identifier": {
        "blockchain": "IoTeX",
        "network": "mainnet"},
    "block_identifier": {"index": 315038}}' http://127.0.0.1:8080/block
