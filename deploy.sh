git add .
git commit -m "feat: otoGonder"
git push 

rm -rf main 

# GOOS=linux GOARCH=386 go build main.go
GOOS=linux GOARCH=amd64 go build main.go 
tar -cvzf static.tar.gz public/static/*

tar -cvzf view.tar.gz public/view/*

tar -cvzf locales.tar.gz public/locales/*

rm -rf CI/* 

mv static.tar.gz CI/static.tar.gz

mv view.tar.gz CI/view.tar.gz

mv locales.tar.gz CI/locales.tar.gz

#tar -xvf static.tar.gz

./deploy