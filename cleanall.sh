MAKE=mingw32-make.exe

cd internal/libspout
$MAKE clean
cd ../..

rm -f cmd/gospout-client/*.exe
rm -f cmd/gospout-server/*.exe
