VERSION=$(grep VERSION core/version.go | cut -d \" -f2)

rm dist/*
VERSION=$VERSION bash release-build-all.sh