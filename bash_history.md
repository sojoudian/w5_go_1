```bash
cd go/src/github.com/sojoudian
mkdir w5_go_1
cd w5_go_1
Code .
open -a Visual\ Studio\ Code .
go mod init github.com/sojoudian/w5_go_1
git init
echo "# w5_go_1" >> README.md\ngit init\ngit add README.md\ngit commit -m "first commit"
git branch -M master
git remote add origin https://github.com/sojoudian/w5_go_1.git
git push -u origin master
git add . && git commit -m "add go mod" && git push
go run main.go
git add . && git commit -m "add main" && git push
go run main.go
cd /Users/maziar/Library/Mobile\ Documents/com\~apple\~CloudDocs/Code/go/src/github.com/sojoudian/w5_go_1
open .
go run main.go
go run *.go
go build -o myapp
ls
vi myapp
MAZ_LASTNAME="Sojoudian\n"
MAZ_LASTNAME="Sojoudian"
echo $MAZ_LASTNAME
GOOS=windows GOARCH=amd64 go build -o win_app
MAZ_LASTNAME="Sojoudian"
git add . && git commit -m "add binary" && git push
./win_app
./myapp
#git clone https://github.com/sojoudian/w5_go_1.git
#cd w5_go_1```