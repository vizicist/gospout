MAKE=mingw32-make.exe

echo "=========== Making libspout.a"
$MAKE clean
$MAKE default
go build
go install
