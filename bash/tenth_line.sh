# 195: https://leetcode.cn/problems/tenth-line/

# Read from the file file.txt and output the tenth line to stdout.

# line 1 and line 10
# sed -n '1p; 10p;' file.txt

# line 1 to line 10
# sed -n '1, 10p;' file.txt

# sed
sed -n 10p file.txt
sed -n "10p" file.txt

# tail + head
tail -n +10 file.txt | head -1

# awk
awk '{if(NR==10){print $0}}' file.txt
awk 'NR == 10' file.txt
