/*
    1484: https://leetcode.cn/problems/group-sold-products-by-the-date/
 */

-- Write your MySQL query statement below

-- group_concat
select sell_date, 
        count(distinct product) num_sold, 
        group_concat(distinct product) products 
from Activities 
group by sell_date 
order by sell_date;


-- group_concat(distinct field order by field desc/asc separator ',')
SELECT sell_date,
    COUNT(DISTINCT product) AS num_sold,
    GROUP_CONCAT(DISTINCT product ORDER BY product asc SEPARATOR ',') AS products
FROM Activities 
GROUP BY sell_date
ORDER BY sell_date;
