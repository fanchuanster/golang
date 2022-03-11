module hello/hello

go 1.17

replace hello/greetings => ../greetings

require hello/greetings v0.0.0-00010101000000-000000000000 // indirect
