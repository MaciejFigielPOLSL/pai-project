cd frontend
npm run build
cd ..
rm -rf ./backend/index/dist
cp -r ./frontend/dist ./backend/index/
cd backend
rm ./build/server
go build -o ./build/server sources
cp ./build/server ../app/server
cp -r ./index ../app/index
cd ..

