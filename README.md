# IoTeX Gateway for Rosetta

This repository implements the [Rosetta](https://github.com/coinbase/rosetta-sdk-go) for the [IoTeX](https://iotex.io) blockchain.

To build the server:

	make

To run tests:

	make test

To clean-up:

	make clean


`make test` will automatically download the [Iotex node][0] and [rosetta-cli][2],
set up a test Oasis network, make some sample transactions, then run the
gateway and validate it using `rosetta-cli`.

[0]: https://github.com/iotexproject/iotex-core
[1]: https://github.com/coinbase/rosetta-sdk-go
[2]: https://github.com/coinbase/rosetta-cli
