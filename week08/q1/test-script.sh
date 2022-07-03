rm -rf ./result_*

echo "================ Performance of 10 bytes ================" >> ./result_10
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 10 >> ./result_10
echo "================ Performance of 20 bytes ================" >> ./result_20
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 20 >> ./result_20
echo "================ Performance of 50 bytes ================" >> ./result_50
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 50 >> ./result_50
echo "================ Performance of 100 bytes ================" >> ./result_100
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 100 >> ./result_100
echo "================ Performance of 200 bytes ================" >> ./result_200
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 200 >> ./result_200
echo "================ Performance of 1000 bytes ================" >> ./result_1000
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 1000 >> ./result_1000
echo "================ Performance of 5000 bytes ================" >> ./result_5000
docker exec -it redis redis-benchmark -p 6379 -t set,get -d 5000 >> ./result_5000