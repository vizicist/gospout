MAKE=mingw32-make.exe

echo "=========== Making libspout.a"
cd internal/libspout
$MAKE clean
$MAKE default
go build
go install
cd ../..

echo "=========== Making gospout-client"
cd cmd/gospout-client
go build
cd ../..

echo "=========== Making gospout-server"
cd cmd/gospout-server
go build
cd ../..
