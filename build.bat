cd frontend
npm run build
rm -rf backend\index\dist
cp frontend\dist backend\index\
cd backend
go build
