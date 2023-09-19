go build -o db_model .
mv db_model ../db_model
cd ..
./db_model
rm db_model