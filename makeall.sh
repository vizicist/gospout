MAKE=mingw32-make.exe

echo "=========== Making libspout.a"
cd internal/libspout
$MAKE clean
$MAKE default
go build
go install
cd ../..

echo "=========== Making gospout-democlient"
cd cmd/gospout-democlient
go build
cd ../..

echo "=========== Making gospout-demoserver"
cd cmd/gospout-demoserver
go build
cd ../..
