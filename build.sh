echo "Build web page"
cd frontend
rm -rf ./frontend/dist
npm run build

cd ..

rm -rf ./backend/index/dist
cp -r ./frontend/dist ./backend/index/

cd backend
echo "Build server"
rm ./build/server
go build -o ./build/server sources
cp ./build/server ../app/server
cp -r ./index ../app/index
cd ..

